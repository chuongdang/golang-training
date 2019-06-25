package entities

type Product struct {
	ID         int64  `db:"id_product" json:"id,omitempty"`
	FkCategory string `db:"fk_category" json:"category_id"`
	Name       string `db:"name" json:"name"`
	Image      string `db:"image" json:"image"`
	CreatedAt  string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt  string `db:"updated_at" json:"updated_at,omitempty"`
}
