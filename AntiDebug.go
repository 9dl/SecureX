package github.com/StreamlineX/SecureX

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)


func AntiDebugRun() {
	for {
		processes, err := exec.Command("tasklist").Output()
		if err != nil {
			log.Fatal(err)
		}

		processList := string(processes)

		for _, blacklistedProcess := range Blacklisted {
			if strings.Contains(processList, blacklistedProcess) {
				err := killProcess(blacklistedProcess)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		if isDebuggerAttached() {
			log.Println("Debugger detected. Terminating...")
			os.Exit(0)
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

	isDebugging := false

	ret, _, _ := syscall.Syscall(uintptr(isDebuggerPresent), 0, 0, 0, 0)
	if ret != 0 {
		isDebugging = true
	}

	return isDebugging
}

var Blacklisted = []string{
	"processh",
	"debug",
	"debugger",
	"hacker",
	"inject",
	"dump",
	"dumper",
	"deobfs",
	"deobfuscator",
	"dnspy",
	"de4dot",
	"dbg",
	"string",
	"decrypt",
	"decryptor",
	"detect it easy",
	"die",
	"unpack",
	"unpacker",
	"http",
}
