package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

var testMessage = "epsilon"

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage(testMessage)
	wg.Wait()
	if msg != testMessage {
		t.Errorf("Expected to find %s, but it is not there\n", testMessage)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	msg = testMessage
	printMessage()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, testMessage) {
		t.Errorf("Expected to find %s, but it is not there\n", testMessage)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	testMessage := "Hello, universe!"
	if !strings.Contains(output, testMessage) {
		t.Errorf("Expected to find %s, but it is not there\n", testMessage)
	}
	testMessage = "Hello, cosmos!"
	if !strings.Contains(output, testMessage) {
		t.Errorf("Expected to find %s, but it is not there\n", testMessage)
	}
	testMessage = "Hello, world!"
	if !strings.Contains(output, testMessage) {
		t.Errorf("Expected to find %s, but it is not there\n", testMessage)
	}

}
