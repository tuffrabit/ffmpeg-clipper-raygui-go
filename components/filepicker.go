package components

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FilePickerType int

const (
	FILE_PICKER_TYPE_BOTH      FilePickerType = 0
	FILE_PICKER_TYPE_DIRECTORY FilePickerType = 1
	FILE_PICKER_TYPE_FILE      FilePickerType = 2
)

type FilePickerState struct {
	WindowOpen              bool
	Path                    string
	currentDirectory        string
	fileListViewScrollIndex int32
	fileListViewActive      int32
	fileList                []os.DirEntry
	fileListEntries         string
}

func (fps *FilePickerState) Reset() {
	fps.WindowOpen = false
	fps.Path = ""
	fps.currentDirectory = ""
	fps.fileListViewScrollIndex = 0
	fps.fileListViewActive = 0
	fps.fileList = nil
	fps.fileListEntries = ""
}

func (fps *FilePickerState) initFileList() error {
	if len(fps.fileList) > 0 {
		return nil
	}

	if fps.currentDirectory == "" {
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			fps.WindowOpen = false
			return fmt.Errorf("failed to obtain home directory: %v", err)
		}
		fps.currentDirectory = homeDirectory
	}

	directoryEntries, err := os.ReadDir(fps.currentDirectory)
	if err != nil {
		fps.WindowOpen = false
		return fmt.Errorf("failed to read home directory: %v", err)
	}

	fps.fileListEntries = ".;..;"

	for _, directoryEntry := range directoryEntries {
		if !directoryEntry.IsDir() || directoryEntry.Name() == "" {
			continue
		}

		fps.fileList = append(fps.fileList, directoryEntry)
		fps.fileListEntries = fmt.Sprintf("%s%s;", fps.fileListEntries, directoryEntry.Name())
	}

	fps.fileListEntries = strings.TrimSuffix(fps.fileListEntries, ";")

	return nil
}

func FilePicker(state FilePickerState) (FilePickerState, error) {
	DrawFullBackground()
	wbRect := RectangleCenterEdgeOffset(20)
	state.WindowOpen = !gui.WindowBox(wbRect, "Open")

	if !state.WindowOpen {
		return state, nil
	}

	var padding float32 = 20
	wbRectXEnd := wbRect.X + wbRect.Width
	yStart := wbRect.Y + padding + 24

	currentDirectoryLabelText := fmt.Sprintf("Current Directory: %s", state.currentDirectory)
	currentDirectoryLabelRect := rl.Rectangle{X: wbRect.X + padding, Y: yStart, Width: MeasureText(currentDirectoryLabelText), Height: 20}
	gui.Label(currentDirectoryLabelRect, currentDirectoryLabelText)

	openButtonRect := rl.Rectangle{X: wbRectXEnd - (150 + padding), Y: yStart, Width: 150, Height: BUTTON_HEIGHT}
	yStart = yStart + openButtonRect.Height + padding
	openButton := gui.Button(openButtonRect, gui.IconText(gui.ICON_FOLDER_OPEN, "Open"))
	if openButton && state.currentDirectory != "" {
		state.Path = state.currentDirectory
		return state, nil
	}

	state.initFileList()
	lvRect := rl.Rectangle{X: wbRect.X, Y: yStart, Width: wbRect.Width, Height: wbRect.Height - yStart + padding}
	state.fileListViewActive = gui.ListView(lvRect, state.fileListEntries, &state.fileListViewScrollIndex, state.fileListViewActive)
	if state.fileListViewActive == 0 {
		return state, nil
	}

	if state.fileListViewActive == 1 {
		state.currentDirectory = filepath.Join(state.currentDirectory, "..")
		state.fileListViewActive = 0
		state.fileList = state.fileList[:0]
		return state, nil
	}

	state.currentDirectory = filepath.Join(state.currentDirectory, state.fileList[state.fileListViewActive-2].Name())
	state.fileListViewActive = 0
	state.fileList = state.fileList[:0]

	return state, nil
}
