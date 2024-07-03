package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type ChapterHtml struct {
	Title   string
	Story   string
	Options []Option
}

func loadStory(filePath string) (map[string]Arc, error) {
	storyFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read story file: %v", err)
	}
	defer storyFile.Close()

	var story map[string]Arc
	// Using decoder to decode an io.Reader
	// To decode []byte (output from os.ReadFile), we use json.Unmarshal([]byte, &story) instead
	// The former has the advantage of not loading the entire file into memory
	if err := json.NewDecoder(storyFile).Decode(&story); err != nil {
		return nil, fmt.Errorf("Failed to decode story data: %v", err)
	}

	return story, nil
}
