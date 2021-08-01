package types

type DBInterface interface {
	GetVideo(id string) *Video
	GetSnippet(id string) *Snippet
	InsertSnippet(s *Snippet)
	GetVideoMedia(id string) *VideoMedia
	InsertVideoMedia(vm *VideoMedia)
}

type MySQL struct{}

func (m *MySQL) GetVideo(id string) *Video {
	return nil
}
func (m *MySQL) GetSnippet(id string) *Snippet {
	return nil
}
func (m *MySQL) InsertSnippet(id *Snippet) {}
func (m *MySQL) GetVideoMedia(id string) *VideoMedia {
	return nil
}
func (m *MySQL) InsertVideoMedia(vm *VideoMedia) {}
