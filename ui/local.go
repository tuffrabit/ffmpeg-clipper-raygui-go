package ui

import (
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PickLocalDirectory(appState *state.AppState) {
	if appState.LocalDirectory != "" {
		return
	}

	var buttonWidth float32 = 150
	halfButtonWidth := buttonWidth / 2
	halfButtonHeight := components.BUTTON_HEIGHT / 2

	gui.SetStyle(gui.BUTTON, gui.TEXT_ALIGNMENT, gui.TEXT_ALIGN_CENTER)
	if gui.Button(rl.Rectangle{X: components.XCenterWithOffset(halfButtonWidth), Y: components.YCenterWithOffset(halfButtonHeight), Width: buttonWidth, Height: components.BUTTON_HEIGHT}, gui.IconText(gui.ICON_FOLDER_OPEN, "Choose a directory")) {
		appState.LocalDirectoryPickerState.WindowOpen = true
	}
}
