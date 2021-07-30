package types

type User struct {
	Username    string
	DisplayName string
	Credentials []*Credential
	Playlists   []*Playlist
	Snippets    []*Snippet
	Follows     []*Celebrity
}

type Session struct {
	User          *User
	IsDirectLogin bool
}

type Credential interface {
	GetSession() *Session
	IsValid() bool
}

type RememberMeToken struct {
	Token string
}

func (rm *RememberMeToken) GetSession() *Session {
	return nil
}
func (rm *RememberMeToken) IsValid() bool {
	return false
}

type Password struct {
	Hash string
	Salt string
}

func (p *Password) GetSession() *Session {
	return nil
}
func (p *Password) IsValid() bool {
	return false
}
