package model

import (
	"os"
	"path/filepath"
)

func getPromptFilePath(promptPath string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(currentDir, "..", "..", "prompt", promptPath), nil
}

func ReadResumeOptimizePrompt() (string, error) {
	filePath, err := getPromptFilePath("resume_optimize.txt")
	if err != nil {
		return "", err
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
