package main

import (
	"bufio"
	"bytes"
	// // "encoding/csv"
	"fmt"
	"github.com/EdlinOrg/prominentcolor"
	"image"
	_ "image/jpeg"
	"net/http"
	"os"
	"strings"
	// "time"
)

var urls = make(chan string, 10)
var results = make(chan image.Image, 10)

func readUrls(fileInput string) {
	fmt.Println("readUrls")
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

// func processUrls(results chan string) {
// 	fmt.Println("processUrls")
// 	var b bytes.Buffer

// 	for url := range urls {
// 		fmt.Println(url)
// 		img, err := loadImage(url)

// 		if err != nil {
// 			checkError(&b, err)
// 			fmt.Println(img)
// 			continue
// 		}

// 		getThreePrevalentColours(img, url, results)
// 	}
// }

// func loadImage(url string) (image.Image, error) {
// 	fmt.Println("loadImage")
// 	var b bytes.Buffer
// 	response, err := http.Get(url)

// 	checkError(&b, err)
// 	defer response.Body.Close()

// 	img, _, err := image.Decode(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return img, nil
// }

// func getThreePrevalentColours(image image.Image, url string, results chan string) {
// 	fmt.Println("getThreePrevalentColours")
// 	var b bytes.Buffer
// 	colours, err := prominentcolor.Kmeans(image)

// 	// pass reference to buffer instead of buffer itself
// 	checkError(&b, err)

// 	assembleLineItem(colours, url, results)
// }

// func assembleLineItem(colours []prominentcolor.ColorItem, url string, results chan string) {
// 	fmt.Println("assembleLineItem")

// 	// build a line item
// 	var str strings.Builder
// 	str.WriteString(url)
// 	for _, colour := range colours {
// 		str.WriteString(",#" + colour.AsString())
// 	}
// 	str.WriteString("\n")
// 	lineItem := str.String()

// 	results <- lineItem
// }

// func createAndWriteCSV(results chan []string) {
// 	var b bytes.Buffer
// 	file, err := os.Create("result.csv")

// 	checkError(&b, err)
// 	defer file.Close()

// 	writeCSV(file, results)
// }

// func writeCSV(file *os.File, results chan []string) (message string, err error) {
// 	var b bytes.Buffer
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	for index, lineItem := range results {
// 		fmt.Println(index, "=>", lineItem)
// 		_, err := file.WriteString(lineItem)
// 		checkError(&b, err)
// 	}
// 	return message, nil
// }

func worker(wg *sync.WaitGroup) {
	for url := range urls {
		fmt.Println(url)
		// output := Result{url, digits(url.randomno)}
		// img, err := loadImage(url)
		// fmt.Println(err)
		// results <- img
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

func result(done chan bool) {
	for result := range results {
		fmt.Printf("result function: ", result)
	}
	done <- true
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
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
	// go processUrls()
	// createAndWriteCSV(results)
}
