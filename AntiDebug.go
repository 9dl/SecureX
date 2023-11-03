package github.com/9dl/SecureX

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var blacklisted = []string{
	"processh", "debug", "debugger", "hacker", "inject", "dump",
	"dumper", "deobfs", "deobfuscator", "dnspy", "de4dot", "dbg",
	"string", "decrypt", "decryptor", "detect it easy", "die",
	"unpack", "unpacker", "http",
}

func AntiDebugRun() {
	for {
		if isDebuggerAttached() {
			log.Println("Debugger detected. Terminating...")
			os.Exit(0)
		}

		processList, err := getProcessList()
		if err != nil {
			log.Fatal(err)
		}

		killBlacklistedProcesses(processList)
	}
}

func getProcessList() (string, error) {
	processes, err := exec.Command("tasklist").Output()
	if err != nil {
		return "", err
	}
	return string(processes), nil
}

func killBlacklistedProcesses(processList string) {
	for _, blacklistedProcess := range blacklisted {
		if strings.Contains(processList, blacklistedProcess) {
			err := killProcess(blacklistedProcess)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func killProcess(processName string) error {
	cmd := exec.Command("taskkill", "/F", "/IM", processName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func isDebuggerAttached() bool {
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return false
	}
	defer syscall.FreeLibrary(kernel32)

	isDebuggerPresent, err := syscall.GetProcAddress(kernel32, "IsDebuggerPresent")
	if err != nil {
		return false
	}

	ret, _, _ := syscall.Syscall(uintptr(isDebuggerPresent), 0, 0, 0, 0)
	return ret != 0
}
