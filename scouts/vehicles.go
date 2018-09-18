package scouts

type Vehicle struct {
	Id int

	Name            string     `json:"name"`
	StateVariables  []string   `json:"state_variables"`
	DefaultDynamics string     `json:"default_dynamics"`
	DynamicsModes   []Dynamics `json:"dynamic_modes"`
	ScriptCompiler  string     `json:"script_compiler"`
}

type Dynamics struct {
}