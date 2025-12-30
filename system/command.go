package system

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func RunFFplay(videopath string, timestamps chan string, playStates chan bool) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := exec.CommandContext(ctx, "ffplay", "-x", "1280", "-loglevel", "-8", "-stats", videopath)
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
		for scanner.Scan() {
			currentScan := scanner.Text()
			if currentScan == "A-V:" && prevScan != "nan" {
				if err == nil {
					timestamps <- prevScan
				}
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
