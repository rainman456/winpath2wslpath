package main

import (
	"fmt"
	"os"
	"strings"
)

// convertWindowsPathToWSL converts a Windows path to the WSL (Ubuntu) format.
func convertWindowsPathToWSL(windowsPath string) string {
	// Replace backslashes with forward slashes
	windowsPath = strings.ReplaceAll(windowsPath, "\\", "/")

	// Check if the path starts with a drive letter (e.g., "C:")
	if strings.Contains(windowsPath, ":") {
		// Split the drive letter and the rest of the path
		driveAndPath := strings.SplitN(windowsPath, ":", 2)
		drive := strings.ToLower(driveAndPath[0]) // Convert drive letter to lowercase
		restOfPath := driveAndPath[1]

		// Construct the WSL path
		wslPath := fmt.Sprintf("/mnt/%s%s", drive, restOfPath)
		return wslPath
	}

	// If no drive letter is found, assume it's already a WSL path or relative path
	return windowsPath
}

func main() {
	// Check if a path argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: wslpath <windows-path>")
		fmt.Println("Example: wslpath 'C:\\Users\\USER\\Desktop\\example'")
		os.Exit(1)
	}

	// Get the Windows path from the command-line argument
	windowsPath := os.Args[1]

	// Convert the Windows path to WSL format
	wslPath := convertWindowsPathToWSL(windowsPath)

	// Output the result
	fmt.Printf("Windows Path: %s\n", windowsPath)
	fmt.Printf("WSL Path:     %s\n", wslPath)
}