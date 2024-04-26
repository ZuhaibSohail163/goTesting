package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestHelloCmd(t *testing.T) {
	// Create a new buffer to capture the output
	var outputBuffer bytes.Buffer

	// Mock Cobra's root command
	rootCmd := &cobra.Command{Use: "test"}
	rootCmd.SetOut(&outputBuffer)

	// Set output for hello command
	var stdout bytes.Buffer
	helloCmd.SetOut(&stdout)

	// Execute the hello command
	// helloCmd.Run(helloCmd, []string{})
	// Define test input arguments
	var testArgs []string

	// Execute the hello command with test arguments
	helloCmd.Run(helloCmd, testArgs)

	// Verify the output
	expectedOutput := "hello called" // Update the expected output
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, testArgs)
	}
}
