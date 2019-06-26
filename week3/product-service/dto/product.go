package dto

type Product struct {
	ID        int64       `json:"id,omitempty"`
	Category  interface{} `json:"category"`
	Name      string      `json:"name"`
	Image     string      `json:"image"`
	CreatedAt string      `json:"created_at,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty"`
}
