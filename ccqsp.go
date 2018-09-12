package gomers

import (
	"encoding/json"
	"fmt"
)

type CCQSP struct {
	Name              string                  `json:"name"`
	Id                int                     `json:"id"`
	Objective         string                  `json:"objective"`
	Events            []Event                 `json:"events"`
	Episodes          []Episode               `json:"episodes"`
	ChanceConstraints []ChanceConstraint      `json:"chanceConstraints"`
	Requirements      []string                `json:"require"`
	Annotations       *map[string]interface{} `json:"annotations"`
	StateSpace        StateSpace              `json:"stateSpace"`
}

type Event struct {
	Name        string                  `json:"name"`
	Id          int                     `json:"id"`
	Annotations *map[string]interface{} `json:"annotations"`
}

type Episode struct {
	Name              string                  `json:"name"`
	Id                int                     `json:"id"`
	FromEvent         int                     `json:"fromEvent"`
	ToEvent           int                     `json:"toEvent"`
	LowerBound        *int                    `json:"lowerBound"`
	UpperBound        *int                    `json:"upperBound"`
	Predicate         Predicate               `json:"predicate"`
	PredicateDuration string                  `json:"predicateDuration"`
	Annotations       *map[string]interface{} `json:"annotations"`
}

type Predicate struct {
	Type        string                  `json:"type"`
	Annotations *map[string]interface{} `json:"annotations"`
}

type ChanceConstraint struct {
	Name               string                  `json:"name"`
	Id                 int                     `json:"id"`
	ConstraintIds      []int                   `json:"constraints"`
	FailureProbability float64                 `json:"failureProbability"`
	Annotations        *map[string]interface{} `json:"annotations"`
}

type StateSpace struct {
	EventAssignments []EventAssignment `json:"eventAssignments"`
	AgentAssignments []AgentAssignment `json:"agentAssignments"`
}

type EventAssignment struct {
	Id         int     `json:"event"`
	LowerBound float64 `json:"lowerBound"`
	UpperBound float64 `json:"upperBound"`
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
