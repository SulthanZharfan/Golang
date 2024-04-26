package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	fileInfo, err := os.Stat("MuhammadSulthanZharfan")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("MuhammadSulthanZharfan adalah sebuah direktori.")
	} else {
		fmt.Println("MuhammadSulthanZharfan adalah sebuah fie.")
	}
}
