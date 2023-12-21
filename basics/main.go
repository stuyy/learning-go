package main

import (
	"bufio"
	"fmt"
	"learning-go/basics/utils"
	"os"
)
func main() {
	fileName := "names.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		println("Error Occured")
		panic(err)
	}

	in := bufio.NewScanner(os.Stdin)
	scanned := in.Scan()
	
	for scanned {
		text := in.Text()
		if text == "exit" {
			in.Err()
			break
		} else {
			fmt.Printf("You wrote: %s\n", text)
			formattedText := fmt.Sprintf("%s\n", text)
			file.WriteString(formattedText)
			scanned = in.Scan()
		}
	}
	file.Sync()
	file.Close()
	utils.ReadFileLineByLine(fileName)
}