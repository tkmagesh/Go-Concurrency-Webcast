//without wg
package main

import (
	"compress/gzip"
	"io"
	"os"
	//"sync"
)

func main() {
	for _, file := range os.Args {
		compress(file)
	}
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}



//with wg

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var i int
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(filename) {
			compress(filename)
			wg.Done()
		}(file)
	}
	wg.Wait()

	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	fmt.Printf("Compressing %s\n", filename)
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
