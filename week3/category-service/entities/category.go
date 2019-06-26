package entities

type Category struct {
	ID        int64  `db:"id_category" json:"id,omitempty"`
	Name      string `db:"name" json:"name"`
	Image     string `db:"image" json:"image"`
	CreatedAt string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt string `db:"updated_at" json:"updated_at,omitempty"`
}
