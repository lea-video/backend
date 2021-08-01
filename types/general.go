package types

import "fmt"

type Title struct {
	Title     string `json:"title"`
	Language  string `json:"lang"`
	PrevTitle *Title `json:"prev"`
	NextTitle *Title `json:"next"`
}

func (t *Title) String() string {
	return fmt.Sprintf("%s(%s)", t.Title, t.Language)
}

type Tag struct {
	Title *Title `json:"title"`
}
