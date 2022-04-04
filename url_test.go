package NovelUpdatesClient

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

func TestBuildSearchStringFromQuery(t *testing.T) {

	data := []struct {
		Query    SearchQuery
		Expected map[string][]string
	}{
		{
			Query: SearchQuery{
				NovelType: []string{
					NovelTypeLightNovel,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlNovelTypeKey:    {NovelTypeLightNovel},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				NovelType: []string{
					NovelTypeLightNovel,
					NovelTypeWebNovel,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlNovelTypeKey:    {NovelTypeLightNovel, NovelTypeWebNovel},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				NovelType: []string{
					NovelTypeLightNovel,
					NovelTypeWebNovel,
					NovelTypePublishedNovel,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlNovelTypeKey:    {NovelTypeLightNovel, NovelTypeWebNovel, NovelTypePublishedNovel},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Language: []string{
					LanguageChinese,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlLanguageKey:     {LanguageChinese},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Language: []string{
					LanguageChinese,
					LanguageFilipino,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlLanguageKey:     {LanguageChinese, LanguageFilipino},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Language: []string{
					LanguageChinese,
					LanguageFilipino,
					LanguageIndonesian,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlLanguageKey:     {LanguageChinese, LanguageFilipino, LanguageIndonesian},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Language: []string{
					LanguageChinese,
					LanguageFilipino,
					LanguageIndonesian,
					LanguageJapanese,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlLanguageKey:     {LanguageChinese, LanguageFilipino, LanguageIndonesian, LanguageJapanese},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Chapters: 10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlChaptersScaleKey: {ScaleMin},
				urlChaptersKey:      {"10"},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ChaptersScale: ScaleMin,
				Chapters:      10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlChaptersScaleKey: {ScaleMin},
				urlChaptersKey:      {"10"},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ChaptersScale: ScaleMax,
				Chapters:      10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlChaptersScaleKey: {ScaleMax},
				urlChaptersKey:      {"10"},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ReleaseFrequency: 10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:          {SeriesFinderEnabled},
				urlReleaseFrequencyScaleKey: {ScaleMax},
				urlReleaseFrequencyKey:      {"10"},
				urlSortKey:                  {SortByLastUpdated},
				urlOrderKey:                 {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ReleaseFrequencyScale: ScaleMin,
				ReleaseFrequency:      10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:          {SeriesFinderEnabled},
				urlReleaseFrequencyScaleKey: {ScaleMin},
				urlReleaseFrequencyKey:      {"10"},
				urlSortKey:                  {SortByLastUpdated},
				urlOrderKey:                 {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ReleaseFrequencyScale: ScaleMax,
				ReleaseFrequency:      10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:          {SeriesFinderEnabled},
				urlReleaseFrequencyScaleKey: {ScaleMax},
				urlReleaseFrequencyKey:      {"10"},
				urlSortKey:                  {SortByLastUpdated},
				urlOrderKey:                 {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Reviews: 10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlReviewsScaleKey: {ScaleMin},
				urlReviewsKey:      {"10"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ReviewsScale: ScaleMin,
				Reviews:      10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlReviewsScaleKey: {ScaleMin},
				urlReviewsKey:      {"10"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ReviewsScale: ScaleMax,
				Reviews:      10,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlReviewsScaleKey: {ScaleMax},
				urlReviewsKey:      {"10"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
	}

	for _, n := range data {
		urlString := buildSearchStringFromQuery(&n.Query)
		parsedUrl, _ := url.ParseQuery(urlString)
		t.Logf("URL: %s", urlString)

		for k, v := range n.Expected {
			actual := parsedUrl.Get(k)
			actualSplit := strings.Split(actual, ",")

			for _, vv := range v {
				assert.Containsf(t, actualSplit, vv, "for query param: %s", k)
			}
		}
	}
}
