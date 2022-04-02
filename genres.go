package NovelUpdatesClient

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type GenreSearchResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Genres() ([]GenreSearchResult, error) {

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

	return findGenreResults(doc), nil
}

func findGenreResults(doc *goquery.Document) []GenreSearchResult {
	var results []GenreSearchResult
	doc.Find("div.g-cols:nth-child(24)").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			results = append(results, GenreSearchResult{
				Name:  findName(s),
				Value: findGenreIdAttr(s),
			})
		})
	})
	return results
}

func findGenreIdAttr(s *goquery.Selection) string {
	v, _ := s.Attr("genreid")
	return v
}
