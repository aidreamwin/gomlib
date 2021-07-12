package disk

import (
	"log"
	"testing"
)

func TestDisk(t *testing.T) {

	avaiable, free, total, err := DiskUsage()
	if err != nil {
		t.Error(err)
		return
	}

	log.Printf("Available  %dmb", avaiable/1024/1024.0)
	log.Printf("Free       %dmb", free/1024/1024.0)
	log.Printf("Total      %dmb", total/1024/1024.0)
}
