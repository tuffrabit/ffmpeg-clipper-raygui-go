package system

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

func RunFFplay(videopath string, timestamps chan string, playStates chan bool) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := exec.CommandContext(ctx, "ffplay", "-y", "680", "-loglevel", "-8", "-stats", videopath)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("could not get stderr pipe: %w", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	go func() {
		prevScan := ""
		capture := false

		for scanner.Scan() {
			currentScan := scanner.Text()

			if prevScan == "vq=" {
				if currentScan == "0KB" {
					capture = false
				} else {
					capture = true
				}
			}

			if capture && (currentScan == "A-V:" || currentScan == "M-V:" || currentScan == "M-A:") && prevScan != "nan" {
				timestamps <- prevScan
			}
			prevScan = currentScan
		}
		wg.Done()
	}()

	if err = cmd.Start(); err != nil {
		return err
	}
	playStates <- true

	wg.Wait()
	playStates <- false
	close(playStates)
	close(timestamps)
	return cmd.Wait()
}

func GetVideoResolution(videopath string) (int, int, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v",
		"error",
		"-select_streams",
		"v:0",
		"-show_entries",
		"stream=width,height",
		"-of",
		"csv=s=x:p=0",
		videopath,
	)
	output, err := RunSystemCommand(cmd)
	if err != nil {
		return 0, 0, fmt.Errorf("system.GetVideoResolution: ffprobe failed\nstderr: %v\nerr: %w", output, err)
	}

	output = strings.TrimSuffix(output, "\r\n")
	output = strings.TrimSuffix(output, "\n")
	output = strings.TrimSuffix(output, "\r")

	if output == "" {
		return 0, 0, errors.New("system.GetVideoResolution: ffprobe did not return resolution")
	}

	parts := strings.Split(output, "x")
	if len(parts) < 2 {
		return 0, 0, fmt.Errorf("system.GetVideoResolution: resolution string %s failed to split correctly", output)
	}

	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("system.GetVideoResolution: width part of resolution string %s failed to parse, error: %w", output, err)
	}

	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("system.GetVideoResolution: height part of resolution string %s failed to parse, error: %w", output, err)
	}

	return width, height, nil
}

func RunSystemCommand(cmd *exec.Cmd) (string, error) {
	var cmdOut bytes.Buffer
	var cmdErr bytes.Buffer
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr
	cmdString := cmd.String()

	log.Printf("Running %v\n", cmdString)

	err := cmd.Run()
	outString := cmdOut.String()
	errString := cmdErr.String()
	if err != nil || errString != "" {
		if outString != "" {
			log.Printf("%v stdout: %v", cmdString, outString)
		}

		if errString != "" {
			log.Printf("%v stderr: %v", cmdString, errString)

			if err == nil {
				err = errors.New(errString)
			}
		}

		log.Println(err)

		return errString, err
	} else {
		if outString != "" {
			log.Printf("%v stdout: %v", cmdString, outString)
		}

		return outString, nil
	}
}
