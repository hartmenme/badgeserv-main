package server

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadBadgeConfig(t *testing.T) {
	// Setup: Create a temporary directory for testing
	tempDir, err := ioutil.TempDir("", "testbadgeconfig")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir) // clean up

	// Test: Call loadBadgeConfig with the temp directory
	_, err = loadBadgeConfig(tempDir)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Additional tests for other scenarios...
}
