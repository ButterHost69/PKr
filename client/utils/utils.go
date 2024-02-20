package utils

import (
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else  {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CreateOTP(size int) int {
	mask := strings.Repeat("9", size)
	intmask, _ := strconv.Atoi(mask)
	return rand.Intn(intmask)
}