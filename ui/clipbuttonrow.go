package ui

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/encoder"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/system"
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

	clipTimestamps chan string
	clipStates     chan bool
	clipErrors     chan error
	clipping       bool
	clipTimestamp  string
)

func populateClipState(appState *state.AppState) {
	if clipTimestamps != nil && clipStates != nil {
	cliptimestamploop:
		for {
			select {
			case time, ok := <-clipTimestamps:
				clipping = true
				clipTimestamp = time

				if !ok {
					clipping = false
					clipTimestamp = ""
					break cliptimestamploop
				}
			case clipState, ok := <-clipStates:
				if !ok {
					clipping = false
					clipTimestamp = ""
					break cliptimestamploop
				}

				if !clipState {
					appState.VideoListState.ResetWithSelection()
					clipping = false
					clipTimestamp = ""
					break cliptimestamploop
				}
			case err, ok := <-clipErrors:
				if !ok {
					clipping = false
					clipTimestamp = ""
					break cliptimestamploop
				}

				if err != nil {
					appState.VideoListState.ResetWithSelection()
					clipping = false
					clipTimestamp = ""
					appState.GlobalMessageModalState.Init("FFmpeg Error", err.Error(), components.MESSAGE_MODAL_TYPE_ERROR)
					break cliptimestamploop
				}
			default:
				break cliptimestamploop
			}
		}
	}
}

func handleClipClick(clicked bool, appState *state.AppState) {
	if !clicked {
		return
	}

	startTime := appState.ClipState.StartInput
	if startTime == "" {
		startTime = "0"
	}
	endTime := appState.ClipState.EndInput
	if endTime == "" {
		endTime = "0"
	}

	cmd, cancel := encoder.GetClipCmd(appState.CurrentVideoState.FullPath, startTime, endTime, appState.ProfileState)
	if cmd == nil {
		return
	}

	clipTimestamps = make(chan string)
	clipStates = make(chan bool)
	clipErrors = make(chan error)
	go system.RunClipCmd(cmd, cancel, clipTimestamps, clipStates, clipErrors)
}

func clipTimeValid(appState *state.AppState) bool {
	if appState.ClipState.EndInput == "" {
		return false
	}

	if appState.ClipState.EndSeconds <= appState.ClipState.StartSeconds {
		return false
	}

	return true
}

func ClipButtonRow(appState *state.AppState) error {
	if appState.LocalDirectory == "" {
		return nil
	}

	profile := appState.ProfileListState.SelectedProfile()
	appState.ProfileState.Profile.PlayAfter = gui.CheckBox(clipButtonRowPlayCheckRect, "Auto Play New Clip", appState.ProfileState.Profile.PlayAfter)
	appState.ProfileListState.UpdateSelectedProfile(profile)

	populateClipState(appState)

	clipButtonText := gui.IconText(gui.ICON_FILE_CUT, "Clip")
	if clipping {
		gui.SetState(gui.STATE_DISABLED)
		clipButtonText = clipTimestamp
	}
	if appState.VideoListState.ListEntries == "" || !clipTimeValid(appState) {
		gui.SetState(gui.STATE_DISABLED)
	}
	clipButton := gui.Button(clipButtonRowClipButtonRect, clipButtonText)
	gui.SetState(gui.STATE_NORMAL)

	handleClipClick(clipButton, appState)

	return nil
}
