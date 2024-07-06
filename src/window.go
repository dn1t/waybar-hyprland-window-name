package main

import (
	"os"
	"os/exec"
	"strings"

	"code.rocketnine.space/tslocum/desktop"
)

func getActiveWindowClass() string {
	hyprCmd := exec.Command("sh", "-c", "hyprctl activewindow | sed -n '/class:/s/class://p' | xargs")
	stdout, err := hyprCmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(stdout))
}

func getDirs(class string) []string {
	home := os.Getenv("HOME")

	prefixes := [...]string{home + "/.local/share/", "/usr/local/share/", "/usr/share/"}
	appDirs := [...]string{"applications/", "applications/kde/", "applications/org.kde.", ""}
	suffixes := [...]string{".desktop", ""}

	dirs := []string{}
	for _, prefix := range prefixes {
		for _, appDir := range appDirs {
			for _, suffix := range suffixes {
				dirs = append(dirs, prefix+appDir+class+suffix)
				dirs = append(dirs, prefix+appDir+strings.ToLower(class)+suffix)
			}
		}
	}

	return dirs
}

func getWindowName() string {
	class := getActiveWindowClass()

	if class != "" {
		dirs := getDirs(class)
		entries, err := desktop.Scan(dirs)

		if err == nil {
			for _, entry := range entries {
				if len(entry) > 0 {
					return entry[0].Name
				}
			}
		}
	}

	return "Hyprland"
}
