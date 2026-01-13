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

func GetClipCmd(video string, startTime string, endTime string, profileState state.ProfileState) (string, *exec.Cmd, context.CancelFunc) {
	videoName := video[:len(video)-len(videoExtension)]
	newVideoName := fmt.Sprintf("%v_clip%v%v", videoName, GetRandomString(), videoExtension)
	var cmd *exec.Cmd
	var cancel context.CancelFunc

	switch profileState.Profile.Encoder.Name {
	case config.Libx264EncoderName:
		cmd, cancel = InitClipLibx264Cmd(video, newVideoName, startTime, endTime, profileState)
	case config.Libx265EncoderName:
		cmd, cancel = InitClipLibx265Cmd(video, newVideoName, startTime, endTime, profileState)
	case config.LibaomAv1EncoderName:
		cmd, cancel = InitClipLibaomAv1Cmd(video, newVideoName, startTime, endTime, profileState)
	case config.NvencH264EncoderName:
		cmd, cancel = InitClipNvencH264Cmd(video, newVideoName, startTime, endTime, profileState)
	case config.NvencHevcEncoderName:
		cmd, cancel = InitClipNvencHevcCmd(video, newVideoName, startTime, endTime, profileState)
	case config.IntelH264EncoderName:
		cmd, cancel = InitClipIntelH264Cmd(video, newVideoName, startTime, endTime, profileState)
	case config.IntelHevcEncoderName:
		cmd, cancel = InitClipIntelHevcCmd(video, newVideoName, startTime, endTime, profileState)
	case config.IntelAv1EncoderName:
		cmd, cancel = InitClipIntelAv1Cmd(video, newVideoName, startTime, endTime, profileState)
	default:
		return "", nil, nil
	}

	return newVideoName, cmd, cancel
}

func GetRandomString() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%X", b)
}
