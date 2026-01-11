package encoder

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/state"
	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/system"
)

func ClipLibx265(video string, startTime string, endTime string, profileState state.ProfileState) (string, error) {
	videoName := video[:len(video)-len(videoExtension)]
	newVideoName := fmt.Sprintf("%v_clip%v%v", videoName, GetRandomString(), videoExtension)

	cmd := exec.Command("ffmpeg",
		"-nostats",
		"-hide_banner",
		"-loglevel",
		"8",
		"-stats",
		fmt.Sprintf("%ss", startTime),
		"-to",
		fmt.Sprintf("%ss", endTime),
		"-i",
		video,
		"-c:v",
		"libx265",
		"-preset",
		profileState.Profile.EncoderSettings.Libx265.EncodingPreset,
		"-crf",
		profileState.Libx265EncodingQualityTargetInput,
		"-vf",
		fmt.Sprintf("scale=iw/%v:-1:flags=bicubic,exposure=%v:black=%v,eq=saturation=%v:contrast=%v:brightness=%v:gamma=%v",
			profileState.ScaleFactorInput,
			profileState.ExposureInput,
			profileState.BlackLevelInput,
			profileState.SaturationInput,
			profileState.ContrastInput,
			profileState.BrightnessInput,
			profileState.GammaInput,
		),
		newVideoName,
	)
	cmdOutput, err := system.RunSystemCommand(cmd)
	if err != nil {
		log.Printf("encoder.ClipLibx265: error running ffmpeg: %v\n", cmdOutput)
	}

	return newVideoName, nil
}

func InitClipLibx265Cmd(video string, startTime string, endTime string, profileState state.ProfileState) (*exec.Cmd, context.CancelFunc) {
	videoName := video[:len(video)-len(videoExtension)]
	newVideoName := fmt.Sprintf("%v_clip%v%v", videoName, GetRandomString(), videoExtension)

	ctx, cancel := context.WithCancel(context.Background())
	return exec.CommandContext(ctx,
		"ffmpeg",
		"-nostats",
		"-hide_banner",
		"-loglevel",
		"8",
		"-stats",
		"-ss",
		fmt.Sprintf("%ss", startTime),
		"-to",
		fmt.Sprintf("%ss", endTime),
		"-i",
		video,
		"-c:v",
		"libx265",
		"-preset",
		profileState.Profile.EncoderSettings.Libx265.EncodingPreset,
		"-crf",
		profileState.Libx265EncodingQualityTargetInput,
		"-vf",
		fmt.Sprintf("scale=iw/%v:-1:flags=bicubic,exposure=%v:black=%v,eq=saturation=%v:contrast=%v:brightness=%v:gamma=%v",
			profileState.ScaleFactorInput,
			profileState.ExposureInput,
			profileState.BlackLevelInput,
			profileState.SaturationInput,
			profileState.ContrastInput,
			profileState.BrightnessInput,
			profileState.GammaInput,
		),
		newVideoName,
	), cancel
}
