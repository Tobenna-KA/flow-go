package reactor

import (
	"github.com/dapperlabs/flow-go/engine/consensus/componentsdef"
	"github.com/dapperlabs/flow-go/engine/consensus/componentsdefConAct"
	"github.com/dapperlabs/flow-go/engine/consensus/componentsreactor/core"
	"github.com/dapperlabs/flow-go/engine/consensus/componentsreactor/forkchoice"
	"github.com/juju/loggo"
)

var ConsensusLogger loggo.Logger

type Reactor struct {
	core       *core.ReactorCore
	forkchoice forkchoice.ForkChoice

	forkchoiceRequests chan uint64
	newQCs             chan *def.QuorumCertificate
	newBlockProposals  chan *def.Block
}

func NewReactor(finalizer *core.ReactorCore, forkchoice forkchoice.ForkChoice) *Reactor {
	return &Reactor{
		core:               finalizer,
		forkchoice:         forkchoice,
		forkchoiceRequests: make(chan uint64, 10),
		newQCs:             make(chan *def.QuorumCertificate, 10),
		newBlockProposals:  make(chan *def.Block, 300),
	}
}

func (r *Reactor) OnBlockProposalTrigger(view uint64) {
	// inspired by https://content.pivotal.io/blog/a-channel-based-ring-buffer-in-go
	select {
	case r.forkchoiceRequests <- view:
	default:
		<-r.forkchoiceRequests
		r.forkchoiceRequests <- view
	}
}

func (r *Reactor) OnQcFromVotes(qc *def.QuorumCertificate) {
	// inspired by https://content.pivotal.io/blog/a-channel-based-ring-buffer-in-go
	select {
	case r.newQCs <- qc:
	default:
		<-r.newQCs
		r.newQCs <- qc
	}
}

func (r *Reactor) OnReceivedViewChange(qc *def.QuorumCertificate) {
	// inspired by https://content.pivotal.io/blog/a-channel-based-ring-buffer-in-go
	select {
	case r.newQCs <- qc:
	default:
		<-r.newQCs
		r.newQCs <- qc
	}
}

func (r *Reactor) OnReceivedBlockProposal(viewChange *defConAct.ViewChange) {
	r.OnReceivedViewChange(viewChange.QC)
}

func (r *Reactor) Run() {
	go r.run()
}

func (r *Reactor) run() {
	for {
		select {
		case view := <-r.forkchoiceRequests:
			fc := r.forkchoice.GenerateForkChoice()
			if fc.View <= view {
				// ToDo trigger Block proposal
			} else {
				ConsensusLogger.Warningf("Dropped Block Proposal Trigger for view %d as it was already stale", view)
			}
		case qc := <-r.newQCs:
			r.forkchoice.ProcessQC(qc)
		case block := <-r.newBlockProposals:
			if r.forkchoice.IsProcessingNeeded(block.BlockMRH, block.View) {
				r.forkchoice.ProcessBlock(block)
			}
		}
	}
}

func (r *Reactor) isFromExpectedPrimary(proposal *def.BlockProposal) bool {
	// ToDo implement
	return true
}
