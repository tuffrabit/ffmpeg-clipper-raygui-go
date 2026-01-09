package ui

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
)

const (
	VIDEO_SIZE_STATS_START_X = VIDEO_LIST_END_X + MAIN_WIDTH_PADDING
	VIDEO_SIZE_STATS_WIDTH   = PROFILE_LIST_WIDTH
	VIDEO_SIZE_STATS_END_X   = PROFILE_LIST_END_X
	VIDEO_SIZE_STATS_START_Y = PROFILE_LIST_BUTTON_END_Y + MAIN_HEIGHT_PADDING

	VIDEO_SIZE_ROW1_END_Y   = VIDEO_SIZE_STATS_START_Y + LABEL_HEIGHT
	VIDEO_SIZE_ROW2_START_Y = VIDEO_SIZE_ROW1_END_Y + LABEL_Y_PADDING
	VIDEO_SIZE_ROW2_END_Y   = VIDEO_SIZE_ROW2_START_Y + LABEL_HEIGHT
)

var (
	videoSizeSourceRect  rl.Rectangle = rl.Rectangle{X: VIDEO_SIZE_STATS_START_X, Y: VIDEO_SIZE_STATS_START_Y, Width: VIDEO_SIZE_STATS_WIDTH, Height: LABEL_HEIGHT}
	videoSizeNewSizeRect rl.Rectangle = rl.Rectangle{X: VIDEO_SIZE_STATS_START_X, Y: VIDEO_SIZE_ROW2_START_Y, Width: VIDEO_SIZE_STATS_WIDTH, Height: LABEL_HEIGHT}
)

func VideoSizeStats(appState *state.AppState) error {
	if appState.LocalDirectory == "" || appState.CurrentVideoState.Name == "" {
		return nil
	}

	gui.Label(videoSizeSourceRect, fmt.Sprintf("Source Resolution: %dx%d", appState.CurrentVideoState.Width, appState.CurrentVideoState.Height))
	gui.Label(videoSizeNewSizeRect, fmt.Sprintf("New Resolution:      %dx%d", appState.CurrentVideoState.NewWidth, appState.CurrentVideoState.NewHeight))

	return nil
}
