package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	if file, err := Opencsv("excel.csv"); err != nil {
		fmt.Println("error occured while try to read csv", err)
	} else {
		fmt.Println(file.Name())
	}

}

func Opencsv(file_path string) (file *os.File, err error) {

	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()

	file, err = os.Open(file_path)

	if nil != err {
		return file, err
	}

	removeEmptyLines(file)

	return file, nil
}

func removeEmptyLines(file *os.File) {
	panic(errors.New("something weired happened while attempting to remove lines"))
}
