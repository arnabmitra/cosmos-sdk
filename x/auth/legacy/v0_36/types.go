// DONTCOVER
// nolint
package v036

import v034auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v0_34"

const (
	ModuleName = "auth"
)

type (
	GenesisState struct {
		Params v034auth.Params `json:"params"`
	}
)

func NewGenesisState(params v034auth.Params) GenesisState {
	return GenesisState{params}
}
