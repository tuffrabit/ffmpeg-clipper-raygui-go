package state

import (
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
	ClipState                 ClipState
}
