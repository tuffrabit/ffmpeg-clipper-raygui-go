package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
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

	gui.Label(startStopInputsStartLabelRect, "Start (hh:mm:ss)")
	if gui.TextBox(startStopInputsStartInputRect, &appState.ClipState.Start, 9, startEditMode) {
		startEditMode = !startEditMode
	}

	gui.Label(startStopInputsEndLabelRect, "End (hh:mm:ss)")
	if gui.TextBox(startStopInputsEndInputRect, &appState.ClipState.End, 9, endEditMode) {
		endEditMode = !endEditMode
	}

	return nil
}
