package models

type SentenceTypes string

const (
	Comprehension SentenceTypes = "comprehension"
	Basic         SentenceTypes = "basic"
)

type SentenceJSON struct {
	SentenceType SentenceTypes `json:"type"`
	Question     string        `json:"question"`
	FollowUp     string        `json:"follow_up"`
	Correct      bool          `json:"correct"`
}
