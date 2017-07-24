package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

var initialCount int = 1

func main() {
	w := newWord()
	var wg sync.WaitGroup

	for i, file_path := range os.Args[1:] {

		wg.Add(1)
		go func(file_path string, i int) {
			fmt.Println("go_routine :", i)
			if err := tally(file_path, w); err != nil {

				fmt.Println(err.Error())
			}
			wg.Done()

		}(file_path, i)
	}

	wg.Wait()
	fmt.Println("Here is the result of tally words")

	for word, count := range w.found {
		fmt.Println(word, count)
	}
}

type word struct {
	sync.Mutex
	found map[string]int
}

func newWord() *word {
	//return &word{make(map[string]int)}
	return &word{found: map[string]int{}}
}

func (w *word) add(word string, count int) {
	w.Lock()
	defer w.Unlock()

	value, ok := w.found[word]

	if !ok {
		w.found[word] = count
	}

	w.found[word] = value + 1

}

func tally(file_path string, w *word) error {
	file, err := os.Open(file_path)

	if nil != err {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		w.add(word, initialCount)
	}

	return scanner.Err()
}
