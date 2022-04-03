package NovelUpdatesClient

import "fmt"

type GenreResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TagResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SearchResult struct {
	Title string `json:"title"`
}

type ErrHTTP struct {
	StatusCode int
	Status     string
}

func (s *ErrHTTP) Error() string {
	return fmt.Sprintf("ErrHTTP: %d:%s", s.StatusCode, s.Status)
}
