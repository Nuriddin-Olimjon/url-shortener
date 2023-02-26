package entity

type Url struct {
	ID             int32  `json:"id"`
	ShortUri       string `json:"short_uri"`
	UserID         int32  `json:"user_id"`
	RequestedCount int32  `json:"requested_count"`
	OriginalUrl    string `json:"original_url"`
}
