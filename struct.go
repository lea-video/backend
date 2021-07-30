package backend

// Interface to be implemented by Database
type PersistentStorage interface {
	SaveVideo()
	LoadVideo() Video
}

// Programm Data
type File struct {
	id        string
	location  string
	encoding  string
	size      uint
	isDefault bool
}

func (m *File) isAudio() bool {
	return false
}

func (m *File) isVideo() bool {
	return false
}

func (m *File) isSubtitle() bool {
	return false
}

type AudioFile struct {
	File
	bitRate uint
}

type VideoFile struct {
	File
	height uint
	width  uint
}

type SubtitleFile struct {
	File
	language string
}

type Title struct {
	prev     *Title
	title    string
	language string
}

type Session struct {
	isDirectLogin bool
}

type Credential interface {
	isValid() bool
	getSession() *Session
}

type User struct {
	displayName string
	username    string
	creds       []*Credential
}

type Tag struct {
	title Title
}

type Video struct {
	title     Title
	medias    []*File
	variation []*Variation
	tags      []*Tag
}

type Variation struct {
	music
	movie
	videos []*Video
	tags   []*Tag
}

type music struct {
	title  Title
	artist []*Celebrity
}

type movie struct {
	title     Title
	actor     []*Celebrity
	regisseur []*Celebrity
}

type Snippet struct {
	Video
	title  Title
	parent Video
	start  uint
	end    uint
}

type Playlist struct {
	title Title
	items []*Video
}

type Celebrity struct{}
