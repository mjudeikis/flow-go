package broadcast_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dapperlabs/flow-go/module"
	"github.com/dapperlabs/flow-go/module/broadcast"
)

// Counter implements a thread-safe counter for tracking received broadcasts.
type Counter struct {
	mu    sync.Mutex
	count int
}

// Inc increments the counter by one.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

// N returns the current count.
func (c *Counter) N() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func TestBroadcast(t *testing.T) {

	// broadcast should be a no-op without subscribers
	t.Run("no subscribers", func(t *testing.T) {
		caster := broadcast.NewBroadcaster()
		done := make(chan struct{})

		go func() {
			caster.Broadcast()
			close(done)
		}()

		select {
		case <-time.After(time.Second):
			t.Fail()
		case <-done:
		}
	})

	t.Run("with subscribers", func(t *testing.T) {
		var (
			caster      = broadcast.NewBroadcaster()
			N           = 1
			subscribers = make([]module.Subscription, N)
			wg          sync.WaitGroup
			counter     Counter
		)

		for i := range subscribers {
			sub := caster.Subscribe()
			subscribers[i] = sub

			// subscribers increment the counter when they receive a broadcast
			// once, then exit
			wg.Add(1)
			go func() {
				defer wg.Done()

				select {
				case <-time.After(time.Second):
					t.Log("timeout")
					t.Fail()
				case <-sub.Ch():
					counter.Inc()
				}
			}()
		}

		caster.Broadcast()
		wg.Wait()

		assert.Equal(t, len(subscribers), counter.N())
	})

	// broadcasts should only be sent to active subscriptions
	t.Run("with some unsubscribed subscribers", func(t *testing.T) {
		var (
			caster      = broadcast.NewBroadcaster()
			N           = 100
			subscribers = make([]module.Subscription, N)
			wg          sync.WaitGroup
			counter     Counter // keep track of received broadcasts
		)

		for i := range subscribers {
			sub := caster.Subscribe()
			subscribers[i] = sub

			wg.Add(1)
			go func() {
				defer wg.Done()

				select {
				case <-time.After(time.Second):
					t.Fail()
				case _, ok := <-sub.Ch():
					if !ok {
						// channel closed: exit and don't increment the counter
						return
					}
					counter.Inc()
				}
			}()
		}

		// unsubscribe half of the subscribers
		// the unsubscribed subscribers should have exited their goroutine but
		// not added to the counter
		for i := 0; i < N/2; i++ {
			subscribers[i].Unsubscribe()
		}

		caster.Broadcast()
		wg.Wait()

		assert.Equal(t, N/2, counter.N())
	})

	// broadcasts to a full subscription channel should be dropped
	t.Run("with busy subscriber", func(t *testing.T) {
		caster := broadcast.NewBroadcaster()

		// create one subscription whose channel isn't handled
		sub := caster.Subscribe()

		done := make(chan struct{})

		// broadcast several times
		go func() {
			caster.Broadcast()
			caster.Broadcast()
			close(done)
		}()

		// the broadcasts should skip when channel is full, and not block
		select {
		case <-time.After(time.Second):
			t.Fail()
		case <-done:
			// both broadcasts completed
		}

		// should be able to read broadcast after-the-fact
		select {
		case <-time.After(time.Second):
			t.Fail()
		case <-sub.Ch():
			// read broadcast successfully
		}
	})

	t.Run("with unsubscribed subscriber", func(t *testing.T) {
		caster := broadcast.NewBroadcaster()

		// create one subscription and immediately unsubscribe
		sub := caster.Subscribe()
		sub.Unsubscribe()

		done := make(chan struct{})

		// broadcast several times
		go func() {
			caster.Broadcast()
			caster.Broadcast()
			close(done)
		}()

		// neither broadcast should block
		select {
		case <-time.After(time.Second):
			t.Fail()
		case <-done:
			// both broadcasts completed
		}
	})
}
