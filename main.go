package main

import (
	"github.com/aidreamwin/gomlib"
	"log"
)

func main() {
	log.Println("gomlib")
	lpFreeBytesAvailable, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes, err := disk.DiskUsage()
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Available  %dmb", lpFreeBytesAvailable/1024/1024.0)
	log.Printf("Total      %dmb", lpTotalNumberOfBytes/1024/1024.0)
	log.Printf("Free       %dmb", lpTotalNumberOfFreeBytes/1024/1024.0)
}
