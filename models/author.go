// Package models contains representations of the various objects in the Ghost API
package models

// Author represents the Author model in Ghost's API
type Author struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	ProfileImage    string `json:"profile_image"`
	CoverImage      string `json:"cover_image"`
	Bio             string `json:"bio"`
	Website         string `json:"website"`
	Location        string `json:"location"`
	Facebook        string `json:"facebook"`
	Twitter         string `json:"twitter"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	URL             string `json:"url"`
}
