package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)

		switch {
		case text == "hi":
			fmt.Println("Hello there")
		default:
			fmt.Println("Your previous input was:", text)
		}
	}

}
