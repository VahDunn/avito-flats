package valueobjects

type UserType int

const (
	Client UserType = iota
	Moderator
)
