package models

type GetTutorRes struct {
	ID                string    `json:"id"`
	Slug              string    `json:"slug"`
	Name              string    `json:"name"`
	Headline          string    `json:"headline"`
	Introduction      string    `json:"introduction"`
	PriceInfo         PriceInfo `json:"price_info"`
	TeachingLanguages []int     `json:"teaching_languages"`
}

type PriceInfo struct {
	Trial  float32 `json:"trial"`
	Normal float32 `json:"normal"`
}
