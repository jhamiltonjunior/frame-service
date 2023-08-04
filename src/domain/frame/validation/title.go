package validation

import (
	"errors"
	"strings"
)

type Title struct {
	Title string
}

func (t *Title) Create (title string) (Title, error) {
	title = strings.Trim(title, " ")
	
	// nao consegui um regex pra isso
	title = strings.ReplaceAll(title, "  ", " ")
	title = strings.ReplaceAll(title, "  ", " ")
	title = strings.ReplaceAll(title, "  ", " ")

	if t.noValid() {
		return Title{ Title: title }, errors.New("Title is not valid")
	}

	// newTitle := Title{ Title: title }

	return Title{ Title: title }, nil
}

func (t *Title) GetValue () string {
	return t.Title
}

func (t *Title) noValid () bool {
	if 
		t.Title == "" ||
		len(t.Title) < 7 ||
		len(t.Title) > 255 {
		return true
	}

	return false;
}