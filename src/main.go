package main

import "fmt"

func main() {
	conn := openConn()
	defer conn.Close()

	onActiveWindowChange(conn, func(class string) {
		name := getWindowName(class)
		fmt.Printf("{\"text\": \"%s\", \"class\": \"waybar-hyprland-window-name\"}\n", name)
	})
}
