package simulation

// DONTCOVER

import (
	"encoding/json"
	//#nosec
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/sge-network/sge/x/house/types"
	"github.com/spf13/cast"
)

// Simulation parameter constants
const (
	HouseParticipationFee = "HouseParticipationFee"
	MinDeposit            = "MinDeposit"
)

// GenHouseParticipationFee randomized bathc settlement count
func GenHouseParticipationFee(r *rand.Rand) sdk.Dec {
	return sdk.NewDec(cast.ToInt64(r.Intn(99)))
}

// GenMinDeposit randomized house by uid query count
func GenMinDeposit(r *rand.Rand) sdk.Int {
	return sdk.NewInt(cast.ToInt64(r.Intn(99)))
}

// RandomizedGenState generates a random GenesisState for house
func RandomizedGenState(simState *module.SimulationState) {
	var houseParticipationFee sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, HouseParticipationFee, &houseParticipationFee, simState.Rand,
		func(r *rand.Rand) { houseParticipationFee = GenHouseParticipationFee(r) },
	)

	var minDeposit sdk.Int
	simState.AppParams.GetOrGenerate(
		simState.Cdc, MinDeposit, &minDeposit, simState.Rand,
		func(r *rand.Rand) { minDeposit = GenMinDeposit(r) },
	)

	defaultGenesis := types.DefaultGenesis()

	_, err := json.MarshalIndent(&defaultGenesis, "", " ")
	if err != nil {
		panic(err)
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(defaultGenesis)
}
