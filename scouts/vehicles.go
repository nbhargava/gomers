package scouts

type Vehicle struct {
	Name            string     `json:"name"`
	StateVariables  []string   `json:"state_variables"`
	DefaultDynamics string     `json:"default_dynamics"`
	DynamicsModes   []Dynamics `json:"dynamic_modes"`
	ScriptCompiler  string     `json:"script_compiler"`
}

type Dynamics struct {
	Name        string      `json:"name"`
	StateBounds StateBounds `json:"stateBounds"`
}

type StateBounds struct {
	StateUpperBounds []*float64 `json:"stateUpperBounds"`
	StateLowerBounds []*float64 `json:"stateLowerBounds"`
}
