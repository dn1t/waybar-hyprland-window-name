package main

import "fmt"

func main() {
	conn := openConn()
	defer conn.Close()

	onActiveWindowChange(conn, func(class string) {
		fmt.Println(getWindowName(class))
	})
}
