package ui

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/config"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/styles"
)

const (
	STYLE_BUTTON_START_X         = PROFILE_LIST_END_X + MAIN_WIDTH_PADDING
	STYLE_BUTTON_WIDTH   float32 = 100
	STYLE_BUTTON_END_X           = STYLE_BUTTON_START_X + STYLE_BUTTON_WIDTH
	STYLE_BUTTON_START_Y         = PROFILE_INPUTS_BLACK_LEVEL_INPUT_END_Y + MAIN_HEIGHT_PADDING
)

var (
	styleButtonRect rl.Rectangle = rl.Rectangle{X: STYLE_BUTTON_START_X, Y: STYLE_BUTTON_START_Y, Width: STYLE_BUTTON_WIDTH, Height: BUTTON_HEIGHT}

	styleList = []string{
		"default",
		"dark",
		"amber",
		"ashes",
		"bluish",
		"candy",
		"cherry",
		"cyber",
		"enefete",
		"genesis",
		"jungle",
		"lavanda",
		"rltech",
		"sunny",
		"terminal",
	}
)

func handleStyleClick(clicked bool, appState *state.AppState) {
	if !clicked {
		return
	}

	prevStyle := ""
	selected := ""
	for _, style := range styleList {
		if prevStyle != config.GetConfig().Style {
			prevStyle = style
			continue
		}

		selected = style
		break
	}

	if selected == "" {
		selected = "default"
	}

	styles.LoadStyle(selected)
	config.GetConfig().Style = selected
	err := config.SaveStyle(selected)
	if err != nil {
		appState.GlobalMessageModalState.Init("Update Style Error", fmt.Sprintf("Failed to update config, error: %v", err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func StyleButton(appState *state.AppState) error {
	styleButton := gui.Button(styleButtonRect, config.GetConfig().Style)
	handleStyleClick(styleButton, appState)

	return nil
}
