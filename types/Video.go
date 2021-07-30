package types

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
	Title     *Title
	Length    uint
	AudioF    []*AudioFormat
	VideoF    []*VideoFormat
	SubtitleF []*SubtitleFormat
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
	Title  *Title
	Start  uint
	Length uint
	Parent *Video
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
