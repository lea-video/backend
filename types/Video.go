package types

import "fmt"

type Playable interface {
	GetStart() uint
	GetLength() uint
	GetAudioFormat() []*AudioFormat
	GetVideoFormat() []*VideoFormat
	GetSubtitleFormat() []*SubtitleFormat
	IsSnippet() bool
	DoSnip(title Title, snipStart uint, snipLength uint) *Snippet
}

type Video struct {
	Title     *Title            `json:"title"`
	Length    uint              `json:"length"`
	AudioF    []*AudioFormat    `json:"audioFormats"`
	VideoF    []*VideoFormat    `json:"videoFormats"`
	SubtitleF []*SubtitleFormat `json:"subtitleFormats"`
}

func (v *Video) String() string {
	return fmt.Sprintf("Video with title '%s' and Length of %dsec", v.Title, v.Length)
}

func (V *Video) GetStart() uint {
	return 0
}
func (V *Video) GetLength() uint {
	return V.Length
}
func (V *Video) GetAudioFormat() []*AudioFormat {
	return V.AudioF
}
func (V *Video) GetVideoFormat() []*VideoFormat {
	return V.VideoF
}
func (V *Video) GetSubtitleFormat() []*SubtitleFormat {
	return V.SubtitleF
}
func (V *Video) DoSnip(title Title, snipStart uint, snipLength uint) *Snippet {
	return &Snippet{
		Title:  &title,
		Start:  snipStart,
		Length: snipLength,
		Parent: V,
	}
}

type Snippet struct {
	Title  *Title `json:"title"`
	Start  uint   `json:"start"`
	Length uint   `json:"length"`
	Parent *Video `json:"parent"`
}

func (S *Snippet) GetStart() uint {
	return S.Start
}
func (S *Snippet) GetLength() uint {
	return S.Length
}
func (S *Snippet) GetAudioFormat() []*AudioFormat {
	return S.Parent.AudioF
}
func (S *Snippet) GetVideoFormat() []*VideoFormat {
	return S.Parent.VideoF
}
func (S *Snippet) GetSubtitleFormat() []*SubtitleFormat {
	return S.Parent.SubtitleF
}
func (S *Snippet) DoSnip(title Title, snipStart uint, snipLength uint) *Snippet {
	return &Snippet{
		Title:  &title,
		Start:  S.Start + snipStart,
		Length: snipLength,
		Parent: S.Parent,
	}
}
