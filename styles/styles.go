package styles

import (
	"embed"
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
)

//go:embed *.rgs
var folder embed.FS

func GetStyleBytes(name string) ([]byte, error) {
	bytes, err := folder.ReadFile(name + ".rgs")
	if err != nil {
		return nil, fmt.Errorf("styles.GetStyleBytes: failed to get read file, error: %w", err)
	}

	return bytes, nil
}

func LoadStyle(name string) {
	if name == "default" {
		gui.LoadStyleDefault()
		return
	}

	styleBytes, err := GetStyleBytes(name)
	if err != nil {
		gui.LoadStyleDefault()
		return
	}

	gui.LoadStyleFromMemory(styleBytes)
}
