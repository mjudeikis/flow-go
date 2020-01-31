package types

type BlockProposal struct {
	Block     *Block
	Signature *Signature // CAUTION: this is sign(Block), i.e. it does NOT include ConsensusPayload
}

func NewBlockProposal(block *Block, sig *Signature) *BlockProposal {
	return &BlockProposal{
		Block:     block,
		Signature: sig,
	}
}

func (b *BlockProposal) QC() *QuorumCertificate { return b.Block.QC }
func (b *BlockProposal) View() uint64           { return b.Block.View }
func (b *BlockProposal) BlockID() []byte        { return b.Block.BlockID() }
func (b *BlockProposal) Height() uint64         { return b.Block.Height }
