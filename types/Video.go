package types

type Playable interface {
	GetStart() uint
	GetLength() uint
	GetFormat() []*Format
	doSnip(title Title, start uint, snipLength uint) *Snippet
}
type Video struct {
	Title   *Title
	Length  uint
	Formats []*Format
}

func (V *Video) GetStart() uint {
	return 0
}
func (V *Video) GetLength() uint {
	return V.Length
}
func (V *Video) GetFormat() []*Format {
	return V.Formats
}
func (V *Video) doSnip(title Title, start uint, snipLength uint) *Snippet {
	vid := Video{&title, snipLength, nil}
	return &Snippet{vid, V, start}
}

type Snippet struct {
	Video
	Parent *Video
	Start  uint
}

func (S *Snippet) GetStart() uint {
	return S.Start
}
func (S *Snippet) GetLength() uint {
	return S.Length
}
func (S *Snippet) GetFormat() []*Format {
	return S.Parent.Formats
}
func (S *Snippet) doSnip(title Title, start uint, snipLength uint) *Snippet {
	vid := Video{&title, snipLength, nil}
	return &Snippet{vid, S.Parent, S.Start + start}
}
