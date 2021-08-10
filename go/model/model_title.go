/*
 * LEAV
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package model

import (
	util "github.com/lea-video/backend/go/utility"
)

type RawTitle struct {
	ID string `json:"id,omitempty"`

	Default string `json:"default,omitempty"`

	OtherIDs []string `json:"otherIDs,omitempty"`
}

type Title struct {
	*RawTitle

	Other []*TitleItem `json:"other,omitempty"`
}

type TitleItem struct {
	ID string `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Language string `json:"language,omitempty"`
}

func NewTitle(defTitle string) *Title {
	return &Title{
		RawTitle: &RawTitle{
			ID:       util.NewID(),
			Default:  defTitle,
			OtherIDs: make([]string, 0),
		},
		Other: make([]*TitleItem, 0),
	}
}

func (t *Title) AddTranslation(title, lang string) {
	e := TitleItem{
		ID:       util.NewID(),
		Title:    title,
		Language: lang,
	}
	t.OtherIDs = append(t.OtherIDs, e.ID)
	t.Other = append(t.Other, &e)
}

func (t *Title) GetTitle(lang string) string {
	for i := 0; i < len(t.Other); i++ {
		if t.Other[i].Language == lang {
			return t.Other[i].Title
		}
	}
	return t.Default
}

func (t *Title) getID() string {
	return t.ID
}
