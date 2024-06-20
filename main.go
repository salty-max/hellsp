package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/salty-max/hellsp/rpc"
)

func main() {
	logger := getLogger("log.txt")
	logger.Println("hellsp started!")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()

		handleMessage(msg)
	}
}

func handleMessage(_ any) {}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("could not open or create file: %s", filename))
	}

	return log.New(logfile, "[hellsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
