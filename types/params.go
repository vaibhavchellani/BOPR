package types

// Params stores all the parameters which are maintained on-chain and keeps updating
// them whenever they change on-chain
type Params struct {
	StakeAmount uint `json:"stakeAmount"`
}