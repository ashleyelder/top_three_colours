package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/EdlinOrg/prominentcolor"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// two bidirectionsl channels so they can send and receive data
// urls will be sent to workers and results collected
// sends to buffered channel are blocked when buffer is full - it is possible to write 10 urls into urls channel without being blocked
// receives from buffered channel are blocked only when buffer is empty
var urls = make(chan string, 10)
var results = make(chan string, 10)

func readUrls(fileInput string) {
	var b bytes.Buffer

	f, err := os.Open(fileInput)

	checkError(&b, err)

	// defer function call to the end of the currently executing function
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		url := scanner.Text()
		// data sent to urls channel, writes are blocked until the worker goroutine reads from this channel
		urls <- url
		fmt.Println("successfully wrote", url, "to channel")
	}
	// close urls channel to notify receivers that no more data will be sent on the channel
	close(urls)
}

// returns actual image from the url
func loadImage(url string) (image.Image, error) {
	var b bytes.Buffer
	response, err := http.Get(url)

	checkError(&b, err)
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getThreePrevalentColours(image image.Image, url string) {
	var b bytes.Buffer
	colours, err := prominentcolor.Kmeans(image)

	// pass buffer by pointer
	checkError(&b, err)
	assembleLineItem(colours, url)
}

func assembleLineItem(colours []prominentcolor.ColorItem, url string) {
	// build a line item
	var str strings.Builder
	str.WriteString(url)
	for _, colour := range colours {
		str.WriteString(",#" + colour.AsString())
	}
	str.WriteString("\n")
	lineItem := str.String()

	results <- lineItem
}

func createAndWriteCSV(wg *sync.WaitGroup) {
	var b bytes.Buffer
	file, err := os.Create("result.csv")

	checkError(&b, err)
	defer file.Close()

	writeCSV(file, wg)
}

func writeCSV(file *os.File, wg *sync.WaitGroup) (message string, err error) {
	var b bytes.Buffer
	writer := csv.NewWriter(file)
	// writes any buffered data to the underlying io.Writer
	defer writer.Flush()

	for lineItem := range results {
		_, err := file.WriteString(lineItem)
		checkError(&b, err)
	}

	defer wg.Done()

	return message, nil
}

// frequent creation of workers receive tasks on the url channel
func worker(wg *sync.WaitGroup) {
	var b bytes.Buffer
	// a defer function to recover is a good idea here because any panics would otherwise crash the entire program
	defer wg.Done()

	// pull urls from queue until it's done/closed so no need to check if channel is closed with ok variable
	for url := range urls {
		fmt.Println("concurrently read value", url, "from channel")
		img, err := loadImage(url)
		if err != nil {
			checkError(&b, err)
			continue
		}
		getThreePrevalentColours(img, url)
	}
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	// wait for the results to complete
	wg.Wait()

	// signal that we are done collecting results
	// close the channel to notify receivers that no more data will be sent on the channel
	close(results)
}

// helper function to print any error
func checkError(writer *bytes.Buffer, err error) (message string, error error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err.Error())
	}
	return message, nil
}

func main() {
	startTime := time.Now()
	filename := "input.txt"
	var wg sync.WaitGroup

	go readUrls(filename)
	wg.Add(1)
	go createAndWriteCSV(&wg)

	// control the number of concurrently running tasks
	// number is tuned to computing resources available, adjustable to optimize performance
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)

	wg.Wait()

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
