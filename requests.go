package NovelUpdatesClient

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
)

func doFetchKeyValueRequestWith(req Requester, selector string, valueAttr string) (r []KeyValueResult, err error) {

	// Create Document From URL
	doc, err := createDocumentFromRequest(req, "https://www.novelupdates.com/series-finder")
	if err != nil {
		return nil, err
	}

	// Find & Extract Data
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		r = append(r, KeyValueResult{
			Name:  s.Text(),
			Value: s.AttrOr(valueAttr, ""),
		})
	})

	return r, nil
}

func DoFetchNovelTypeRequestWith(req Requester) (r []KeyValueResult, err error) {
	return doFetchKeyValueRequestWith(req, ".rankfl > div:nth-child(3) a", "genreid")
}

func DoFetchLanguageRequestWith(req Requester) (r []KeyValueResult, err error) {
	return doFetchKeyValueRequestWith(req, "div.g-cols:nth-child(5) a", "genreid")
}

func DoFetchGenresRequestWith(req Requester) (r []KeyValueResult, err error) {
	return doFetchKeyValueRequestWith(req, "div.g-cols:nth-child(24) a", "genreid")
}

func DoFetchTagsRequestWith(req Requester) (r []KeyValueResult, err error) {
	return doFetchKeyValueRequestWith(req, "select#tags_include option", "value")
}

func DoFetchStoryStatusRequestWith(req Requester) (r []KeyValueResult, err error) {
	return doFetchKeyValueRequestWith(req, ".storystatus option", "value")
}

func DoFetchSortByRequestWith(req Requester) (r []KeyValueResult, err error) {
	return doFetchKeyValueRequestWith(req, ".sortresults option", "value")
}

func DoFetchOrderByRequestWith(req Requester) (r []KeyValueResult, err error) {
	return doFetchKeyValueRequestWith(req, ".sortorder option", "value")
}

func DoSearchRequestWith(req Requester, q *SearchQuery) (r []SearchResult, err error) {

	doc, err := createDocumentFromRequest(req, fmt.Sprintf("https://www.novelupdates.com/series-finder/?%s", buildSearchStringFromQuery(q)))
	if err != nil {
		return nil, err
	}

	// Find & Extract Data
	doc.Find("div.search_main_box_nu").Each(func(i int, s *goquery.Selection) {

		///////////////////////////////////////////////////////////
		Title := "ERR"
		{
			oTitle := s.Find(".search_title").Text()
			oTitle = strings.Trim(oTitle, " ")
			if oTitle != "" {
				Title = oTitle
			}
		}
		///////////////////////////////////////////////////////////
		ID := "ERR"
		URL := "ERR"
		{
			oURL := s.Find(".search_title a").AttrOr("href", "")
			oURL = strings.Trim(oURL, " ")
			if _, err := url.ParseRequestURI(oURL); oURL != "" && err == nil {
				URL = oURL
			}
			oURLSplit := strings.Split(oURL, "/")
			if len(oURLSplit) == 6 {
				if oURL != "" {
					ID = oURLSplit[4]
				}
			}
		}
		///////////////////////////////////////////////////////////
		Chapters := -1
		{
			oChapters := s.Find(".search_stats span:nth-child(1)").Text()
			oChapters = strings.Replace(oChapters, "Chapters", "", -1)
			oChapters = strings.Trim(oChapters, " ")
			if Chapters, err = strconv.Atoi(oChapters); err != nil {
				Chapters = -1
			}
		}
		///////////////////////////////////////////////////////////
		ReleaseFrequency := -1.0
		{
			oReleaseFrequency := s.Find(".search_stats span:nth-child(2)").Text()
			oReleaseFrequency = strings.Replace(oReleaseFrequency, "Day(s)", "", -1)
			oReleaseFrequency = strings.Replace(oReleaseFrequency, "Every", "", -1)
			oReleaseFrequency = strings.Trim(oReleaseFrequency, " ")
			if ReleaseFrequency, err = strconv.ParseFloat(oReleaseFrequency, 64); err != nil {
				ReleaseFrequency = -1.0
			}
		}
		///////////////////////////////////////////////////////////
		Readers := -1
		{
			oReaders := s.Find(".search_stats span:nth-child(3)").Text()
			oReaders = strings.Replace(oReaders, "Readers", "", -1)
			oReaders = strings.Trim(oReaders, " ")
			if Readers, err = strconv.Atoi(oReaders); err != nil {
				Readers = -1.0
			}
		}
		///////////////////////////////////////////////////////////
		Reviews := -1
		{
			oReviews := s.Find(".search_stats span:nth-child(4)").Text()
			oReviews = strings.Replace(oReviews, "Reviews", "", -1)
			oReviews = strings.Trim(oReviews, " ")
			if Reviews, err = strconv.Atoi(oReviews); err != nil {
				Reviews = -1.0
			}
		}
		///////////////////////////////////////////////////////////
		LastUpdated := "00-00-0000"
		{
			oLastUpdated := s.Find(".search_stats span:nth-child(5)").Text()
			oLastUpdated = strings.Trim(oLastUpdated, " ")
			if oLastUpdated != "" {
				LastUpdated = oLastUpdated
			}
		}
		///////////////////////////////////////////////////////////
		var genreList GenreList
		{
			s.Find(".search_genre a").Each(func(i int, selection *goquery.Selection) {
				oGenre := selection.Text()
				oGenre = strings.Trim(oGenre, " ")
				if genre, ok := DisplayStringToGenre[oGenre]; ok && oGenre != "" {
					genreList = append(genreList, genre)
				}
			})
		}
		///////////////////////////////////////////////////////////
		Rating := -1.0
		ShortLanguage := "ERR"
		{
			oRating := s.Find(".search_ratings").Text()
			oRatingSplit := strings.Split(oRating, " ")

			if len(oRatingSplit) == 2 {
				oShortLanguage := oRatingSplit[0]
				oShortLanguage = strings.Trim(oShortLanguage, " ")
				if oShortLanguage != "" {
					ShortLanguage = oShortLanguage
				}
				oRating = strings.Trim(oRatingSplit[1], "()")
				if Rating, err = strconv.ParseFloat(oRating, 64); err != nil {
					Rating = -1.0
				}
			}
		}
		///////////////////////////////////////////////////////////
		Description := "ERR"
		{
			oDescription := s.Find(".search_body_nu").Contents().FilterFunction(func(i int, s *goquery.Selection) bool {
				return !s.Is(".search_title,.search_stats,.search_genre,div,.dots,.list")
			}).Text()
			oDescription = strings.Replace(oDescription, "<<less", "", -1)
			oDescription = strings.Trim(oDescription, " \n")
			if oDescription != "" {
				Description = oDescription
			}
		}

		r = append(r, SearchResult{
			Title:                  Title,
			ID:                     ID,
			URL:                    URL,
			Chapters:               Chapters,
			ReleaseFrequencyInDays: ReleaseFrequency,
			Readers:                Readers,
			Reviews:                Reviews,
			LastUpdated:            LastUpdated,
			Genre:                  genreList,
			Description:            Description,
			ShortLanguage:          ShortLanguage,
			Rating:                 Rating,
		})
	})

	return r, nil
}
