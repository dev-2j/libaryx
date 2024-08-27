package main

import (
	"fmt"
	"log"

	"example.com/m/routex"
)

func main() {

	// routex
	routex.Route()

	// เริ่มเซิร์ฟเวอร์
	fmt.Println("Server is running on port 4000")
	log.Fatal(`Server is running on port 4000`)
}
