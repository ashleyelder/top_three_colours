package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"image"
	_ "image/jpeg"
	// "io"
	"log"
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
			// fmt.Println(err)
			checkError(&b, err)
			continue
		}
		// checkError(&b, err)

		getThreePrevalentColor(img, urls[i])

	}
}

func loadImage(url string) (image.Image, error) {
	response, err := http.Get(url)

	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getThreePrevalentColor(image image.Image, url string) {
	var b bytes.Buffer
	colours, err := prominentcolor.Kmeans(image)

	// pass reference to buffer instead of buffer itself
	checkError(&b, err)

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
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for index, lineItem := range results {
		fmt.Println(index, "=>", lineItem)
		_, err := file.WriteString(lineItem)
		if err != nil {
			log.Fatalln("Failed to write data:", err)
		}
	}
	return message, nil
}

// err error
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
