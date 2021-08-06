/*
 * LEAV
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package leav

type Song struct {
	Title *Title `json:"title,omitempty"`

	Artists []*Celebrity `json:"artists,omitempty"`

	Featuring []*Celebrity `json:"featuring,omitempty"`

	Variants []*Media `json:"variants,omitempty"`

	Covers []*SongCovers `json:"covers,omitempty"`

	CleanAudio *Media `json:"cleanAudio,omitempty"`

	DefaultVideo *Media `json:"defaultVideo,omitempty"`
}

func (*Song) getType() string {
	return "Song"
}

func NewSong(title *Title) *Song {
	return &Song{
		Title:     title,
		Artists:   make([]*Celebrity, 0),
		Featuring: make([]*Celebrity, 0),
		Variants:  make([]*Media, 0),
		Covers:    make([]*SongCovers, 0),
	}
}
