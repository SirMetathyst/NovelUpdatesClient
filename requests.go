package NovelUpdatesClient

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func DoFetchNovelTypeRequest() (r []NovelTypeResult, err error) {

	// Create Document From URL
	doc, err := createDocumentFromURL("https://www.novelupdates.com/series-finder")
	if err != nil {
		return nil, err
	}

	// Find & Extract Data
	doc.Find(".rankfl > div:nth-child(3) a").Each(func(i int, s *goquery.Selection) {
		r = append(r, NovelTypeResult{
			Name:  s.Text(),
			Value: s.AttrOr("genreid", ""),
		})
	})

	return r, nil
}

func DoFetchLanguageRequest() (r []NovelTypeResult, err error) {

	// Create Document From URL
	doc, err := createDocumentFromURL("https://www.novelupdates.com/series-finder")
	if err != nil {
		return nil, err
	}

	// Find & Extract Data
	doc.Find("div.g-cols:nth-child(5) a").Each(func(i int, s *goquery.Selection) {
		r = append(r, NovelTypeResult{
			Name:  s.Text(),
			Value: s.AttrOr("genreid", ""),
		})
	})

	return r, nil
}

func DoFetchGenresRequest() (r []GenreResult, err error) {

	// Create Document From URL
	doc, err := createDocumentFromURL("https://www.novelupdates.com/series-finder")
	if err != nil {
		return nil, err
	}

	// Find & Extract Data
	doc.Find("div.g-cols:nth-child(24) a").Each(func(i int, s *goquery.Selection) {
		r = append(r, GenreResult{
			Name:  s.Text(),
			Value: s.AttrOr("genreid", ""),
		})
	})

	return r, nil
}

func DoFetchTagsRequest() (r []TagResult, err error) {

	doc, err := createDocumentFromURL("https://www.novelupdates.com/series-finder")
	if err != nil {
		return nil, err
	}

	// Find & Extract Data
	doc.Find("select#tags_include option").Each(func(i int, s *goquery.Selection) {
		r = append(r, TagResult{
			Name:  s.Text(),
			Value: s.AttrOr("value", ""),
		})
	})

	return r, nil
}

func DoSearchRequest(q *SearchQuery) (r []SearchResult, err error) {

	doc, err := createDocumentFromURL(fmt.Sprintf("https://www.novelupdates.com/series-finder/%s", buildParamStringFromSearchQuery(q)))
	if err != nil {
		return nil, err
	}

	// Find & Extract Data
	doc.Find("div.search_main_box_nu").Each(func(i int, s *goquery.Selection) {
		r = append(r, SearchResult{Title: s.Find(".search_title").Text()})
	})

	return r, nil
}
