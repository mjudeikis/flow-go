package module

import (
	"github.com/dapperlabs/flow-go/model/flow"
)

type HandleFunc func(originID flow.Identifier, entity flow.Entity) error

type Requester interface {
	Request(entityID flow.Identifier, handle HandleFunc) error
}
