package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	// is "gotest.tools/assert/cmp"
	// "gotest.tools/v3/assert"
	"github.com/stretchr/testify/assert"
)

// example of dependency injection
// need to inject (pass in) the dependency of printing
// function doesn't need to know where or how printing happens, so we should accept an interface rather than concrete type
// under the hood Printf uses Fprintf which uses io.Writer
func TestCheckError(t *testing.T) {
	// buffer type from the bytes package implements the Writer interface
	buffer := bytes.Buffer{}

	// example of table
	tests := []struct {
		writer *bytes.Buffer
		err    error
	}{
		{writer: &buffer, err: nil},
	}
	for _, test := range tests {
		result, err := checkError(test.writer, test.err)

		// assert error type
		assert.IsType(t, test.err, err)

		// assert no result when error is nil
		assert.Equal(t, result, "")

		// assert for nil (good for errors)
		assert.Nil(t, err)
	}
}

// example of test helper
// creates temp file
// helper really should be used for very reused logic that doesn't fail often - easier to trace if logic within the test itself insted of abstraction
// should never return an error, accesses T structures so fails the test if there is an error
func testTempFile(t *testing.T) (*os.File, func()) {
	// improves stack trace output better during a panic situation in test helper
	t.Helper()

	tf, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	tf.Close()

	// as cleanup returns closure that deletes the temp file
	return tf, func() { os.Remove(tf.Name()) }
}

func TestWriteCSV(t *testing.T) {
	tf, tfclose := testTempFile(t)

	defer tfclose()

	writeCSV(tf)

	content, err := writeCSV(tf)
	if err != nil {
		t.Error("Failed to read csv data.")
	}
	fmt.Print(content)
}

func TestLoadImage(t *testing.T) {
	// load valid images of cute gophers
	validTests := []struct {
		url string
	}{
		{url: "http://i.imgur.com/FApqk3D.jpg"},
		{url: "https://i.redd.it/xae65ypfqycy.jpg"},
		{url: "https://i.redd.it/1nlgrn49x7ry.jpg"},
	}
	for _, test := range validTests {
		result, err := loadImage(test.url)

		// assert there there is a result - not testing actual decode function so not checking for an image specifically
		assert.NotNil(t, result)

		// assert there is no error
		assert.NoError(t, err)
	}

	// load the invalid image from input.txt
	// why do pictures of cute gophers not decode? unfortunately format not recognized for some reason
	invalidTests := []struct {
		url string
	}{
		{url: "https://i.redd.it/nrafqoujmety.jpg"},
		{url: "https://raw.githubusercontent.com/ashleymcnamara/gophers/master/NERDY.png"},
		{url: "https://raw.githubusercontent.com/ashleymcnamara/gophers/master/BELGIUM.png"},
		{url: "https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GOPHER_DENVER.png"},
	}
	for _, test := range invalidTests {
		result, err := loadImage(test.url)

		// assert result is nil when invalid image passed
		assert.Equal(t, result, nil)

		// assert there is an error with "image: unknown format"
		assert.Error(t, err, "image: unknown format")
	}
}

// func TestAssembleLineItem(t *testing.T) {
// test creates assembles a lineItem and appends to results array
// }
