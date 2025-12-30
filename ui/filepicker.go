package ui

import (
	"fmt"
	"os"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FILE_PICKER_TYPE_BOTH int = iota
	FILE_PICKER_TYPE_DIRECTORY
	FILE_PICKER_TYPE_FILE
)

var listViewScrollIndex int32
var listViewActive int32
var listEntries string

func FilePicker() (bool, error) {
	rl.DrawRectangle(int32(XStart()), int32(YStart()), int32(XFromEnd(0)), int32(YFromEnd(0)), rl.Fade(rl.RayWhite, 0.8))
	wb := gui.WindowBox(rl.Rectangle{
		X:      20,
		Y:      20,
		Width:  XFromEnd(40),
		Height: YFromEnd(40),
	}, "Open")

	if wb {
		return wb, nil
	}

	if listEntries == "" {
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			return true, fmt.Errorf("failed to obtain home directory: %v", err)
		}

		directoryEntries, err := os.ReadDir(homeDirectory)
		if err != nil {
			return true, fmt.Errorf("failed to read home directory: %v", err)
		}

		for _, directoryEntry := range directoryEntries {
			name := directoryEntry.Name()

			if directoryEntry.IsDir() {
				name = name + " (DIR)"
			}

			listEntries = fmt.Sprintf("%s%s;", listEntries, name)
		}
	}

	gui.ListView(rl.Rectangle{
		X:      40,
		Y:      40,
		Width:  XFromEnd(80),
		Height: YFromEnd(80),
	}, listEntries, &listViewScrollIndex, listViewActive)

	return wb, nil
}
