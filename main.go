package main

import (
	"log"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/ffmpeg"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1280, 720, "ffmpeg clipper raygui")
	defer rl.CloseWindow()

	appState := state.AppState{}
	rl.SetTargetFPS(30)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		ui.PickLocalDirectory(&appState)
		ui.VideoList(&appState)
		ui.Statusbar(&appState)

		if appState.LocalDirectoryPickerState.WindowOpen {
			var err error
			appState.LocalDirectoryPickerState, err = components.FilePicker(appState.LocalDirectoryPickerState)
			if err != nil {
				log.Println("failed to open file picker:", err)
				appState.LocalDirectoryPickerState.Reset()
				appState.GlobalMessageModalState.Init("File Picker Error", err.Error(), components.MESSAGE_MODAL_TYPE_ERROR)
			}
			if appState.LocalDirectoryPickerState.Path != "" {
				if ffmpeg.TotalFFmpegHealth(appState.LocalDirectoryPickerState.Path) {
					appState.LocalDirectory = appState.LocalDirectoryPickerState.Path
					appState.StatusText = appState.LocalDirectory
					appState.LocalDirectoryPickerState.Reset()
				} else {
					appState.LocalDirectoryPickerState.Reset()
					appState.GlobalMessageModalState.Init("Directory Error", "ffmpeg, ffprobe, and ffplay NOT found in system or selected path!", components.MESSAGE_MODAL_TYPE_ERROR)
				}
			}
		}

		if appState.GlobalMessageModalState.WindowOpen {
			var err error
			appState.GlobalMessageModalState, err = components.MessageModal(appState.GlobalMessageModalState)
			if err != nil {
				log.Println("global message modal failed", err)
				appState.GlobalMessageModalState.Reset()
			}
		}

		if appState.GlobalConfirmModalState.WindowOpen {
			var err error
			appState.GlobalConfirmModalState, err = components.ConfirmModal(appState.GlobalConfirmModalState)
			if err != nil {
				log.Println("global confirm modal failed", err)
				appState.GlobalConfirmModalState.Reset()
			}
		}

		rl.EndDrawing()
	}
}
