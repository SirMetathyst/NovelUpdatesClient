package NovelUpdatesClient

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func DoFetchNovelTypeRequestWith(req Requester) (r []NovelTypeResult, err error) {

	// Create Document From URL
	doc, err := createDocumentFromRequest(req, "https://www.novelupdates.com/series-finder")
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

func DoFetchLanguageRequestWith(req Requester) (r []NovelTypeResult, err error) {

	// Create Document From URL
	doc, err := createDocumentFromRequest(req, "https://www.novelupdates.com/series-finder")
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

func DoFetchGenresRequestWith(req Requester) (r []GenreResult, err error) {

	// Create Document From URL
	doc, err := createDocumentFromRequest(req, "https://www.novelupdates.com/series-finder")
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

func DoFetchTagsRequestWith(req Requester) (r []TagResult, err error) {

	doc, err := createDocumentFromRequest(req, "https://www.novelupdates.com/series-finder")
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

func DoSearchRequestWith(req Requester, q *SearchQuery) (r []SearchResult, err error) {

	doc, err := createDocumentFromRequest(req, fmt.Sprintf("https://www.novelupdates.com/series-finder/?%s", buildSearchStringFromQuery(q)))
	if err != nil {
		return nil, err
	}

	// Find & Extract Data
	doc.Find("div.search_main_box_nu").Each(func(i int, s *goquery.Selection) {

		Title := "ERR"
		{
			oTitle := s.Find(".search_title").Text()
			oTitle = strings.Trim(oTitle, " ")
			if oTitle != "" {
				Title = oTitle
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
		var Genre []string
		{
			s.Find(".search_genre a").Each(func(i int, selection *goquery.Selection) {
				oGenre := selection.Text()
				oGenre = strings.Trim(oGenre, " ")
				if oGenre != "" {
					Genre = append(Genre, oGenre)
				}
			})
		}
		///////////////////////////////////////////////////////////
		Rating := -1.0
		{
			oRating := s.Find(".search_ratings").Text()
			oRatingSplit := strings.Split(oRating, "(")
			if len(oRatingSplit) == 2 {
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
				return !s.Is(".search_title,.search_stats,.search_genre,.dots,.list")
			}).Text()
			oDescription = strings.Replace(oDescription, "<<less", "", -1)
			oDescription = strings.Trim(oDescription, " \n")
			if oDescription != "" {
				Description = oDescription
			}
		}

		r = append(r, SearchResult{
			Title:                  Title,
			Chapters:               Chapters,
			ReleaseFrequencyInDays: ReleaseFrequency,
			Readers:                Readers,
			Reviews:                Reviews,
			LastUpdated:            LastUpdated,
			Genre:                  Genre,
			Description:            Description,
			Rating:                 Rating,
		})
	})

	return r, nil
}
