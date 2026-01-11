package encoder

import (
	"context"
	"crypto/rand"
	"fmt"
	"os/exec"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/config"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
)

var videoExtension string = ".mp4"

func GetClipCmd(video string, startTime string, endTime string, profileState state.ProfileState) (*exec.Cmd, context.CancelFunc) {
	switch profileState.Profile.Encoder.Name {
	case config.Libx264EncoderName:
		return InitClipLibx264Cmd(video, startTime, endTime, profileState)
	case config.Libx265EncoderName:
		return InitClipLibx265Cmd(video, startTime, endTime, profileState)
	case config.LibaomAv1EncoderName:
		return InitClipLibaomAv1Cmd(video, startTime, endTime, profileState)
	case config.NvencH264EncoderName:
		return InitClipNvencH264Cmd(video, startTime, endTime, profileState)
	case config.NvencHevcEncoderName:
		return InitClipNvencHevcCmd(video, startTime, endTime, profileState)
	case config.IntelH264EncoderName:
		return InitClipIntelH264Cmd(video, startTime, endTime, profileState)
	case config.IntelHevcEncoderName:
		return InitClipIntelHevcCmd(video, startTime, endTime, profileState)
	case config.IntelAv1EncoderName:
		return InitClipIntelAv1Cmd(video, startTime, endTime, profileState)
	}

	return nil, nil
}

func GetRandomString() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%X", b)
}
