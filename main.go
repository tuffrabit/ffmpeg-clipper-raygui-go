package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/config"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/ffmpeg"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/styles"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/ui"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	logFile := setupLog()
	if logFile != nil {
		defer logFile.Close()
	}

	configLoadErr := config.LoadConfig()

	rl.InitWindow(ui.WINDOW_WIDTH, ui.WINDOW_HEIGHT, "ffmpeg clipper raygui")
	defer rl.CloseWindow()

	appState := state.CreateAppState()
	if configLoadErr != nil {
		appState.StatusText = fmt.Sprintf("failed to load config, error: %v", configLoadErr)
		log.Println(appState.StatusText)
	}

	rl.SetTargetFPS(30)
	styles.LoadStyle(config.GetConfig().Style)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(uint(gui.GetStyle(gui.DEFAULT, gui.BACKGROUND_COLOR))))

		if configLoadErr != nil {
			rl.EndDrawing()
			continue
		}

		ui.PickLocalDirectory(appState)
		ui.VideoList(appState)
		ui.FFplayHelp(appState)
		ui.ProfileList(appState)
		ui.VideoSizeStats(appState)
		ui.StartStopInputs(appState)
		ui.ClipButtonRow(appState)
		ui.ProfileInputs(appState)
		ui.StyleButton(appState)
		ui.Statusbar(appState)

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

		if appState.GlobalInputModalState.WindowOpen {
			var err error
			appState.GlobalInputModalState, err = components.InputModal(appState.GlobalInputModalState)
			if err != nil {
				log.Println("global input modal failed", err)
				appState.GlobalInputModalState.Reset()
			}
		}

		rl.EndDrawing()
	}
}

func setupLog() *os.File {
	log.SetOutput(os.Stdout)

	useFile := true
	logFile, err := os.OpenFile("ffmpeg-clipper.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		useFile = false
		log.Println("could not open ffmpeg-clipper.log file for write")
	}

	if useFile {
		multi := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multi)
		return logFile
	}

	return nil
}
