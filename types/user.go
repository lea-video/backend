package types

type User struct {
	Username    string        `json:"username"`
	DisplayName string        `json:"displayName"`
	Credentials []*Credential `json:"credentials"`
	Playlists   []*Playlist   `json:"playlists"`
	Snippets    []*Snippet    `json:"snippets"`
	Follows     []*Celebrity  `json:"follows"`
}

type Session struct {
	User          *User `json:"user"`
	IsDirectLogin bool  `json:"isDirectLogin"`
}

type Credential interface {
	GetSession() *Session
	IsValid() bool
}

type RememberMeToken struct {
	Token string `json:"token"`
}

func (rm *RememberMeToken) GetSession() *Session {
	return nil
}
func (rm *RememberMeToken) IsValid() bool {
	return false
}

type Password struct {
	Hash string `json:"hash"`
	Salt string `json:"salt"`
}

func (p *Password) GetSession() *Session {
	return nil
}
func (p *Password) IsValid() bool {
	return false
}
