package main

import (
	"fmt"
	"log"
	"stock-pick/listeners"
	"stock-pick/readers"
)

func main() {

	activityChannel := make(chan string)

	StartListeners(activityChannel)

	select {
	case message := <-activityChannel:
		fmt.Println(message)
	}
}

//StartListeners gets all the symbols and creates listeners for them
func StartListeners(activityChannel chan string) {
	reader, error := readers.NewReader(readers.InMemory)

	if error != nil {
		log.Fatal("Failure to create symbol reader")
	}

	symbols, error := reader.GetSymbols()

	if error != nil {
		log.Fatal("Failure to retrieve symbols")
	}

	for _, symbol := range symbols {
		listener, err := listeners.NewListener()
		if err != nil {
			panic(err)
		}
		go listener.ListenToSymbol(symbol, activityChannel)
	}
}
