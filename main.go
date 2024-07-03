package main

import (
	"log"
	"net/http"
)

func main() {
	storyFileName := "story.json"
	story, err := loadStory(storyFileName)
	if err != nil {
		log.Fatal(err)
	}

	handler, err := newStoryHandler("story.html", story)
	if err != nil {
		log.Fatalf("Failed to create the story handler: %v", err)
	}
	http.Handle("/", handler)

	log.Println("Starting the server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
