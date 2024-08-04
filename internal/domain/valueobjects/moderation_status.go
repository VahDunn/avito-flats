package valueobjects

type ModerationStatus int

const (
	Created ModerationStatus = iota
	Approved
	Declined
	OnModeration
)
