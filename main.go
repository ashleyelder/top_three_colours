package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"image"
	_ "image/jpeg"
	// "log"
	"net/http"
	"os"
	"strings"

	"github.com/EdlinOrg/prominentcolor"
)

var results []string

func readUrls(fileInput string, urls []string) {
	var b bytes.Buffer
	f, err := os.Open(fileInput)

	checkError(&b, err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	for i := 0; i < len(urls); i++ {

		img, err := loadImage(urls[i])

		if err != nil {
			checkError(&b, err)
			continue
		}

		getThreePrevalentColours(img, urls[i])

	}
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
	var result strings.Builder
	result.WriteString(url)

	for _, colour := range colours {
		result.WriteString(",#" + colour.AsString())
	}
	result.WriteString("\n")

	s := result.String()

	results = append(results, s)
}

func createAndWriteCSV() {
	var b bytes.Buffer
	file, err := os.Create("result.csv")

	checkError(&b, err)
	defer file.Close()

	writeCSV(file)
}

func writeCSV(file *os.File) (message string, err error) {
	var b bytes.Buffer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for index, lineItem := range results {
		fmt.Println(index, "=>", lineItem)
		_, err := file.WriteString(lineItem)
		checkError(&b, err)
	}
	return message, nil
}

func checkError(writer *bytes.Buffer, err error) (message string, error error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err.Error())
	}
	return message, nil
}

func main() {
	var filename = "input.txt"

	// creates dynamic array
	var links []string

	readUrls(filename, links)

	createAndWriteCSV()
}
