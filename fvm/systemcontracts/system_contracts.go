// Package systemcontracts stores canonical address locations for all system
// smart contracts and service events.
//
// System contracts are special smart contracts controlled by the service account,
// a Flow account with special privileges to administer the network.
//
// Service events are special events defined within system contracts which
// are included within execution receipts and processed by the consensus committee
// to enable message-passing to the protocol state.
//
// For transient networks, all system contracts can be deployed to the service
// account. For long-lived networks, system contracts are spread across several
// accounts for historical reasons.
package systemcontracts

import (
	"fmt"

	"github.com/onflow/flow-go/model/flow"
)

const (

	// Unqualified names of system smart contracts (not including address prefix)

	ContractNameEpoch     = "FlowEpoch"
	ContractNameClusterQC = "FlowClusterQC"
	ContractNameDKG       = "FlowDKG"

	// Unqualified names of service events (not including address prefix or contract name)

	EventNameEpochSetup  = "EpochSetup"
	EventNameEpochCommit = "EpochCommitted"

	// Format strings for qualified service event names (including address prefix)

	serviceEventTypeIDFormat = "A.%s.%s.%s"
)

// SystemContract represents a system contract on a particular chain.
type SystemContract struct {
	Address flow.Address
	Name    string
}

// ServiceEvent represents a service event on a particular chain.
type ServiceEvent struct {
	Address       flow.Address
	ContractName  string
	Name          string
	QualifiedType flow.EventType
}

// QualifiedIdentifier returns the Cadence qualified identifier of the service
// event, which includes the contract name and the event type name.
func (se ServiceEvent) QualifiedIdentifier() string {
	return fmt.Sprintf("%s.%s", se.ContractName, se.Name)
}

// EventType returns the full event type identifier, including the address, the
// contract name, and the event type name.
func (se ServiceEvent) EventType() flow.EventType {
	return flow.EventType(fmt.Sprintf("A.%s.%s.%s", se.Address, se.ContractName, se.Name))
}

// SystemContracts is a container for all system contracts on a particular chain.
type SystemContracts struct {
	Epoch     SystemContract
	ClusterQC SystemContract
	DKG       SystemContract
}

// ServiceEvents is a container for all service events on a particular chain.
type ServiceEvents struct {
	EpochSetup  ServiceEvent
	EpochCommit ServiceEvent
}

// All returns all service events as a slice.
func (se ServiceEvents) All() []ServiceEvent {
	return []ServiceEvent{
		se.EpochSetup,
		se.EpochCommit,
	}
}

// SystemContractsForChain returns the system contract configuration for the given chain.
func SystemContractsForChain(chainID flow.ChainID) (*SystemContracts, error) {
	addresses, ok := contractAddressesByChainID[chainID]
	if !ok {
		return nil, fmt.Errorf("unknown chain id (%s)", chainID.String())
	}

	contracts := &SystemContracts{
		Epoch: SystemContract{
			Address: addresses[ContractNameEpoch],
			Name:    ContractNameEpoch,
		},
		ClusterQC: SystemContract{
			Address: addresses[ContractNameClusterQC],
			Name:    ContractNameClusterQC,
		},
		DKG: SystemContract{
			Address: addresses[ContractNameDKG],
			Name:    ContractNameDKG,
		},
	}

	return contracts, nil
}

// ServiceEventsForChain returns the service event confirmation for the given chain.
func ServiceEventsForChain(chainID flow.ChainID) (*ServiceEvents, error) {
	addresses, ok := contractAddressesByChainID[chainID]
	if !ok {
		return nil, fmt.Errorf("unknown chain id (%s)", chainID.String())
	}

	events := &ServiceEvents{
		EpochSetup: ServiceEvent{
			Address:      addresses[ContractNameEpoch],
			ContractName: ContractNameEpoch,
			Name:         EventNameEpochSetup,
		},
		EpochCommit: ServiceEvent{
			Address:      addresses[ContractNameEpoch],
			ContractName: ContractNameEpoch,
			Name:         EventNameEpochCommit,
		},
	}

	return events, nil
}

// serviceEventQualifiedType returns the qualified event type for the
// given service event format string and address location.
func serviceEventQualifiedType(format string, addr flow.Address) flow.EventType {
	return flow.EventType(fmt.Sprintf(format, addr))
}

// contractAddressesByChainID stores the default system smart contract
// addresses for each chain.
var contractAddressesByChainID map[flow.ChainID]map[string]flow.Address

func init() {
	contractAddressesByChainID = make(map[flow.ChainID]map[string]flow.Address)

	// Main Flow network
	// TODO need these address values
	mainnet := map[string]flow.Address{
		ContractNameEpoch:     flow.EmptyAddress,
		ContractNameClusterQC: flow.EmptyAddress,
		ContractNameDKG:       flow.EmptyAddress,
	}
	contractAddressesByChainID[flow.Mainnet] = mainnet

	// Long-lived test networks
	// TODO need these address values
	testnet := map[string]flow.Address{
		ContractNameEpoch:     flow.EmptyAddress,
		ContractNameClusterQC: flow.EmptyAddress,
		ContractNameDKG:       flow.EmptyAddress,
	}
	contractAddressesByChainID[flow.Testnet] = testnet
	contractAddressesByChainID[flow.Canary] = testnet

	// Transient test networks
	// All system contracts are deployed to the service account
	transient := map[string]flow.Address{
		ContractNameEpoch:     flow.Emulator.Chain().ServiceAddress(),
		ContractNameClusterQC: flow.Emulator.Chain().ServiceAddress(),
		ContractNameDKG:       flow.Emulator.Chain().ServiceAddress(),
	}
	contractAddressesByChainID[flow.Emulator] = transient
	contractAddressesByChainID[flow.Localnet] = transient
	contractAddressesByChainID[flow.Benchnet] = transient
}
