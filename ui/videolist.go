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

const (
	VIDEO_LIST_START_X         = MAIN_WIDTH_PADDING
	VIDEO_LIST_WIDTH   float32 = 500
	VIDEO_LIST_END_X           = VIDEO_LIST_START_X + VIDEO_LIST_WIDTH
	VIDEO_LIST_START_Y         = MAIN_HEIGHT_PADDING
	VIDEO_LIST_HEIGHT  float32 = 500
	VIDEO_LIST_END_Y           = VIDEO_LIST_START_Y + VIDEO_LIST_HEIGHT

	VIDEO_LIST_BUTTON_WIDTH   = (VIDEO_LIST_WIDTH - MAIN_WIDTH_PADDING) / 2
	VIDEO_LIST_BUTTON_START_Y = VIDEO_LIST_END_Y + MAIN_HEIGHT_PADDING
	VIDEO_LIST_BUTTON_END_Y   = VIDEO_LIST_BUTTON_START_Y + BUTTON_HEIGHT

	VIDEO_LIST_PLAY_BUTTON_START_X = VIDEO_LIST_START_X
	VIDEO_LIST_PLAY_BUTTON_END_X   = VIDEO_LIST_START_X + VIDEO_LIST_BUTTON_WIDTH

	VIDEO_LIST_DELETE_BUTTON_START_X = VIDEO_LIST_PLAY_BUTTON_END_X + MAIN_WIDTH_PADDING
	VIDEO_LIST_DELETE_BUTTON_END_X   = VIDEO_LIST_DELETE_BUTTON_START_X + PROFILE_LIST_BUTTON_WIDTH

	VIDEO_LIST_BUTTON_START_Y2 = VIDEO_LIST_BUTTON_END_Y + MAIN_HEIGHT_PADDING
	VIDEO_LIST_BUTTON_END_Y2   = VIDEO_LIST_BUTTON_START_Y2 + BUTTON_HEIGHT

	VIDEO_LIST_CLIP_START_BUTTON_START_X = VIDEO_LIST_START_X
	VIDEO_LIST_CLIP_START_BUTTON_END_X   = VIDEO_LIST_START_X + VIDEO_LIST_BUTTON_WIDTH

	VIDEO_LIST_CLIP_STOP_BUTTON_START_X = VIDEO_LIST_CLIP_START_BUTTON_END_X + MAIN_WIDTH_PADDING

	VIDEO_LIST_DELETE_CONFIRM_CONTEXT = "video_list_delete"
)

var (
	videoListRect                rl.Rectangle = rl.Rectangle{X: VIDEO_LIST_START_X, Y: VIDEO_LIST_START_Y, Width: VIDEO_LIST_WIDTH, Height: VIDEO_LIST_HEIGHT}
	videoListPlayButtonRect      rl.Rectangle = rl.Rectangle{X: VIDEO_LIST_PLAY_BUTTON_START_X, Y: VIDEO_LIST_BUTTON_START_Y, Width: VIDEO_LIST_BUTTON_WIDTH, Height: BUTTON_HEIGHT}
	videoListDeleteButtonRect    rl.Rectangle = rl.Rectangle{X: VIDEO_LIST_DELETE_BUTTON_START_X, Y: VIDEO_LIST_BUTTON_START_Y, Width: VIDEO_LIST_BUTTON_WIDTH, Height: BUTTON_HEIGHT}
	videoListClipStartButtonRect rl.Rectangle = rl.Rectangle{X: VIDEO_LIST_CLIP_START_BUTTON_START_X, Y: VIDEO_LIST_BUTTON_START_Y2, Width: VIDEO_LIST_BUTTON_WIDTH, Height: BUTTON_HEIGHT}
	videoListClipStopButtonRect  rl.Rectangle = rl.Rectangle{X: VIDEO_LIST_CLIP_STOP_BUTTON_START_X, Y: VIDEO_LIST_BUTTON_START_Y2, Width: VIDEO_LIST_BUTTON_WIDTH, Height: BUTTON_HEIGHT}
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
var playErrors chan error
var playing bool
var timestamp string

func populatePlayState(appState *state.AppState) {
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
			case err, ok := <-clipErrors:
				if !ok {
					clipping = false
					clipTimestamp = ""
					break timestamploop
				}

				if err != nil {
					clipping = false
					clipTimestamp = ""
					appState.GlobalMessageModalState.Init("FFmpeg Error", err.Error(), components.MESSAGE_MODAL_TYPE_ERROR)
					break timestamploop
				}
			default:
				break timestamploop
			}
		}
	}
}

func handleVideoSelect(appState *state.AppState) {
	if len(appState.VideoListState.EntryList) == 0 {
		return
	}

	file := appState.VideoListState.EntryList[appState.VideoListState.Active]

	if appState.CurrentVideoState.Name != file.Name() {
		appState.CurrentVideoState.Name = file.Name()
		appState.CurrentVideoState.FullPath = path.Join(appState.LocalDirectory, file.Name())

		width, height, err := system.GetVideoResolution(appState.CurrentVideoState.FullPath)
		if err != nil {
			appState.GlobalMessageModalState.Init("Video Resolution Error", fmt.Sprintf("Failed to calculate video resolution for %s, error: %v", appState.CurrentVideoState.FullPath, err), components.MESSAGE_MODAL_TYPE_ERROR)
		}

		appState.CurrentVideoState.Update(width, height)
	}
}

func handlePlayClick(clicked bool, appState *state.AppState) {
	if clicked {
		timestamps = make(chan string)
		playStates = make(chan bool)
		playErrors = make(chan error)
		go system.RunFFplay(appState.CurrentVideoState.FullPath, timestamps, playStates, playErrors)
	}
}

func handleVideoDeleteClick(clicked bool, appState *state.AppState) {
	if clicked {
		appState.GlobalConfirmModalState.Init("Delete Video?", fmt.Sprintf("Are you sure you want to delete %s?", appState.CurrentVideoState.FullPath), VIDEO_LIST_DELETE_CONFIRM_CONTEXT)
	}
}

func handleSetStartClick(clicked bool, appState *state.AppState) {
	if !clicked {
		return
	}

	err := appState.ClipState.SetStart(timestamp)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid Start", fmt.Sprintf("Failed to parse start value %s, error: %v", timestamp, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleSetStopClick(clicked bool, appState *state.AppState) {
	if !clicked {
		return
	}

	err := appState.ClipState.SetEnd(timestamp)
	if err != nil {
		appState.GlobalMessageModalState.Init("Invalid End", fmt.Sprintf("Failed to parse end value %s, error: %v", timestamp, err), components.MESSAGE_MODAL_TYPE_ERROR)
	}
}

func handleVideoDeleteGlobalConfirm(appState *state.AppState) {
	if appState.GlobalConfirmModalState.Completed(VIDEO_LIST_DELETE_CONFIRM_CONTEXT) {
		if appState.GlobalConfirmModalState.Result {
			err := os.Remove(appState.CurrentVideoState.FullPath)
			if err != nil {
				appState.GlobalMessageModalState.Init("Delete Video Error", fmt.Sprintf("Failed to delete %s, error: %v", appState.CurrentVideoState.FullPath, err), components.MESSAGE_MODAL_TYPE_ERROR)
			}
			appState.VideoListState.Reset()
		}

		appState.GlobalConfirmModalState.Reset()
	}
}

func VideoList(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	appState.VideoListState.InitFileList(appState.LocalDirectory, allowedExtensions)
	appState.VideoListState.Active = gui.ListView(videoListRect, appState.VideoListState.ListEntries, &appState.VideoListState.ScrollIndex, appState.VideoListState.Active)

	populatePlayState(appState)
	handleVideoSelect(appState)

	playButtonText := gui.IconText(gui.ICON_PLAYER_PLAY, "Play")
	if playing || appState.VideoListState.ListEntries == "" {
		gui.SetState(gui.STATE_DISABLED)
		playButtonText = timestamp
	}
	playButton := gui.Button(videoListPlayButtonRect, playButtonText)
	deleteButton := gui.Button(videoListDeleteButtonRect, gui.IconText(gui.ICON_FILE_DELETE, "Delete"))
	gui.SetState(gui.STATE_NORMAL)

	if !playing {
		gui.SetState(gui.STATE_DISABLED)
	}
	setStartButton := gui.Button(videoListClipStartButtonRect, gui.IconText(gui.ICON_PLAYER_PREVIOUS, "Set Clip Start"))
	setStopButton := gui.Button(videoListClipStopButtonRect, gui.IconText(gui.ICON_PLAYER_NEXT, "Set Clip End"))
	gui.SetState(gui.STATE_NORMAL)

	handlePlayClick(playButton, appState)
	handleVideoDeleteClick(deleteButton, appState)
	handleVideoDeleteGlobalConfirm(appState)
	handleSetStartClick(setStartButton, appState)
	handleSetStopClick(setStopButton, appState)

	return nil
}
