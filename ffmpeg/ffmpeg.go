package ffmpeg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/tuffrabit/ffmpeg-clipper-raygui-go/optional"
)

var systemExists optional.Optional[Exists]

type Paths struct {
	FfmpegPath  string
	FfprobePath string
	FfplayPath  string
}

type Exists struct {
	FFmpegExists  bool
	FFprobeExists bool
	FFplayExists  bool
}

func CheckSystemFFmpegCommand(command string) bool {
	cmd := exec.Command(command)
	var cmdOut bytes.Buffer
	var cmdErr bytes.Buffer
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr
	cmd.Run()
	outString := cmdOut.String()
	errString := cmdErr.String()

	if strings.HasPrefix(outString, command) {
		return true
	}

	if strings.HasPrefix(errString, command) {
		return true
	}

	return false
}

func CheckSystemFFmpeg() Exists {
	if systemExists.IsSet() {
		return systemExists.Get()
	}

	// pathVar := os.Getenv("PATH")
	// paths := strings.Split(pathVar, ";")

	// for _, pathEntry := range paths {
	// 	if strings.Contains(pathEntry, "ffmpeg") {
	// 		systemPaths := Paths{
	// 			FfmpegPath:  path.Join(pathEntry, "ffmpeg"),
	// 			FfprobePath: path.Join(pathEntry, "ffprobe"),
	// 			FfplayPath:  path.Join(pathEntry, "ffplay"),
	// 		}

	// 		systemExists.Set(CheckFFmpeg(systemPaths))
	// 		return systemExists.Get()
	// 	}
	// }

	exists := Exists{
		FFmpegExists:  CheckSystemFFmpegCommand("ffmpeg"),
		FFprobeExists: CheckSystemFFmpegCommand("ffprobe"),
		FFplayExists:  CheckSystemFFmpegCommand("ffplay"),
	}

	systemExists.Set(exists)
	return systemExists.Get()

	// systemExists.Set(Exists{
	// 	FFmpegExists:  false,
	// 	FFprobeExists: false,
	// 	FFplayExists:  false,
	// })
	// return systemExists.Get()
}

func CheckLocalFFmpeg(path string) Exists {
	pathSeparator := string(os.PathSeparator)
	systemPaths := Paths{
		FfmpegPath:  fmt.Sprintf("%v%vffmpeg.exe", path, pathSeparator),
		FfprobePath: fmt.Sprintf("%v%vffprobe.exe", path, pathSeparator),
		FfplayPath:  fmt.Sprintf("%v%vffplay.exe", path, pathSeparator),
	}

	return CheckFFmpeg(systemPaths)
}

func CheckFFmpeg(paths Paths) Exists {
	result := Exists{
		FFmpegExists:  false,
		FFprobeExists: false,
		FFplayExists:  false,
	}

	matches, err := filepath.Glob(fmt.Sprintf("%s*", paths.FfmpegPath))
	if err == nil && len(matches) > 0 {
		result.FFmpegExists = true
	}

	matches, err = filepath.Glob(fmt.Sprintf("%s*", paths.FfprobePath))
	if err == nil && len(matches) > 0 {
		result.FFprobeExists = true
	}

	matches, err = filepath.Glob(fmt.Sprintf("%s*", paths.FfplayPath))
	if err == nil && len(matches) > 0 {
		result.FFplayExists = true
	}

	return result
}

func SystemFFmpegHealth() bool {
	exists := CheckSystemFFmpeg()
	return FFmpegHealth(exists)
}

func LocalFFmpegHealth(path string) bool {
	exists := CheckLocalFFmpeg(path)
	return FFmpegHealth(exists)
}

func TotalFFmpegHealth(path string) bool {
	if SystemFFmpegHealth() {
		return true
	}

	return LocalFFmpegHealth(path)
}

func FFmpegHealth(exists Exists) bool {
	if exists.FFmpegExists && exists.FFprobeExists && exists.FFplayExists {
		return true
	}

	return false
}
