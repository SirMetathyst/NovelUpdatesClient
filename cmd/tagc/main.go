package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type TagsSearchResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TagsResponse struct {
	Results []TagsSearchResult `json:"results"`
}

func Tags() (*TagsResponse, error) {

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

	return &TagsResponse{Results: findTagResults(doc)}, nil
}

func findTagResults(doc *goquery.Document) []TagsSearchResult {
	var results []TagsSearchResult
	doc.Find("select#tags_include").Each(func(i int, s *goquery.Selection) {
		s.Find("option").Each(func(i int, s *goquery.Selection) {
			results = append(results, TagsSearchResult{
				Name:  findName(s),
				Value: findValue(s),
			})
		})
	})
	return results
}

func findValue(s *goquery.Selection) string {
	v, _ := s.Attr("value")
	return v
}

func findName(s *goquery.Selection) string {
	return s.Text()
}

func normalisedName(n string) string {
	n = strings.Replace(n, " ", "", -1)
	n = strings.Replace(n, "-", "", -1)
	n = strings.Replace(n, "/", "", -1)
	n = strings.Replace(n, "'", "", -1)
	return n
}

func main() {

	flagPackageName := flag.String("package", "NovelUpdatesClient", "The name of the package for the generated output")

	resp, _ := Tags()

	b := strings.Builder{}

	//////// Header
	b.WriteString("// Code generated by tagc. DO NOT EDIT.\n")
	b.WriteString(fmt.Sprintf("package %s\n\n", *flagPackageName))
	b.WriteString("type Tag string\n\n")

	//////// Tags
	b.WriteString(fmt.Sprintf("// Tags: Total(%d)\n", len(resp.Results)))
	b.WriteString("const (\n")

	for _, result := range resp.Results {
		b.WriteString(fmt.Sprintf("\tTag%s Tag = \"%s\"\n", normalisedName(result.Name), result.Value))
	}

	b.WriteString(")\n")

	//////// TagToName
	b.WriteString("var (\n")
	b.WriteString("\tTagToName = map[Tag]string{\n")

	for _, result := range resp.Results {
		b.WriteString(fmt.Sprintf("\tTag%s:\"%s\",\n", normalisedName(result.Name), result.Name))
	}

	b.WriteString("}\n")
	b.WriteString(")\n\n")

	///////// NameToTag
	b.WriteString("var (\n")
	b.WriteString("\tNameToTag = map[string]Tag{\n")

	for _, result := range resp.Results {
		b.WriteString(fmt.Sprintf("\t\"%s\":Tag%s,\n", result.Name, normalisedName(result.Name)))
	}

	b.WriteString("}\n")
	b.WriteString(")")

	fmt.Println(b.String())
}
