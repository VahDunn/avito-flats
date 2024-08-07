package valueobjects

type ModerationStatus int64

const (
	Created ModerationStatus = iota
	Approved
	Declined
	OnModeration
)
