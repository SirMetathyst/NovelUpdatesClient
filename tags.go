package NovelUpdatesClient

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type TagsSearchResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Tags() ([]TagsSearchResult, error) {

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

	return findTagResults(doc), nil
}

func findTagResults(doc *goquery.Document) []TagsSearchResult {
	var results []TagsSearchResult
	doc.Find("select#tags_include").Each(func(i int, s *goquery.Selection) {
		s.Find("option").Each(func(i int, s *goquery.Selection) {
			results = append(results, TagsSearchResult{
				Name:  findName(s),
				Value: findValueAttr(s),
			})
		})
	})
	return results
}

func findValueAttr(s *goquery.Selection) string {
	v, _ := s.Attr("value")
	return v
}

func findName(s *goquery.Selection) string {
	return s.Text()
}
