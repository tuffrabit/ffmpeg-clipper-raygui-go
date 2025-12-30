package components

import (
	gui "github.com/gen2brain/raylib-go/raygui"
)

type ConfirmModalState struct {
	WindowOpen bool
	Title      string
	Message    string
	Confirmed  bool
	Result     bool
}

func (cps *ConfirmModalState) Reset() {
	cps.WindowOpen = false
	cps.Title = ""
	cps.Message = ""
	cps.Confirmed = false
	cps.Result = false
}

func (cps *ConfirmModalState) Init(title string, message string) {
	cps.Reset()
	cps.Title = title
	cps.Message = message
	cps.WindowOpen = true
}

func (cps *ConfirmModalState) Completed() bool {
	if !cps.WindowOpen && cps.Confirmed {
		return true
	}

	return false
}

func ConfirmModal(state ConfirmModalState) (ConfirmModalState, error) {
	if state.WindowOpen == false {
		return state, nil
	}

	DrawFullBackground()
	mbRect := RectangleCenterEdgeOffset(20)

	title := state.Title

	if title == "" {
		title = "Confirm?"
	}

	result := gui.MessageBox(mbRect, gui.IconText(gui.ICON_EXIT, title), state.Message, "Yes;No")
	switch result {
	case 0:
	case 2:
		state.WindowOpen = false
		state.Confirmed = true
		state.Result = false
	case 1:
		state.WindowOpen = false
		state.Confirmed = true
		state.Result = true
	}

	return state, nil
}
