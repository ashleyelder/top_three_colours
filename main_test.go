package main

import "testing"
import "os"
import "io/ioutil"
import "fmt"

// example test helper that creates temp file - really should be used for very reused logic that doesn't fail often - easier to trace if logic within the test itself insted of abstraction
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
		t.Error("Failed to read csv data")
	}
	fmt.Print(content)
}
