package utils

import "os/exec"

func Cmd(cmd string) []byte {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
	panic("some error found")
	}
	return out
}
