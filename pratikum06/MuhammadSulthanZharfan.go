package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	// membuat direktori baru
	err = os.Chmod("MuhammadSulthanZharfan", 0016)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Izin 'MuhammadSulthanZharfan' telah diubah menjadi 0016")
}
