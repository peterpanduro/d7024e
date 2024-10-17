package main

import (
	"d7024e/kademlia"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestHandleCommandPutAndGet(t *testing.T) {
	// Setup a Kademlia instance
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")
	kad := kademlia.NewKademlia(contact)

	// Create a temporary file with test content
	content := []byte("Test content")
	tmpfile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up the file afterward

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Capture stdout
	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Test the "put" command
	putCommand := fmt.Sprintf("put %s\n", tmpfile.Name())
	handleCommand(putCommand, kad)

	// Flush and restore stdout
	wOut.Close()
	os.Stdout = oldStdout
	outBytes, _ := ioutil.ReadAll(rOut)
	output := string(outBytes)

	// Extract the hash from the output
	var hash string
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "File uploaded successfully. Hash:") {
			hash = strings.TrimSpace(strings.TrimPrefix(line, "File uploaded successfully. Hash:"))
			break
		}
	}

	if hash == "" {
		t.Fatal("Failed to retrieve hash from output")
	}

	// Capture stdout for the "get" command
	rOut, wOut, _ = os.Pipe()
	os.Stdout = wOut

	// Test the "get" command
	getCommand := fmt.Sprintf("get %s\n", hash)
	handleCommand(getCommand, kad)

	// Flush and restore stdout
	wOut.Close()
	os.Stdout = oldStdout
	outBytes, _ = ioutil.ReadAll(rOut)
	output = string(outBytes)

	if !strings.Contains(output, "Data retrieved: Test content") {
		t.Errorf("Expected 'Test content' in output, got '%s'", output)
	}
}

func TestHandleCommandInvalid(t *testing.T) {
	// Setup a Kademlia instance
	contact := kademlia.NewContact(kademlia.NewRandomKademliaID(), "127.0.0.1:8080")
	kad := kademlia.NewKademlia(contact)

	// Capture stdout
	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Test an invalid command
	handleCommand("invalid\n", kad)

	// Flush and restore stdout
	wOut.Close()
	os.Stdout = oldStdout
	outBytes, _ := ioutil.ReadAll(rOut)
	output := string(outBytes)

	expected := "Invalid command\n"
	if output != expected {
		t.Errorf("Expected output '%s', got '%s'", expected, output)
	}
}
