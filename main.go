package main

import (
	"log"
	"sync"

	"github.com/meson10/cdnlysis"
)

var wg sync.WaitGroup

func influxReceiver(logChannel <-chan *cdnlysis.LogRecord) {
	defer wg.Done()

	log.Println("Waiting for Records")

	for {
		record, ok := <-logChannel
		if !ok {
			return
		}

		AddToInflux(record)
	}
}

func main() {
	start := ""

	GetConfig()

	RefreshInfluxSession()

	cdnlysis.Setup()

	outputChan := make(chan *cdnlysis.LogRecord)
	go cdnlysis.Start(&start, outputChan)

	wg.Add(1)
	go influxReceiver(outputChan)

	log.Println("Waiting for done!!")
	wg.Wait()
}
