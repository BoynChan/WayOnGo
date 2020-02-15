package du

import (
	"fmt"
	"testing"
	"time"
)

/**
author:Boyn
date:2020/2/15
*/

func TestWalkDir(t *testing.T) {
	initDir := []string{"F:\\Code"}
	fileSizes := make(chan int64)
	used := time.Now()
	go func() {
		for _, root := range initDir {
			WalkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes {
		nbytes += size
		nfiles++
	}
	printDiskUsage(nfiles, nbytes)
	fmt.Printf("Time Used:%.2f ms", float64(used.Nanosecond())/1000000)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files : %.1f MB Uses\n", nfiles, float64(nbytes)/1e6)
}
