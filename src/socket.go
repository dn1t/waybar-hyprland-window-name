package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func getHyprSocket(name string) string {
	instance := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if instance == "" {
		panic("Hyprland is not running")
	}

	dir := os.Getenv("XDG_RUNTIME_DIR")
	if dir != "" {
		dir += "/hypr"
	} else {
		panic("Couldn't find current runtime dir")
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		dir = "/tmp/hypr"
	}

	return fmt.Sprintf("%s/%s/.%s.sock", dir, instance, name)
}

func openConn() net.Conn {
	socket, err := net.Dial("unix", getHyprSocket("socket2"))
	if err != nil {
		panic(err)
	}

	return socket
}

func onActiveWindowChange(conn net.Conn, listener func(data string)) {
	buf := make([]byte, 4096)

	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			return
		}

		listener(strings.TrimSpace(strings.Split(strings.Split(strings.Split(string(buf[0:n]), "\n")[0], ">>")[1], ",")[0]))
	}
}
