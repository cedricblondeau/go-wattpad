package wattpad

// Story defines a simple story struct
type Story struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"user"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// StoriesEnvelope defines an envelope containing a list of stories
type StoriesEnvelope struct {
	Stories []Story `json:"stories"`
}
