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
	i := 0
	for _, file_path := range os.Args[1:] {

		i += 1
		wg.Add(1)
		go func(file_url string) {
			compress(file_url)

			wg.Done()

		}(file_path)
	}

	wg.Wait()
	fmt.Printf("%f - Gzip completed \n", i)
}

//accepts a path to the file to compress
//compresses the file to gzip into the same folder
func compress(file_path string) error {

	file, err := os.Open(file_path)

	if err != nil {
		return err
	}

	defer file.Close()

	newFile, err := os.Create(file_path + ".gz")

	if err != nil {
		return err
	}

	defer newFile.Close()

	gzipOut := gzip.NewWriter(newFile)

	_, err = io.Copy(newFile, file)

	gzipOut.Close()

	return err

}
