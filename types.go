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

const (
	StoryStatusDefault = StoryStatusAll
	SortByDefault      = SortByLastUpdated
	OrderByDefault     = OrderByDescending
)

func SlugStringToStoryStatusOrDefault(s string) StoryStatus {
	return SlugStringToStoryStatusOr(s, StoryStatusDefault)
}

func SlugStringToSortByOrDefault(s string) SortBy {
	return SlugStringToSortByOr(s, SortByDefault)
}

func SlugStringToOrderByOrDefault(s string) OrderBy {
	return SlugStringToOrderByOr(s, OrderByDefault)
}

const (
	ScaleMin = "min"
	ScaleMax = "max"
)

const (
	OperatorAnd = "and"
	OperatorOr  = "or"
)

type KeyValueResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SearchQuery struct {
	NovelType             NovelTypeList `json:"novel_type"`
	Language              LanguageList  `json:"language"`
	ChaptersScale         string        `json:"chapters_scale"`
	Chapters              int           `json:"chapters"`
	ReleaseFrequencyScale string        `json:"release_frequency_scale"`
	ReleaseFrequency      int           `json:"release_frequency"`
	ReviewsScale          string        `json:"reviews_scale"`
	Reviews               int           `json:"reviews"`
	RatingScale           string        `json:"rating_scale"`
	Rating                int           `json:"rating"`
	NumberOfRatingsScale  string        `json:"number_of_ratings_scale"`
	NumberOfRatings       int           `json:"number_of_ratings"`
	ReadersScale          string        `json:"readers_scale"`
	Readers               int           `json:"readers"`
	FirstReleaseDateScale string        `json:"first_release_date_scale"`
	FirstReleaseDate      string        `json:"first_release_date"`
	LastReleaseDateScale  string        `json:"last_release_date_scale"`
	LastReleaseDate       string        `json:"last_release_date"`
	GenreOperator         string        `json:"genre_operator"`
	GenreInclude          GenreList     `json:"genre_include"`
	GenreExclude          GenreList     `json:"genre_exclude"`
	TagOperator           string        `json:"tag_operator"`
	TagInclude            TagList       `json:"tag_include"`
	TagExclude            TagList       `json:"tag_exclude"`
	StoryStatus           StoryStatus   `json:"story_status"`
	SortBy                SortBy        `json:"sort_by"`
	OrderBy               OrderBy       `json:"order_by"`
	Page                  int           `json:"page"`
}

type SearchResult struct {
	ID                     string    `json:"id"`
	URL                    string    `json:"url"`
	Title                  string    `json:"title"`
	Chapters               int       `json:"chapters"`
	ReleaseFrequencyInDays float64   `json:"release_frequency_in_days"`
	Readers                int       `json:"readers"`
	Reviews                int       `json:"reviews"`
	LastUpdated            string    `json:"last_updated"`
	Genre                  GenreList `json:"genre"`
	Description            string    `json:"description"`
	ShortLanguage          string    `json:"short_language"`
	Rating                 float64   `json:"rating"`
}
