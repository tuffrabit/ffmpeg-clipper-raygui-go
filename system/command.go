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

func RunClipCmd(cmd *exec.Cmd, cancel context.CancelFunc, timestamps chan string, playStates chan bool, errorChan chan error) {
	defer cancel()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		err := fmt.Errorf("system.RunClipCmd: could not get stderr pipe: %w", err)
		fmt.Println(err)
		errorChan <- err
		close(playStates)
		close(timestamps)
		close(errorChan)
		return
	}

	var wg sync.WaitGroup
	var errorBuilder strings.Builder
	isGood := false

	wg.Add(1)
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	go func() {
		for scanner.Scan() {
			currentScan := scanner.Text()

			if !isGood {
				errorBuilder.WriteString(currentScan + " ")
			}

			if strings.Contains(currentScan, "time=") {
				isGood = true
				timestamps <- currentScan[5:]
			}
		}
		wg.Done()
	}()

	if err := cmd.Start(); err != nil {
		err := fmt.Errorf("system.RunClipCmd: failed to start command: %w", err)
		fmt.Println(err)
		errorChan <- err
		close(playStates)
		close(timestamps)
		close(errorChan)
		return
	}

	playStates <- true
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		errorBuilder.WriteString(err.Error())
		err := fmt.Errorf("system.RunClipCmd: failed to run command: %s", errorBuilder.String())
		fmt.Println(err)
		errorChan <- err
		close(playStates)
		close(timestamps)
		close(errorChan)
		return
	}

	playStates <- false
	if !isGood {
		errorChan <- fmt.Errorf("system.RunClipCmd: error with ffmpeg run, error: %s", errorBuilder.String())
	}
	close(playStates)
	close(timestamps)
	close(errorChan)
}

func RunFFplay(videopath string, timestamps chan string, playStates chan bool, errorChan chan error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := exec.CommandContext(ctx, "ffplay", "-y", "680", "-loglevel", "-8", "-stats", videopath)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		err = fmt.Errorf("system.RunFFplay: could not get stderr pipe: %w", err)
		fmt.Println(err)
		errorChan <- err
		close(playStates)
		close(timestamps)
		close(errorChan)
		return
	}

	var wg sync.WaitGroup
	var errorBuilder strings.Builder
	isGood := false

	wg.Add(1)
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)

	go func() {
		prevScan := ""
		capture := false

		for scanner.Scan() {
			currentScan := scanner.Text()

			if !isGood {
				errorBuilder.WriteString(currentScan + " ")
			}

			if strings.Contains(prevScan, "vq=") {
				if currentScan == "0KB" {
					capture = false
				} else {
					capture = true
				}
			}

			if capture && (currentScan == "A-V:" || currentScan == "M-V:" || currentScan == "M-A:") && prevScan != "nan" {
				isGood = true
				timestamps <- prevScan
			}
			prevScan = currentScan
		}
		wg.Done()
	}()

	if err = cmd.Start(); err != nil {
		err = fmt.Errorf("system.RunFFplay: failed to run command: %w", err)
		fmt.Println(err)
		errorChan <- err
		close(playStates)
		close(timestamps)
		close(errorChan)
		return
	}

	playStates <- true
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		errorBuilder.WriteString(err.Error())
		err = fmt.Errorf("system.RunFFplay: failed to run command: %s", errorBuilder.String())
		fmt.Println(err)
		errorChan <- err
		close(playStates)
		close(timestamps)
		close(errorChan)
		return
	}

	playStates <- false
	if !isGood {
		errorChan <- fmt.Errorf("system.RunFFplay: error with ffplay run, error: %s", errorBuilder.String())
	}
	close(playStates)
	close(timestamps)
	close(errorChan)
}

func Play(videopath string) {
	cmd := exec.Command("ffplay", "-y", "680", "-loglevel", "-8", "-stats", videopath)
	RunSystemCommand(cmd)
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
