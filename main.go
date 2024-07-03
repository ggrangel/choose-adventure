package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
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

type StoryHandler struct {
	story map[string]Arc
}

func (h *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("story.html")
	if err != nil {
		log.Fatalf("Failed to parse the template: %v", err)
	}
	arc := r.URL.Path[1:]
	if arc == "" {
		arc = "intro"
	}
	story, ok := h.story[arc]
	if !ok {
		log.Printf("The arc %s does not exist", arc)
		http.NotFound(w, r)
		return
	}
	chapterHtml := ChapterHtml{
		Title:   story.Title,
		Story:   strings.Join(story.Story, " "),
		Options: story.Options,
	}
	err = tmpl.Execute(w, chapterHtml)
	if err != nil {
		log.Fatalf("Failed to execute the template: %v", err)
	}
}

func main() {
	storyFileName := "story.json"
	storyFile, err := os.Open(storyFileName)
	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}
	defer storyFile.Close()

	var story map[string]Arc
	json.NewDecoder(storyFile).Decode(&story)

	handler := &StoryHandler{story: story}
	http.Handle("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
