package state

import (
	"fmt"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/components"
)

type AppState struct {
	StatusText                string
	LocalDirectory            string
	GlobalMessageModalState   components.MessageModalState
	GlobalConfirmModalState   components.ConfirmModalState
	GlobalInputModalState     components.InputModalState
	LocalDirectoryPickerState components.FilePickerState
	CurrentVideoState         CurrentVideoState
	VideoListState            DirEntryListState
	ProfileListState          ProfileListState
	ProfileState              ProfileState
	ClipState                 ClipState
	EncoderPresetsState       map[string]EncoderPresetState
}

func CreateAppState() AppState {
	appState := AppState{}
	return appState
}

func (as *AppState) GetEncoderPresetState(encoderTypeName string) (EncoderPresetState, error) {
	for t, encoderPresetState := range as.EncoderPresetsState {
		if t == encoderTypeName {
			return encoderPresetState, nil
		}
	}

	return EncoderPresetState{}, fmt.Errorf("state.AppState.GetEncoderPresetState: encoder preset state for type %s does not exist", encoderTypeName)
}
