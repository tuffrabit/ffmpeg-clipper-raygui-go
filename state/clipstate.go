package state

import (
	"fmt"
	"strconv"
)

type ClipState struct {
	StartSeconds float64
	StartInput   string
	EndSeconds   float64
	EndInput     string
}

func (cs *ClipState) Reset() {
	cs.StartSeconds = 0
	cs.StartInput = ""
	cs.EndSeconds = 0
	cs.EndInput = ""
}

func (cs *ClipState) SetStart(input string) error {
	if input == cs.StartInput {
		return nil
	}

	if input == "" {
		cs.StartSeconds = 0
		cs.StartInput = input
		return nil
	}

	v, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return fmt.Errorf("state.ClipState.SetStart: failed to convert input to float, error: %w", err)
	}
	cs.StartSeconds = v
	cs.StartInput = input

	return nil
}

func (cs *ClipState) SetEnd(input string) error {
	if input == cs.EndInput {
		return nil
	}

	if input == "" {
		cs.EndSeconds = 0
		cs.EndInput = input
		return nil
	}

	v, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return fmt.Errorf("state.ClipState.SetEnd: failed to convert input to float, error: %w", err)
	}
	cs.EndSeconds = v
	cs.EndInput = input

	return nil
}
