package main

import (
	"github.com/aidreamwin/gomlib/disk"
	"github.com/aidreamwin/gomlib/log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Trace("gomlib", "ddddd有意思")
	log.Debug("gomlib")
	log.Info("gomlib")
	log.Warn("gomlib")
	log.Error("gomlib")
	lpFreeBytesAvailable, lpTotalNumberOfBytes, lpTotalNumberOfFreeBytes, err := disk.DiskUsage("C:")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Available  %dmb", lpFreeBytesAvailable/1024/1024.0)
	log.Printf("Total      %dmb", lpTotalNumberOfBytes/1024/1024.0)
	log.Printf("Free       %dmb", lpTotalNumberOfFreeBytes/1024/1024.0)
}
