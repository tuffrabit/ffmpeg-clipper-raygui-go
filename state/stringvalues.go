package state

import (
	"fmt"
	"strconv"
)

type StringValue interface {
	Set(input string) error
}

type IntStringValue struct {
	Input string
	Value int
}

func (sv *IntStringValue) Set(input string) error {
	if input == sv.Input {
		return nil
	}

	v, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("state.IntStringValue.Set: failed to convert input to int, error: %w", err)
	}
	sv.Value = v
	sv.Input = input

	return nil
}

type Float32StringValue struct {
	Input string
	Value float32
}

func (sv *Float32StringValue) Set(input string) error {
	if input == sv.Input {
		return nil
	}

	v, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return fmt.Errorf("state.Float32StringValue.Set: failed to convert input to float32, error: %w", err)
	}
	sv.Value = float32(v)
	sv.Input = input

	return nil
}
