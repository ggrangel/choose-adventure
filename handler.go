package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type StoryHandler struct {
	tmpl  *template.Template
	story map[string]Arc
}

func newStoryHandler(templatePath string, story map[string]Arc) (*StoryHandler, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse the template: %v", err)
	}

	return &StoryHandler{
		tmpl:  tmpl,
		story: story,
	}, nil
}

func (h *StoryHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	arc := r.URL.Path[1:]
	if arc == "" {
		arc = "intro"
	}

	story, ok := h.story[arc]
	if !ok {
		http.NotFound(w, r)
		return
	}

	chapterHtml := ChapterHtml{
		Title:   story.Title,
		Story:   strings.Join(story.Story, " "),
		Options: story.Options,
	}

	if err := h.tmpl.Execute(w, chapterHtml); err != nil {
		log.Printf("Failed to execute template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
