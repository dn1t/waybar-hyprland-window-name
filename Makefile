default: build

build:
	go build -ldflags "-s -w" -o ./waybar-hyprland-window-name ./src
