package hotstuff

import (
	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/model/flow"
)

// RandomBeaconReconstructor collects signature shares, and reconstructs the
// group signature with enough shares.
type RandomBeaconReconstructor interface {
	// Verify verifies the signature under the stored public key corresponding to the signerID, and the stored message agreed about upfront.
	Verify(signerID flow.Identifier, sig crypto.Signature) (bool, error)

	// TrustedAdd adds the signature share to the reconstructors internal
	// state. Validity of signature is not checked. It is up to the
	// implementation, whether it still adds a signature or not, when the
	// minimal number of required sig shares has already been reached,
	// because the reconstructed group signature is the same.
	// It returns:
	//  - (true, nil) if and only if enough signature shares were collected
	//  - (false, nil) if not enough shares were collected
	//  - (false, error) if there is exception adding the sig share)
	TrustedAdd(signerID flow.Identifier, sig crypto.Signature) (hasSufficientShares bool, err error)

	// HasSufficientShares returns true if and only if reconstructor
	// has collected a sufficient number of signature shares.
	HasSufficientShares() bool

	// Reconstruct reconstructs the group signature.
	// The reconstructed signature is verified against the overall group public key and the message agreed upon.
	// This is a sanity check that is necessary since "TrustedAdd" allows adding non-verified signatures.
	// Reconstruct returns an error if the reconstructed signature fails the sanity verification, or if not enough shares have been collected.
	Reconstruct() (crypto.Signature, error)
}

// SigType is the aggregable signature type.
type SigType int

// SigType specifies the role of the signature in the protocol. SigTypeRandomBeacon type is for random beacon signatures. SigTypeStaking is for Hotstuff sigantures. Both types are aggregatable cryptographic signatures.
const (
	SigTypeStaking SigType = iota
	SigTypeRandomBeacon
	SigTypeInvalid
)

// WeightedSignatureAggregator aggregates signatures of the same signature scheme and the same message from different signers.
// The public keys and message are aggreed upon upfront.
// It is also recommended to only aggregate signatures generated with keys representing equivalent security-bit level.
// The module is aware of weights assigned to each signer, as well as a total weight threshold.
// Implementation of SignatureAggregator must be concurrent safe.
type WeightedSignatureAggregator interface {
	// Verify verifies the signature under the stored public key corresponding to the signerID, and the stored message.
	Verify(signerID flow.Identifier, sig crypto.Signature) (bool, error)

	// TrustedAdd adds a signature to the internal set of signatures.
	// It adds the signer's weight to the total collected weight and returns the total weight regardless
	// of the returned error.
	// The function errors if a signature from the signerID was already collected.
	TrustedAdd(signerID flow.Identifier, sig crypto.Signature) (totalWeight uint64, exception error)

	// TotalWeight returns the total weight presented by the collected signatures.
	TotalWeight() uint64

	// Aggregate assumes enough weights have been collected, it aggregates the signatures
	// and return the aggregated signature.
	// If called concurrently, only one thread will be running the aggregation.
	// Aggregate attempts to aggregate the internal signatures and returns the resulting signature data.
	// It errors if not enough weights have been collected.
	// The function performs a final verification and errors if the aggregated signature is not valid. This is
	// required for the function safety since "TrustedAdd" allows adding invalid signatures.
	// If called concurrently, only one thread will be running the aggregation.
	Aggregate() ([]flow.Identifier, []byte, error)
}

// CombinedSigAggregator aggregates the staking signatures and random beacon signatures,
// and keep track of the total weights represented by each signature share. And report whether
// sufficient weights for representing the majority of stakes have been collected. If yes, then aggregate
// the signatures.
type CombinedSigAggregator interface {
	// Verify verifies the signature under the stored public key corresponding to the signerID and the stored message.
	// `sigType` specifies the type of the input signature (random beacon or hotstuff), which helps the module pick the right stored public key and message.
	Verify(signerID flow.Identifier, sig crypto.Signature, sigType SigType) (bool, error)

	// TrustedAdd adds the signature to staking signatures store or random beacon signature store
	// based on the given sig type.
	// It returns:
	//  - (false, nil) if the sig share is added, but the total stake weight represented by the collected
	//    signatures can not represent the majority.
	//  - (true, nil) if the sig share is added, and sufficient stake weight has been collected to represent
	//    the majority.
	//  - (false, exception) if there is any exception adding the signature.
	TrustedAdd(signerID flow.Identifier, sig crypto.Signature, sigType SigType) (hasSufficientWeight bool, exception error)

	// HasSufficientWeight returns whether enough signatures have been collected to represent
	// stake majority.
	HasSufficientWeight() bool

	// Aggregate assumes enough shares have been collected, and aggregates the signatures.
	// Note we don't mix the staking sig and random beacon sig when aggregating them,
	// Instead, they are aggregated separately and returned separately.
	Aggregate() (aggregatedStakingSig []byte, aggregatedRandomBeaconSig []byte, exception error)
}

// BlockSignatureData is an intermediate struct for Packer to pack the
// aggregated signature data into raw bytes or unpack from raw bytes.
type BlockSignatureData struct {
	StakingSigners               []flow.Identifier
	RandomBeaconSigners          []flow.Identifier
	AggregatedStakingSig         []byte // if BLS is used, this is equivalent to crypto.Signature
	AggregatedRandomBeaconSig    []byte // if BLS is used, this is equivalent to crypto.Signature
	ReconstructedRandomBeaconSig crypto.Signature
}

// Packer packs aggregated signature data into raw bytes to be used in block header.
type Packer interface {
	// blockID is the block that the aggregated signature is for.
	// sig is the aggregated signature data.
	Pack(blockID flow.Identifier, sig *BlockSignatureData) ([]flow.Identifier, []byte, error)

	// blockID is the block that the aggregated sig is signed for
	// sig is the aggregated signature data
	Unpack(blockID flow.Identifier, signerIDs []flow.Identifier, sigData []byte) (*BlockSignatureData, error)
}
