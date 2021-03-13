package domain

type User struct {
	ID      int64   `json:"id,omitempty"`
	FName   string  `json:"fname,omitempty"`
	City    string  `json:"city,omitempty"`
	Phone   string  `json:"phone,omitempty"`
	Height  float32 `json:"height,omitempty"`
	Married bool    `json:"married,omitempty"`
}
