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

type Album struct {
	Id string `json:"id,omitempty"`

	Title *Title `json:"title,omitempty"`

	Songs []*Song `json:"songs,omitempty"`
}

func NewAlbum(title *Title) *Album {
	return &Album{
		Id:    util.NewID(),
		Title: title,
		Songs: make([]*Song, 0),
	}
}

func (a *Album) getID() string {
	return a.Id
}

func (*Album) getType() string {
	return "Album"
}