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

type Title struct {
	Id string `json:"id,omitempty"`

	Default string `json:"default,omitempty"`

	Other []*TitleItem `json:"other,omitempty"`
}

type TitleItem struct {
	Title string `json:"title,omitempty"`

	Language string `json:"language,omitempty"`
}

func NewTitle(defTitle string) *Title {
	return &Title{
		Id:      util.NewID(),
		Default: defTitle,
		Other:   make([]*TitleItem, 0),
	}
}

func (t *Title) AddTranslation(title, lang string) {
	e := TitleItem{
		Title:    title,
		Language: lang,
	}
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