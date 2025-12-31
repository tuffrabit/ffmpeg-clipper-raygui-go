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
	Context    string
}

func (cps *ConfirmModalState) Reset() {
	cps.WindowOpen = false
	cps.Title = ""
	cps.Message = ""
	cps.Confirmed = false
	cps.Result = false
	cps.Context = ""
}

func (cps *ConfirmModalState) Init(title string, message string, context string) {
	cps.Reset()
	cps.Title = title
	cps.Message = message
	cps.Context = context
	cps.WindowOpen = true
}

func (cps *ConfirmModalState) Completed(context string) bool {
	if cps.Context != context {
		return false
	}

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
