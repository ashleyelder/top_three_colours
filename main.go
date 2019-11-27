package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/EdlinOrg/prominentcolor"
	"image"
	_ "image/jpeg"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var urls = make(chan string, 10)
var results = make(chan string, 10)

func readUrls(fileInput string) {
	var b bytes.Buffer

	f, err := os.Open(fileInput)

	checkError(&b, err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		url := scanner.Text()
		urls <- url
	}
	close(urls)
}

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

	// pass reference to buffer instead of buffer itself
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

func createAndWriteCSV(done chan bool) {
	var b bytes.Buffer
	file, err := os.Create("result.csv")

	checkError(&b, err)
	defer file.Close()

	writeCSV(file, done)
}

func writeCSV(file *os.File, done chan bool) (message string, err error) {
	var b bytes.Buffer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for lineItem := range results {
		_, err := file.WriteString(lineItem)
		checkError(&b, err)
	}

	done <- true

	return message, nil
}

func worker(wg *sync.WaitGroup) {
	var b bytes.Buffer
	for url := range urls {
		img, err := loadImage(url)
		if err != nil {
			checkError(&b, err)
			continue
		}

		getThreePrevalentColours(img, url)
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func checkError(writer *bytes.Buffer, err error) (message string, error error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err.Error())
	}
	return message, nil
}

func main() {
	startTime := time.Now()
	filename := "input.txt"

	go readUrls(filename)
	done := make(chan bool)
	go createAndWriteCSV(done)

	noOfWorkers := 10
	createWorkerPool(noOfWorkers)

	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
