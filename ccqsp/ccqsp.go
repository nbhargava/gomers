package ccqsp

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
	Type        string
	Args        map[string]interface{}
	Annotations *map[string]interface{}
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
	Parameters *[]float64
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

func (a *Assignment) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		a.Time,
		a.Parameters,
	})
}

func (p *Predicate) UnmarshalJSON(buf []byte) error {
	var predicate map[string]interface{}
	if err := json.Unmarshal(buf, &predicate); err != nil {
		return err
	}

	predType, ok := predicate["type"].(string)
	if !ok {
		return fmt.Errorf("`type` not provided in %v", predicate)
	}
	p.Type = predType
	delete(predicate, "type")

	annotations, ok := predicate["annotations"].(map[string]interface{})
	if ok {
		p.Annotations = &annotations
		delete(predicate, "annotations")
	}

	// Check for required fields
	// TODO: Do this for things other than Dynamics
	switch predType {
	case "dynamics":
		_, ok = predicate["dynamicsMode"].(string)
		if !ok {
			return fmt.Errorf("`dynamicsMode` not properly specified for dynamics predicate in %v", predicate)
		}

		agents, ok := predicate["agents"].([]interface{})
		agentInts := []int{}
		for _, agent := range agents {
			a, ok := agent.(float64)
			if !ok {
				return fmt.Errorf("`agents` must be a list of ids in dynamics predicate in %v", predicate)
			}
			agentInts = append(agentInts, int(a))
		}
		predicate["agents"] = agentInts
		if !ok {
			return fmt.Errorf("`agents` not properly specified for dynamics predicate in %v", predicate)
		}
	}

	p.Args = predicate
	return nil
}

func (a *Predicate) MarshalJSON() ([]byte, error) {
	jsonMap := map[string]interface{}{}
	jsonMap["type"] = a.Type
	jsonMap["annotations"] = a.Annotations
	for k, v := range a.Args {
		jsonMap[k] = v
	}
	return json.Marshal(jsonMap)
}
