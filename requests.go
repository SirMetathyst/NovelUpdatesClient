package NovelUpdatesClient

import "github.com/PuerkitoBio/goquery"

func DoFetchGenresRequest(client Interface) (r []GenreResult, err error) {

	// Do request
	doc, err := NewGetRequest(client, "https://www.novelupdates.com/series-finder")
	if err != nil {
		return nil, err
	}

	doc.Find("div.g-cols:nth-child(24) a").Each(func(i int, s *goquery.Selection) {
		r = append(r, GenreResult{
			Name:  s.Text(),
			Value: s.AttrOr("genreid", ""),
		})
	})

	return r, nil
}

func DoFetchTagsRequest(client Interface) (r []TagResult, err error) {

	// Do request
	doc, err := NewGetRequest(client, "https://www.novelupdates.com/series-finder")
	if err != nil {
		return nil, err
	}

	// Find data
	doc.Find("select#tags_include option").Each(func(i int, s *goquery.Selection) {
		r = append(r, TagResult{
			Name:  s.Text(),
			Value: s.AttrOr("value", ""),
		})
	})

	return r, nil
}

func DoSearchRequest(client Interface) (r []SearchResult, err error) {

	// Do request
	doc, err := NewGetRequest(client, "https://www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=asc")
	if err != nil {
		return nil, err
	}

	// Find data
	doc.Find("div.search_main_box_nu").Each(func(i int, s *goquery.Selection) {
		r = append(r, SearchResult{Title: s.Find(".search_title").Text()})
	})

	return r, nil
}
