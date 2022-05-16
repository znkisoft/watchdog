package ssm

import (
	"log"
	"os/exec"
)

type IUptime interface {
	Uptime() (int, error)
}

func Uptime() (int, error) {
	cmd := exec.Command("uptime")

	if err := cmd.Run(); err != nil {
		return 0, err
	}
	defer cmd.Process.Kill()

	data, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	log.Printf("%s", data)
	return 1, nil
}
