package strategicreserve

import (
	//#nosec
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sge-network/sge/testutil/sample"
	srsimulation "github.com/sge-network/sge/x/strategicreserve/simulation"
	"github.com/sge-network/sge/x/strategicreserve/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = srsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	//#nosec
	opWeightMsgAddMarket          = "op_weight_msg_create_chain"
	defaultWeightMsgAddMarket int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	srsimulation.RandomizedGenState(simState)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return srsimulation.ParamChanges(r)
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.StoreKey] = srsimulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddMarket int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddMarket, &weightMsgAddMarket, nil,
		func(_ *rand.Rand) {
			weightMsgAddMarket = defaultWeightMsgAddMarket
		},
	)

	return operations
}
