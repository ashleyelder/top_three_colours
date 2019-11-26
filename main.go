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

// makes global because difficulty accessing in readUrls and then processUrls (which isn't called from readUrls)
var urls = make(chan string)

func readUrls(fileInput string, results chan string) {
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
}

func processUrls(results chan string) {
	fmt.Println("processUrls")
	var b bytes.Buffer

	for url := range urls {
		fmt.Println(url)
		img, err := loadImage(url)

		if err != nil {
			checkError(&b, err)
			fmt.Println(img)
			continue
		}

		getThreePrevalentColours(img, url, results)
	}
}

func loadImage(url string) (image.Image, error) {
	fmt.Println("loadImage")
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

func getThreePrevalentColours(image image.Image, url string, results chan string) {
	fmt.Println("getThreePrevalentColours")
	var b bytes.Buffer
	colours, err := prominentcolor.Kmeans(image)

	// pass reference to buffer instead of buffer itself
	checkError(&b, err)

	assembleLineItem(colours, url, results)
}

func assembleLineItem(colours []prominentcolor.ColorItem, url string, results chan string) {
	fmt.Println("assembleLineItem")

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

func checkError(writer *bytes.Buffer, err error) (message string, error error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err.Error())
	}
	return message, nil
}

func main() {
	res := make(chan string)

	filename := "input.txt"

	go readUrls(filename, res)
	go processUrls(res)

	s := <-res
	fmt.Println("s", s)

	fmt.Scanln()

	// createAndWriteCSV(results)
}
