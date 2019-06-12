package data

import (
	"errors"

	"github.com/dapperlabs/bamboo-emulator/crypto"
)

// WorldState represents the current state of the blockchain.
type WorldState struct {
	Blocks       map[crypto.Hash]Block
	Collections  map[crypto.Hash]Collection
	Transactions map[crypto.Hash]Transaction
	Blockchain   []Block
}

// NewWorldState returns a new empty world state.
func NewWorldState() *WorldState {
	return &WorldState{
		Blocks:       make(map[crypto.Hash]Block),
		Collections:  make(map[crypto.Hash]Collection),
		Transactions: make(map[crypto.Hash]Transaction),
		Blockchain:   make([]Block, 0),
	}
}

func (s *WorldState) GetLatestBlock() *Block {
	currHeight := len(s.Blockchain)
	return &s.Blockchain[currHeight-1]
}

func (s *WorldState) GetBlockByNumber(n uint64) (*Block, error) {
	currHeight := len(s.Blockchain)
	if int(n) < currHeight {
		return &s.Blockchain[n], nil
	}

	return &Block{}, errors.New("invalid Block number: Block number exceeds blockchain length")
}

func (s *WorldState) GetBlockByHash(h crypto.Hash) (*Block, error) {
	if block, ok := s.Blocks[h]; ok {
		return &block, nil
	}

	return &Block{}, errors.New("invalid Block hash: Block doesn't exist")

}

func (s *WorldState) GetCollection(h crypto.Hash) (*Collection, error) {
	if collection, ok := s.Collections[h]; ok {
		return &collection, nil
	}

	return &Collection{}, errors.New("invalid Collection hash: Collection doesn't exist")
}

func (s *WorldState) GetTransaction(h crypto.Hash) (*Transaction, error) {
	if tx, ok := s.Transactions[h]; ok {
		return &tx, nil
	}

	return &Transaction{}, errors.New("invalid Transaction hash: Transaction doesn't exist")
}

func (s *WorldState) AddBlock(block *Block) error {
	// TODO: add to block map and chain
	return nil
}

func (s *WorldState) InsertCollection(col *Collection) error {
	// TODO: add to collection map
	return nil
}

func (s *WorldState) InsertTransaction(tx *Transaction) error {
	if _, exists := s.Transactions[tx.Hash()]; exists {
		return errors.New("transaction exists")
	}

	s.Transactions[tx.Hash()] = *tx

	return nil
}

func (s *WorldState) SealBlock(h crypto.Hash) error {
	// TODO: seal the block
	return nil
}
