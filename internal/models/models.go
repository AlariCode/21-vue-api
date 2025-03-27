package models

type Category struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Alias     string     `json:"alias"`
	Bookmarks []Bookmark `json:"bookmarks"`
}

type Bookmark struct {
	ID         int64  `json:"id"`
	CategoryID int64  `json:"category_id"`
	URL        string `json:"url"`
	Title      string `json:"title,omitempty"`
	Image      string `json:"image,omitempty"`
	CreatedAt  string `json:"created_at"`
}
