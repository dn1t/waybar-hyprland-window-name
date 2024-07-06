package main

import (
	"os"
	"strings"

	"code.rocketnine.space/tslocum/desktop"
)

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

func getWindowName(class string) string {
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
