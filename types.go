package NovelUpdatesClient

import (
	"io"
)

type Requester interface {
	Request(url string) (io.ReadCloser, error)
}

type RequesterFunc func(url string) (io.ReadCloser, error)

func (f RequesterFunc) Request(url string) (io.ReadCloser, error) {
	return f(url)
}

type Scale string

const (
	ScaleMin Scale = "min"
	ScaleMax Scale = "max"
)

func (s Scale) String() string {
	return string(s)
}

func SlugStringToScaleOr(s string, def Scale) Scale {
	if Scale(s) == ScaleMin {
		return ScaleMin
	} else if Scale(s) == ScaleMax {
		return ScaleMax
	}
	return def
}

type Operator string

const (
	OperatorAnd Operator = "and"
	OperatorOr  Operator = "or"
)

func (s Operator) String() string {
	return string(s)
}

func SlugStringToOperatorOr(s string, def Operator) Operator {
	if Operator(s) == OperatorOr {
		return OperatorOr
	} else if Operator(s) == OperatorAnd {
		return OperatorAnd
	}
	return def
}

type KeyValueResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SearchQuery struct {
	NovelType             NovelTypeList `json:"novel_type"`
	Language              LanguageList  `json:"language"`
	ChaptersScale         Scale         `json:"chapters_scale"`
	Chapters              int           `json:"chapters"`
	ReleaseFrequencyScale Scale         `json:"release_frequency_scale"`
	ReleaseFrequency      int           `json:"release_frequency"`
	ReviewsScale          Scale         `json:"reviews_scale"`
	Reviews               int           `json:"reviews"`
	RatingScale           Scale         `json:"rating_scale"`
	Rating                int           `json:"rating"`
	NumberOfRatingsScale  Scale         `json:"number_of_ratings_scale"`
	NumberOfRatings       int           `json:"number_of_ratings"`
	ReadersScale          Scale         `json:"readers_scale"`
	Readers               int           `json:"readers"`
	FirstReleaseDateScale Scale         `json:"first_release_date_scale"`
	FirstReleaseDate      string        `json:"first_release_date"`
	LastReleaseDateScale  Scale         `json:"last_release_date_scale"`
	LastReleaseDate       string        `json:"last_release_date"`
	GenreOperator         Operator      `json:"genre_operator"`
	GenreInclude          GenreList     `json:"genre_include"`
	GenreExclude          GenreList     `json:"genre_exclude"`
	TagOperator           Operator      `json:"tag_operator"`
	TagInclude            TagList       `json:"tag_include"`
	TagExclude            TagList       `json:"tag_exclude"`
	StoryStatus           StoryStatus   `json:"story_status"`
	SortBy                SortBy        `json:"sort_by"`
	OrderBy               OrderBy       `json:"order_by"`
	Page                  int           `json:"page"`
}

type SearchResult struct {
	ID                     string    `json:"id,omitempty"`
	URL                    string    `json:"url,omitempty"`
	Title                  string    `json:"title,omitempty"`
	Chapters               int       `json:"chapters,omitempty"`
	ReleaseFrequencyInDays float64   `json:"release_frequency_in_days,omitempty"`
	Readers                int       `json:"readers,omitempty"`
	Reviews                int       `json:"reviews,omitempty"`
	LastUpdated            string    `json:"last_updated,omitempty"`
	Genre                  GenreList `json:"genre,omitempty"`
	Description            string    `json:"description,omitempty"`
	ShortLanguage          string    `json:"short_language,omitempty"`
	Rating                 float64   `json:"rating,omitempty"`
}
