package models

// Tag represents the Tag model in Ghost's API
type Tag struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Description     string `json:"description"`
	FeatureImage    string `json:"feature_image"`
	Visibility      string `json:"visibility"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	URL             string `json:"url"`
}
