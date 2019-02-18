package models

import "time"

// Post represents the Post model in Ghost's API
type Post struct {
	ID                 string    `json:"id"`
	UUID               string    `json:"uuid"`
	Title              string    `json:"title"`
	Slug               string    `json:"slug"`
	HTML               string    `json:"html"`
	CommentID          string    `json:"comment_id"`
	FeatureImage       string    `json:"feature_image"`
	Featured           bool      `json:"featured"`
	Page               bool      `json:"page"`
	MetaTitle          string    `json:"meta_title"`
	MetaDescription    string    `json:"meta_description"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	PublishedAt        time.Time `json:"published_at"`
	CustomExcerpt      string    `json:"custom_excerpt"`
	Excerpt            string    `json:"excerpt"`
	CodeInjectionHead  string    `json:"codeinjection_head"`
	CodeInjectionFoot  string    `json:"codeinjection_foot"`
	OgImage            string    `json:"og_image"`
	OgTitle            string    `json:"og_title"`
	OgDescription      string    `json:"og_description"`
	TwitterImage       string    `json:"twitter_image"`
	TwitterTitle       string    `json:"twitter_title"`
	TwitterDescription string    `json:"twitter_description"`
	CustomTemplate     string    `json:"custom_template"`
	URL                string    `json:"url"`

	PrimaryAuthor *Author   `json:"primary_author"`
	PrimaryTag    *Tag      `json:"primary_tag"`
	Authors       []*Author `json:"authors"`
	Tags          []*Tag    `json:"tags"`
}
