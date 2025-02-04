package xtesting

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RebuildProto() error {
	cmd := exec.Command("go", "install")
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dockerArgs := []string{
		"run",
		"--rm",
		"-v", filepath.Dir(pwd) + ":/app",
		"-v", home + "/go/bin/protoc-gen-monkeyc:/bin/protoc-gen-monkeyc",
		"-w", "/app/Example",
		"rvolosatovs/protoc",
		"-I.",
		"--monkeyc_out=.",
		"--go_out=../protoc-gen-monkeyc/example",
		"--go_opt=Mexample.proto=github.com/domenipavec/protobuf-monkeyc/protoc-gen-monkeyc/example",
		"--go_opt=paths=source_relative",
		"example.proto",
	}
	log.Print(strings.Join(dockerArgs, " "))
	cmd = exec.Command(
		"docker", dockerArgs...,
	)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
