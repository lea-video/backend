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

type Movie struct {
	Id string `json:"id,omitempty"`

	Title *Title `json:"title,omitempty"`

	Sequel *Movie `json:"sequel,omitempty"`

	Cuts []*Media `json:"cuts,omitempty"`
}

func NewMovie(title *Title) *Movie {
	return &Movie{
		Id:    util.NewID(),
		Title: title,
		Cuts:  make([]*Media, 0),
	}
}

func (m *Movie) getID() string {
	return m.Id
}

func (*Movie) getType() string {
	return "Movie"
}