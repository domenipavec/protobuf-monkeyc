package xtesting

import (
	"log"
	"net"
	"os/exec"
	"time"
)

func StartSimulator() error {
	if isSimulatorRunning() {
		log.Println("Simulator already running")
		return nil
	}

	log.Println("Starting simulator")
	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-d",
		"-p", "1234:1234",
		"--entrypoint", "/bin/bash",
		"ghcr.io/matco/connectiq-tester:latest",
		"-c", "Xvfb :1 -screen 0 1280x1024x24 & DISPLAY=:1 simulator",
	)
	err := cmd.Run()
	if err != nil {
		return err
	}

	for !isSimulatorRunning() {
		log.Println("Waiting for simulator...")
		time.Sleep(time.Second)
	}

	return nil
}

func isSimulatorRunning() bool {
	conn, err := net.DialTimeout("tcp", "localhost:1234", time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
