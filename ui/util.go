package ui

import (
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/optional"

	rl "github.com/gen2brain/raylib-go/raylib"
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
