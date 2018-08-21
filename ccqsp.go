package gomers

import (
	"encoding/json"
	"fmt"
)

type CCQSP struct {
	Name              string             `json:"name"`
	Id                int                `json:"id"`
	Objective         string             `json:"objective"`
	Events            []Event            `json:"events"`
	Episodes          []Episode          `json:"episodes"`
	ChanceConstraints []ChanceConstraint `json:"chanceConstraints"`
	Requirements      []string           `json:"require"`
	StateSpace        StateSpace         `json:"stateSpace"`
}

type Event struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Episode struct {
	Name              string    `json:"name"`
	Id                int       `json:"id"`
	FromEvent         int       `json:"fromEvent"`
	ToEvent           int       `json:"toEvent"`
	Predicate         Predicate `json:"predicate"`
	PredicateDuration string    `json:"predicateDuration"`
}

type Predicate struct {
	Type string `json:"type"`
}

type ChanceConstraint struct {
	Name               string  `json:"name"`
	Id                 int     `json:"id"`
	ConstraintIds      []int   `json:"constraints"`
	FailureProbability float64 `json:"failureProbability"`
}

type StateSpace struct {
	AgentAssignments []AgentAssignment `json:"agentAssignments"`
}

type AgentAssignment struct {
	AgentId           int                 `json:"agent"`
	StateAssignment   ParameterAssignment `json:"stateAssignment"`
	ControlAssignment ParameterAssignment `json:"controlAssignment"`
}

type ParameterAssignment struct {
	Type        string       `json:"type"`
	Assignments []Assignment `json:"assignment"`
}

type Assignment struct {
	Time       float64
	Parameters interface{}
}

// Adapted from http://eagain.net/articles/go-json-array-to-struct/
func (a *Assignment) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&a.Time, &a.Parameters}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Assignment: %d != %d", g, e)
	}
	return nil
}
