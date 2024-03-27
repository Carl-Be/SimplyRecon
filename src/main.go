package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go StartNmapScanning(&wg)

	wg.Wait()
}
