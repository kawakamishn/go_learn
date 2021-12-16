//./du3 -v [pass] [pass]

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

// どのrootに所属するかをnameに記録する
type directoryInformation struct {
	name string
	size int64
}

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan *directoryInformation)
	var waitgroup sync.WaitGroup
	directories := make(map[string]*directoryInformation)

	for _, root := range roots {
		directories[root] = &directoryInformation{name: root}
		waitgroup.Add(1)
		go walkDir(root, root, &waitgroup, fileSizes)
	}
	go func() {
		waitgroup.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case intermediateDirectoryInformation, ok := <-fileSizes:
			if !ok {
				break loop
			}
			directories[intermediateDirectoryInformation.name].size += intermediateDirectoryInformation.size
		case <-tick:
			printDiskUsage(roots, directories)
		}
	}
	printDiskUsage(roots, directories)
	fmt.Printf("\n")
}

func printDiskUsage(roots []string, directories map[string]*directoryInformation) {
	fmt.Printf("\r")
	for _, root := range roots {
		fmt.Print(root, ": ", float64(directories[root].size)/1e9, " GB")
	}
}

func walkDir(dir string, root string, waitgroup *sync.WaitGroup, fileSizes chan<- *directoryInformation) {
	defer waitgroup.Done()

	for _, entry := range direntries(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			waitgroup.Add(1)
			go walkDir(subdir, root, waitgroup, fileSizes)
		} else {
			fileSizes <- &directoryInformation{name: root, size: entry.Size()} //ここでrootのnameの情報も付加
		}
	}
}

var sema = make(chan struct{}, 32)

func direntries(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4: %v\n", err)
		return nil
	}
	return entries
}
