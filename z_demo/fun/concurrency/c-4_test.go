package concurrency

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

var sema = make(chan struct{}, 20)
var wg sync.WaitGroup
var done = make(chan struct{})

func TestC4(t *testing.T) {
	roots := []string{
		`C:\`,
	}
	// Traverse the file tree.
	fileSizes := make(chan int64)
	for _, dir := range roots {
		wg.Add(1)
		go walkDirSync(dir, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Print the results.
	var tick <-chan time.Time
	tick = time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// Do nothing.
				log.Println("clear")
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func walkDirSync(s string, wg *sync.WaitGroup, fileSizes chan int64) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(s) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(s, entry.Name())
			go walkDirSync(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() {
		<-sema
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
