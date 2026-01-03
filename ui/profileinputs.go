package ui

import (
	"fmt"
	"strings"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/config"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
)

const (
	PROFILE_INPUTS_START_X = PROFILE_LIST_END_X + MAIN_WIDTH_PADDING
	PROFILE_INPUTS_WIDTH   = float32(WINDOW_WIDTH) - PROFILE_LIST_END_X - (MAIN_WIDTH_PADDING * 2)
	PROFILE_INPUTS_END_X   = PROFILE_INPUTS_START_X + PROFILE_INPUTS_WIDTH
	PROFILE_INPUTS_START_Y = MAIN_HEIGHT_PADDING

	PROFILE_INPUTS_SCALE_DOWN_LABEL_END_Y   = PROFILE_INPUTS_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_SCALE_DOWN_INPUT_START_Y = PROFILE_INPUTS_SCALE_DOWN_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_SCALE_DOWN_INPUT_END_Y   = PROFILE_INPUTS_SCALE_DOWN_INPUT_START_Y + TEXTBOX_HEIGHT

	PROFILE_INPUTS_ENCODERS_LABEL_START_Y = PROFILE_INPUTS_SCALE_DOWN_INPUT_END_Y + MAIN_HEIGHT_PADDING
	PROFILE_INPUTS_ENCODERS_LABEL_END_Y   = PROFILE_INPUTS_ENCODERS_LABEL_START_Y + LABEL_HEIGHT
	PROFILE_INPUTS_ENCODERS_INPUT_START_Y = PROFILE_INPUTS_ENCODERS_LABEL_END_Y + LABEL_Y_PADDING
	PROFILE_INPUTS_ENCODERS_INPUT_END_Y   = PROFILE_INPUTS_ENCODERS_INPUT_START_Y + DROPDOWNBOX_HEIGHT
)

var (
	profileInputsScaleDownLabelRect rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsScaleDownInputRect rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_SCALE_DOWN_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsEncodersLabelRect  rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_ENCODERS_LABEL_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}
	profileInputsEncodersInputRect  rl.Rectangle = rl.Rectangle{X: PROFILE_INPUTS_START_X, Y: PROFILE_INPUTS_ENCODERS_INPUT_START_Y, Width: PROFILE_INPUTS_WIDTH, Height: LABEL_HEIGHT}

	scaleFactorEditMode bool   = false
	encoderList         string = EncoderList()
	encoderListActive   int32  = 0
	encoderListEditMode bool   = false
)

func EncoderList() string {
	encoderTypes := config.GetEncoderTypes()
	encoderList := ""

	for _, encoderName := range encoderTypes {
		encoderList = fmt.Sprintf("%s%s;", encoderList, encoderName)
	}

	encoderList = strings.TrimSuffix(encoderList, ";")

	return encoderList
}

func ProfileInputs(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	profile := appState.ProfileListState.SelectedProfile()
	gui.Label(profileInputsScaleDownLabelRect, "Scale Down Factor")
	scaleFactor := fmt.Sprintf("%f", profile.ScaleFactor)
	if gui.TextBox(profileInputsScaleDownInputRect, &scaleFactor, 30, scaleFactorEditMode) {
		scaleFactorEditMode = !scaleFactorEditMode
	}

	gui.Label(profileInputsEncodersLabelRect, "Encoder")
	if gui.DropdownBox(profileInputsEncodersInputRect, encoderList, &encoderListActive, encoderListEditMode) {
		encoderListEditMode = !encoderListEditMode
	}

	appState.ProfileListState.UpdateSelectedProfile(profile)

	return nil
}
