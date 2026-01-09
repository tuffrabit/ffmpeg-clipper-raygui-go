package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
)

const (
	CLIP_BUTTON_ROW_START_X = VIDEO_LIST_END_X + MAIN_WIDTH_PADDING
	CLIP_BUTTON_ROW_WIDTH   = PROFILE_LIST_WIDTH
	CLIP_BUTTON_ROW_END_X   = PROFILE_LIST_END_X
	CLIP_BUTTON_ROW_START_Y = START_STOP_INPUTS_END_INPUT_END_Y + MAIN_HEIGHT_PADDING

	CLIP_BUTTON_ROW_PLAY_CHECK_END_X   = CLIP_BUTTON_ROW_START_X + CHECKBOX_HEIGHT
	CLIP_BUTTON_ROW_PLAY_CHECK_START_Y = CLIP_BUTTON_ROW_START_Y + ((BUTTON_HEIGHT / 2) - (CHECKBOX_HEIGHT / 2))
	CLIP_BUTTON_ROW_PLAY_CHECK_END_Y   = CLIP_BUTTON_ROW_PLAY_CHECK_START_Y + CHECKBOX_HEIGHT

	CLIP_BUTTON_ROW_CLIP_BUTTON_WIDTH   = 210
	CLIP_BUTTON_ROW_CLIP_BUTTON_START_X = CLIP_BUTTON_ROW_END_X - CLIP_BUTTON_ROW_CLIP_BUTTON_WIDTH
	CLIP_BUTTON_ROW_CLIP_BUTTON_START_Y = CLIP_BUTTON_ROW_START_Y
)

var (
	clipButtonRowPlayCheckRect  rl.Rectangle = rl.Rectangle{X: CLIP_BUTTON_ROW_START_X, Y: CLIP_BUTTON_ROW_PLAY_CHECK_START_Y, Width: CHECKBOX_HEIGHT, Height: CHECKBOX_HEIGHT}
	clipButtonRowClipButtonRect rl.Rectangle = rl.Rectangle{X: CLIP_BUTTON_ROW_CLIP_BUTTON_START_X, Y: CLIP_BUTTON_ROW_CLIP_BUTTON_START_Y, Width: CLIP_BUTTON_ROW_CLIP_BUTTON_WIDTH, Height: BUTTON_HEIGHT}
)

func ClipButtonRow(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	profile := appState.ProfileListState.SelectedProfile()
	appState.ProfileState.Profile.PlayAfter = gui.CheckBox(clipButtonRowPlayCheckRect, "Auto Play New Clip", appState.ProfileState.Profile.PlayAfter)
	appState.ProfileListState.UpdateSelectedProfile(profile)
	gui.Button(clipButtonRowClipButtonRect, gui.IconText(gui.ICON_FILE_CUT, "Clip"))

	return nil
}
