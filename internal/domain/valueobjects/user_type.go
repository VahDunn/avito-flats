package valueobjects

type UserType int

const (
	Client UserType = iota
	Moderator
)

type Token struct {
	Token string `json:"token"`
}
