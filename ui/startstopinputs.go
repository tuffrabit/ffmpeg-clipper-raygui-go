package ui

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
)

const (
	START_STOP_INPUTS_START_X = VIDEO_LIST_END_X + MAIN_WIDTH_PADDING
	START_STOP_INPUTS_WIDTH   = PROFILE_LIST_WIDTH
	START_STOP_INPUTS_END_X   = PROFILE_LIST_END_X
	START_STOP_INPUTS_START_Y = VIDEO_SIZE_ROW2_END_Y + MAIN_HEIGHT_PADDING

	START_STOP_INPUTS_START_LABEL_END_Y   = START_STOP_INPUTS_START_Y + LABEL_HEIGHT
	START_STOP_INPUTS_START_INPUT_START_Y = START_STOP_INPUTS_START_LABEL_END_Y + LABEL_Y_PADDING
	START_STOP_INPUTS_START_INPUT_END_Y   = START_STOP_INPUTS_START_INPUT_START_Y + TEXTBOX_HEIGHT

	START_STOP_INPUTS_END_LABEL_START_Y = START_STOP_INPUTS_START_INPUT_END_Y + MAIN_HEIGHT_PADDING
	START_STOP_INPUTS_END_LABEL_END_Y   = START_STOP_INPUTS_END_LABEL_START_Y + LABEL_HEIGHT
	START_STOP_INPUTS_END_INPUT_START_Y = START_STOP_INPUTS_END_LABEL_END_Y + LABEL_Y_PADDING
	START_STOP_INPUTS_END_INPUT_END_Y   = START_STOP_INPUTS_END_INPUT_START_Y + TEXTBOX_HEIGHT
)

var (
	startStopInputsStartLabelRect rl.Rectangle = rl.Rectangle{X: START_STOP_INPUTS_START_X, Y: START_STOP_INPUTS_START_Y, Width: START_STOP_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	startStopInputsStartInputRect rl.Rectangle = rl.Rectangle{X: START_STOP_INPUTS_START_X, Y: START_STOP_INPUTS_START_INPUT_START_Y, Width: START_STOP_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}
	startStopInputsEndLabelRect   rl.Rectangle = rl.Rectangle{X: START_STOP_INPUTS_START_X, Y: START_STOP_INPUTS_END_LABEL_START_Y, Width: START_STOP_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	startStopInputsEndInputRect   rl.Rectangle = rl.Rectangle{X: START_STOP_INPUTS_START_X, Y: START_STOP_INPUTS_END_INPUT_START_Y, Width: START_STOP_INPUTS_WIDTH, Height: TEXTBOX_HEIGHT}

	startEditMode bool = false
	endEditMode   bool = false
)

func StartStopInputs(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	startInput := appState.ClipState.StartInput
	gui.Label(startStopInputsStartLabelRect, "Start (seconds decimal)")
	if gui.TextBox(startStopInputsStartInputRect, &startInput, 9, startEditMode) {
		startEditMode = !startEditMode
	}
	err := appState.ClipState.SetStart(startInput)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Start", fmt.Sprintf("Failed to parse start value %s, error: %v", startInput, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}

	endInput := appState.ClipState.EndInput
	gui.Label(startStopInputsEndLabelRect, "End (seconds decimal)")
	if gui.TextBox(startStopInputsEndInputRect, &endInput, 9, endEditMode) {
		endEditMode = !endEditMode
	}
	err = appState.ClipState.SetEnd(endInput)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid End", fmt.Sprintf("Failed to parse end value %s, error: %v", endInput, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}

	return nil
}
