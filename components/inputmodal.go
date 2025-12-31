package components

import (
	gui "github.com/gen2brain/raylib-go/raygui"
)

type InputModalState struct {
	WindowOpen bool
	Title      string
	Message    string
	Confirmed  bool
	Result     string
	Context    string
}

func (ips *InputModalState) Reset() {
	ips.WindowOpen = false
	ips.Title = ""
	ips.Message = ""
	ips.Confirmed = false
	ips.Result = ""
	ips.Context = ""
}

func (ips *InputModalState) Init(title string, message string, context string) {
	ips.Reset()
	ips.Title = title
	ips.Message = message
	ips.Context = context
	ips.WindowOpen = true
}

func (ips *InputModalState) Completed(context string) bool {
	if ips.Context != context {
		return false
	}

	if !ips.WindowOpen && ips.Confirmed {
		return true
	}

	return false
}

func InputModal(state InputModalState) (InputModalState, error) {
	if state.WindowOpen == false {
		return state, nil
	}

	DrawFullBackground()

	title := state.Title

	if title == "" {
		title = "Input"
	}

	mbRect := RectangleCenterEdgeOffset(20)
	secret := true
	result := gui.TextInputBox(mbRect, title, state.Message, "Ok;Cancel", &state.Result, 255, &secret)

	switch result {
	case 0:
	case 2:
		state.WindowOpen = false
		state.Confirmed = true
		state.Result = ""
	case 1:
		state.WindowOpen = false
		state.Confirmed = true
	}

	return state, nil
}
