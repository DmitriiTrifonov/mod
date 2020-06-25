package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	message       = "mod: try mod --help or -h for more information"
	shortHelpFlag = "-h"
	helpFlag      = "--help"
	notInt        = "number is not an int type"
	helpMessage   = `Usage: mod [number1] [number2]
    -h or --help: shows help message`
)

func main() {
	switch len(os.Args) {
	case 3:
		f, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalln("The first", notInt)
		}

		s, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalln("The second", notInt)
		}
		fmt.Println(f % s)
	case 2:
		if os.Args[1] == shortHelpFlag || os.Args[1] == helpFlag {
			fmt.Println(helpMessage)
		}
	default:
		fmt.Println(message)
	}
}
