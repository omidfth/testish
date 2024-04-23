package testish

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func filename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

// Dirname is the __dirname equivalent
func dirname() (string, error) {
	filename, err := filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}

func runDockerCompose(path string) {
	cmd := exec.Command("docker", "compose", "-f", path, "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}

func stopDockerCompose(path string) {
	cmd := exec.Command("docker", "compose", "-f", path, "down")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}
