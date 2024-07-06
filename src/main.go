package main

import "fmt"

func main() {
	conn := openConn()
	defer conn.Close()

	onActiveWindowChange(conn, func(data string) {
		windowName := getWindowName()

		fmt.Println(windowName)
	})
}
