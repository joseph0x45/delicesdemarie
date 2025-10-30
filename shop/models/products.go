package models

type Product struct {
	ID          string `json:"id" db:"id"`
	Label       string `json:"label" db:"label"`
	Variant     string `json:"variant" db:"variant"`
	Price       int    `json:"price" db:"price"`
	Picture     string `json:"picture" db:"picture"`
	Description string `json:"description" db:"description"`
	Available   bool   `json:"available" db:"available"`
	Brand       string `json:"brand" db:"brand"`
}
