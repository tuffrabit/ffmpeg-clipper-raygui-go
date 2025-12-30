package ui

import (
	"fmt"
	"os"
	"path"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/system"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var allowedExtensions []string = []string{
	".mp4",
	".mkv",
	".avi",
	".flv",
	".mov",
	".wmv",
	".ogg",
	".webm",
}

var timestamps chan string
var playStates chan bool
var playing bool
var timestamp string

func VideoList(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	appState.VideoListState.InitFileList(appState.LocalDirectory, allowedExtensions)

	vlRect := rl.Rectangle{X: components.START_PADDING, Y: components.START_PADDING, Width: 500, Height: 500}
	appState.VideoListState.Active = gui.ListView(vlRect, appState.VideoListState.ListEntries, &appState.VideoListState.ScrollIndex, appState.VideoListState.Active)

	var padding float32 = 20
	yStart := vlRect.Y + vlRect.Height + padding

	if timestamps != nil && playStates != nil {
	timestamploop:
		for {
			select {
			case time, ok := <-timestamps:
				playing = true
				timestamp = time

				if !ok {
					playing = false
					timestamp = ""
					break timestamploop
				}
			case playState, ok := <-playStates:
				if !playState || !ok {
					playing = false
					timestamp = ""
					break timestamploop
				}
			default:
				break timestamploop
			}
		}
	}

	playButtonText := "Play"
	if playing {
		gui.SetState(gui.STATE_DISABLED)
		playButtonText = timestamp
	}
	playButtonRect := rl.Rectangle{X: components.START_PADDING, Y: yStart, Width: 240, Height: components.BUTTON_HEIGHT}
	playButton := gui.Button(playButtonRect, gui.IconText(gui.ICON_PLAYER_PLAY, playButtonText))

	deleteButtonRectX := vlRect.X + vlRect.Width - 240
	deleteButtonRect := rl.Rectangle{X: deleteButtonRectX, Y: yStart, Width: 240, Height: components.BUTTON_HEIGHT}
	yStart = yStart + playButtonRect.Height + padding
	deleteButton := gui.Button(deleteButtonRect, gui.IconText(gui.ICON_FILE_DELETE, "Delete"))
	if playing {
		gui.SetState(gui.STATE_NORMAL)
	}

	if playButton {
		file := appState.VideoListState.EntryList[appState.VideoListState.Active]
		filepath := path.Join(appState.LocalDirectory, file.Name())

		timestamps = make(chan string)
		playStates = make(chan bool)
		go system.RunFFplay(filepath, timestamps, playStates)
	}

	if deleteButton {
		file := appState.VideoListState.EntryList[appState.VideoListState.Active]
		filepath := path.Join(appState.LocalDirectory, file.Name())
		appState.GlobalConfirmModalState.Init("Delete?", fmt.Sprintf("Are you sure you want to delete %s?", filepath))
	}

	if appState.GlobalConfirmModalState.Completed() {
		if appState.GlobalConfirmModalState.Result {
			file := appState.VideoListState.EntryList[appState.VideoListState.Active]
			filepath := path.Join(appState.LocalDirectory, file.Name())
			err := os.Remove(filepath)
			if err != nil {
				appState.GlobalMessageModalState.Init("Delete Error", fmt.Sprintf("Failed to delete %s, error: %v", filepath, err), components.MESSAGE_MODAL_TYPE_ERROR)
			}
			appState.VideoListState.Reset()
		}

		appState.GlobalConfirmModalState.Reset()
	}

	return nil
}
