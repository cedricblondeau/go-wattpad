package wattpad

// Category defines a simple category struct
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CategoriesEnvelope defines an envelope containing a list of categories
type CategoriesEnvelope struct {
	Categories []Category `json:"categories"`
}
