package entity

import "time"

type URL struct {
	ID             int32     `json:"id"`
	ShortUri       string    `json:"short_uri"`
	UserID         int32     `json:"user_id"`
	RequestedCount int32     `json:"requested_count"`
	OriginalUrl    string    `json:"original_url"`
	ExpiresAt      time.Time `json:"expires_at"`
}

type CreateURIParams struct {
	Username    string `json:"-"`
	ShortURI    string `json:"short_uri"`
	OriginalUrl string `json:"original_url" binding:"required,min=5,url"`
}

type UpdateURIParams struct {
	Username    string `json:"-"`
	OldShortURI string `json:"old_short_uri" binding:"required"`
	NewShortURI string `json:"new_short_uri" binding:"required"`
}
