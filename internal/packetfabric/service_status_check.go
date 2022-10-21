package packetfabric

import "time"

type ServiceState struct {
	CircuitID string `json:"circuit_id"`
	Status    Status `json:"status"`
}
type Object struct {
	State   string `json:"state"`
	Deleted bool   `json:"deleted"`
}
type Current struct {
	State       string `json:"state"`
	Description string `json:"description"`
	Warning     bool   `json:"warning"`
}
type Progress struct {
	Position int `json:"position"`
	Steps    int `json:"steps"`
}
type States struct {
	State       string `json:"state"`
	Description string `json:"description"`
}
type LastWorkflow struct {
	Name         string   `json:"name"`
	Root         string   `json:"root"`
	Current      string   `json:"current"`
	State        string   `json:"state"`
	CurrentState string   `json:"current_state"`
	PrevState    string   `json:"prev_state"`
	IsFinal      bool     `json:"is_final"`
	Progress     Progress `json:"progress"`
}
type Status struct {
	Object       Object       `json:"object"`
	Current      Current      `json:"current"`
	LastWorkflow LastWorkflow `json:"last_workflow"`
}

func (c *PFClient) CheckServiceStatus(ch chan bool, fn func() (*ServiceState, error)) {
	ticker := time.NewTicker(5 * time.Second)
	var count int
	for range ticker.C {
		count = count + 1
		state, serviceErr := fn()
		if serviceErr != nil && count == 0 {
			ch <- false
		}
		if state != nil {
			if state.Status.Current.State == "COMPLETE" {
				ticker.Stop()
				ch <- true
			} else if state.Status.LastWorkflow.Progress.Position ==
				state.Status.LastWorkflow.Progress.Steps && state.Status.LastWorkflow.IsFinal {
				ticker.Stop()
				ch <- true
			}
		}
		if serviceErr != nil && count > 0 {
			ticker.Stop()
			ch <- true
		}
	}
}

func (c *PFClient) CheckIPSecStatus(ch chan bool, fn func() (*ServiceState, error)) {
	ticker := time.NewTicker(5 * time.Second)
	var count int
	for range ticker.C {
		count = count + 1
		state, serviceErr := fn()
		if serviceErr != nil && count == 0 {
			ch <- false
		}
		if state != nil {
			if state.Status.Current.State == "BGP_NOT_CREATED" {
				ticker.Stop()
				ch <- true
			}
		}
		if serviceErr != nil && count > 0 {
			ticker.Stop()
			ch <- true
		}
	}
}
