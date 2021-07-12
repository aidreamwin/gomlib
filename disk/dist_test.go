package disk

import (
	"log"
	"testing"
	"time"
)

func TestDisk(t *testing.T) {
	sTime := time.Now().UnixNano()
	avaiable, free, total, err := DiskUsage()
	if err != nil {
		t.Error(err)
		return
	}
	eTime := time.Now().UnixNano()

	log.Printf("Available  %dmb", avaiable/1024/1024.0)
	log.Printf("Free       %dmb", free/1024/1024.0)
	log.Printf("Total      %dmb", total/1024/1024.0)
	log.Println(sTime, eTime-sTime, (eTime-sTime)/10e6)
}
