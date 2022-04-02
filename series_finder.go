package NovelUpdatesClient

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type SeriesFinderSearchResult struct {
	Title string `json:"title"`
}

type SeriesFinderSearchResponse struct {
	Results []SeriesFinderSearchResult `json:"results"`
}

func SeriesFinder() (*SeriesFinderSearchResponse, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}

	// Create client
	client := &http.Client{Transport: tr}

	// Create request
	req, err := http.NewRequest("GET", "https://www.novelupdates.com/series-finder/?sf=1&sort=sdate&order=asc", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:98.0) Gecko/20100101 Firefox/98.0")

	// Request NovelUpdates Page
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return &SeriesFinderSearchResponse{Results: findSFSearchResults(doc)}, nil
}

func findSFSearchResults(doc *goquery.Document) []SeriesFinderSearchResult {
	var results []SeriesFinderSearchResult
	doc.Find("div.search_main_box_nu").Each(func(i int, s *goquery.Selection) {
		results = append(results, SeriesFinderSearchResult{
			Title: findSFTitle(s),
		})
	})
	return results
}

func findSFTitle(s *goquery.Selection) string {
	return s.Find(".search_title").Text()
}
