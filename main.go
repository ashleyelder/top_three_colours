package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/EdlinOrg/prominentcolor"
)

var results []string

func readUrls(fileInput string, urls []string) {

	f, err := os.Open(fileInput)

	if err != nil {
		log.Fatalln("Failure opening file:", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	for i := 0; i < len(urls); i++ {

		img, err := loadImage(urls[i])

		if err != nil {
			fmt.Println(err)
			continue
		}

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

	colours, err := prominentcolor.Kmeans(image)

	checkError("Cannot get colours", err)

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

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writeCSV(file)
}

func writeCSV(file *os.File) {

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for index, lineItem := range results {
		fmt.Println(index, "=>", lineItem)
		_, err := file.WriteString(lineItem)
		if err != nil {
			log.Fatalln("Failed to write data:", err)
		}
	}
}

func checkError(message string, err error) {

	if err != nil {
		log.Fatal(message, err)
	}
}

func main() {

	var filename = "input.txt"

	// creates dynamic array
	var links []string

	readUrls(filename, links)

	createAndWriteCSV()
}
