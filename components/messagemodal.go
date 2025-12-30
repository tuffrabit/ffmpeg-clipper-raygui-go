package components

import (
	gui "github.com/gen2brain/raylib-go/raygui"
)

type MessageModalType int

const (
	MESSAGE_MODAL_TYPE_INFO  MessageModalType = 0
	MESSAGE_MODAL_TYPE_ERROR MessageModalType = 1
)

type MessageModalState struct {
	WindowOpen bool
	Title      string
	Message    string
	Type       MessageModalType
}

func (mps *MessageModalState) Reset() {
	mps.WindowOpen = false
	mps.Title = ""
	mps.Message = ""
	mps.Type = MESSAGE_MODAL_TYPE_INFO
}

func (mps *MessageModalState) Init(title string, message string, messageModalType MessageModalType) {
	mps.Reset()
	mps.Title = title
	mps.Message = message
	mps.Type = messageModalType
	mps.WindowOpen = true
}

func MessageModal(state MessageModalState) (MessageModalState, error) {
	if state.WindowOpen == false {
		return state, nil
	}

	DrawFullBackground()
	mbRect := RectangleCenterEdgeOffset(20)

	title := state.Title

	if title == "" {
		switch state.Type {
		case MESSAGE_MODAL_TYPE_INFO:
			title = "Info"
		default:
			title = "Alert"
		}
	}

	result := gui.MessageBox(mbRect, gui.IconText(gui.ICON_EXIT, title), state.Message, "Ok")
	if result == 0 || result == 1 {
		state.WindowOpen = false
	}

	return state, nil
}
