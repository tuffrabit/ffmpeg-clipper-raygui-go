package components

import (
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/optional"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BUTTON_HEIGHT      float32 = 50
	LABEL_HEIGHT       float32 = 24
	TEXTBOX_HEIGHT     float32 = 24
	CHECKBOX_HEIGHT    float32 = 24
	DROPDOWNBOX_HEIGHT float32 = 24
)

var xCenter optional.Optional[float32]
var yCenter optional.Optional[float32]

func XStart() float32 {
	return 0
}

func XCenter() float32 {
	if xCenter.IsSet() {
		return xCenter.Get()
	}

	xCenter.Set(float32(rl.GetScreenWidth()) / 2)
	return xCenter.Get()
}

func XCenterWithOffset(offset float32) float32 {
	return XCenter() - offset
}

func XFromEnd(offset float32) float32 {
	return float32(rl.GetScreenWidth()) - offset
}

func YStart() float32 {
	return 0
}

func YCenter() float32 {
	if yCenter.IsSet() {
		return yCenter.Get()
	}

	yCenter.Set(float32(rl.GetScreenHeight()) / 2)
	return yCenter.Get()
}

func YCenterWithOffset(offset float32) float32 {
	return YCenter() - offset
}

func YFromEnd(offset float32) float32 {
	return float32(rl.GetScreenHeight()) - offset
}

func FullWidth() float32 {
	return float32(rl.GetScreenWidth())
}

func FullHeight() float32 {
	return float32(rl.GetScreenHeight())
}

func DrawFullBackground() {
	rl.DrawRectangle(0, 0, int32(XFromEnd(0)), int32(YFromEnd(0)), rl.Fade(rl.RayWhite, 0.8))
}

func RectangleCenterEdgeOffsets(xOffset float32, yOffset float32) rl.Rectangle {
	return rl.Rectangle{
		X:      xOffset,
		Y:      yOffset,
		Width:  XFromEnd(xOffset * 2),
		Height: YFromEnd(yOffset * 2),
	}
}

func RectangleCenterEdgeOffset(offset float32) rl.Rectangle {
	return RectangleCenterEdgeOffsets(offset, offset)
}

func MeasureText(text string) float32 {
	currentFont := gui.GetFont()
	textSize := gui.GetStyle(gui.DEFAULT, gui.TEXT_SIZE)
	textSpacing := gui.GetStyle(gui.DEFAULT, gui.TEXT_SPACING)
	textDimensions := rl.MeasureTextEx(currentFont, text, float32(textSize), float32(textSpacing))

	return textDimensions.X
}
