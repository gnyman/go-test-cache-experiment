package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestWorkingDirectory(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/hello.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
