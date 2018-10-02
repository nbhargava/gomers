package scouts

type Vehicle struct {
	Name            string     `json:"name"`
	StateVariables  []string   `json:"stateVariables"`
	DefaultDynamics string     `json:"defaultDynamics"`
	DynamicsModes   []Dynamics `json:"dynamicModes"`
	ScriptCompiler  string     `json:"scriptCompiler"`
}

type Dynamics struct {
	Name        string      `json:"name"`
	StateBounds StateBounds `json:"stateBounds"`
}

type StateBounds struct {
	StateUpperBounds []*float64 `json:"stateUpperBounds"`
	StateLowerBounds []*float64 `json:"stateLowerBounds"`
}
