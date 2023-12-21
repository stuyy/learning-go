package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileLineByLine(fileName string) {
	reader, err := os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	in := bufio.NewScanner(reader)
	scanned := in.Scan()
	for scanned {
		text := in.Text()
		fmt.Printf("Name: %s\n", text)
		scanned = in.Scan()
	}
}

func ReadFileAndPrint(fileName string) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	str := string(contents)
	println(str)
}