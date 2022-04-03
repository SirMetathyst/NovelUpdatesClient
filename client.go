package NovelUpdatesClient

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

var Default = NewClient()

type Interface interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
}

func NewClient() *http.Client {
	return &http.Client{Transport: NewTransport()}
}

func NewGetRequest(client Interface, url string) (*goquery.Document, error) {

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:98.0) Gecko/20100101 Firefox/98.0")

	// Request NovelUpdates Page
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, &ErrHTTP{Status: res.Status, StatusCode: res.StatusCode}
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
