// package main

// import "testing"

// import "bytes"

// func TestWriteCSV(t *testing.T) {

// 	var b bytes.Buffer
// 	b.WriteString("fake, csv, data")

// buffer cannot be used as am argument to os.File because it doesn't have a file descriptor, a number managed by the os
// content, err := writeCSV(&b)
// if err != nil {
// 	t.Error("Failed to read csv data")
// }
// fmt.Print(content)
// }

