package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
)

const (
	FFPLAY_HELP_START_X = MAIN_WIDTH_PADDING
	FFPLAY_HELP_WIDTH   = VIDEO_LIST_WIDTH
	FFPLAY_HELP_END_X   = FFPLAY_HELP_START_X + FFPLAY_HELP_WIDTH

	FFPLAY_HELP_PAUSE_LABEL_START_Y = VIDEO_LIST_BUTTON_END_Y2 + CLOSE_LABEL_HEIGHT_PADDING
	FFPLAY_HELP_PAUSE_LABEL_END_Y   = FFPLAY_HELP_PAUSE_LABEL_START_Y + LABEL_HEIGHT

	FFPLAY_HELP_STEP_LABEL_START_Y = FFPLAY_HELP_PAUSE_LABEL_END_Y + CLOSE_LABEL_HEIGHT_PADDING
	FFPLAY_HELP_STEP_LABEL_END_Y   = FFPLAY_HELP_STEP_LABEL_START_Y + LABEL_HEIGHT
)

var (
	ffplayPauseLabelRect rl.Rectangle = rl.Rectangle{X: FFPLAY_HELP_START_X, Y: FFPLAY_HELP_PAUSE_LABEL_START_Y, Width: FFPLAY_HELP_WIDTH, Height: LABEL_HEIGHT}
	ffplayStepLabelRect  rl.Rectangle = rl.Rectangle{X: FFPLAY_HELP_START_X, Y: FFPLAY_HELP_STEP_LABEL_START_Y, Width: FFPLAY_HELP_WIDTH, Height: LABEL_HEIGHT}
)

func FFplayHelp(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	gui.Label(ffplayPauseLabelRect, "Play/Pause: p, SPC              Seek backward/forward 10 seconds: left/right")
	gui.Label(ffplayStepLabelRect, "Step to the next frame: s      Seek backward/forward 1 minute: down/up")

	return nil
}
