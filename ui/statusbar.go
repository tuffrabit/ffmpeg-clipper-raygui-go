package ui

import (
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/ffmpeg"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func getStatusText(appState *state.AppState) string {
	if appState.StatusText != "" {
		return appState.StatusText
	}

	statusText := ""
	if ffmpeg.SystemFFmpegHealth() {
		statusText = "ffmpeg, ffprobe, and ffplay found in system path!"
	} else {
		statusText = "ffmpeg, ffprobe, and ffplay NOT found in system path!"
	}

	return statusText
}

func Statusbar(appState *state.AppState) {
	statusText := getStatusText(appState)

	gui.StatusBar(rl.Rectangle{
		X:      components.XStart(),
		Y:      components.YFromEnd(20),
		Width:  components.FullWidth(),
		Height: 20,
	}, statusText)
}
