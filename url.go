package NovelUpdatesClient

import (
	"fmt"
	"strings"
)

const (
	urlSeriesFinderKey          = "sf"
	urlNovelTypeKey             = "nt"
	urlLanguageKey              = "org"
	urlChaptersScaleKey         = "mrl"
	urlChaptersKey              = "rl"
	urlReleaseFrequencyScaleKey = "mrf"
	urlReleaseFrequencyKey      = "rf"
	urlReviewsScaleKey          = "mrvc"
	urlReviewsKey               = "rvc"
	urlSortKey                  = "sort"
	urlOrderKey                 = "order"
)

func appendNonEmpty(dst []string, src string) []string {
	if src != "" {
		dst = append(dst, src)
	}
	return dst
}

func buildNovelType(q *SearchQuery) string {
	if q != nil && len(q.NovelType) > 0 {
		return fmt.Sprintf("%s=%s", urlNovelTypeKey, strings.Join(q.NovelType, ","))
	}
	return ""
}

func buildLanguage(q *SearchQuery) string {
	if q != nil && len(q.Language) > 0 {
		return fmt.Sprintf("%s=%s", urlLanguageKey, strings.Join(q.Language, ","))
	}
	return ""
}

func buildChapters(q *SearchQuery) string {
	if q != nil && q.Chapters > 0 {
		v := fmt.Sprintf("%s=%d", urlChaptersKey, q.Chapters)
		if q.ChaptersScale == "" {
			v += fmt.Sprintf("&%s=%s", urlChaptersScaleKey, ScaleMin)
		} else {
			v += fmt.Sprintf("&%s=%s", urlChaptersScaleKey, q.ChaptersScale)
		}
		return v
	}
	return ""
}

func buildReleaseFrequency(q *SearchQuery) string {
	if q != nil && q.ReleaseFrequency > 0 {
		v := fmt.Sprintf("%s=%d", urlReleaseFrequencyKey, q.ReleaseFrequency)
		if q.ReleaseFrequencyScale == "" {
			v += fmt.Sprintf("&%s=%s", urlReleaseFrequencyScaleKey, ScaleMax)
		} else {
			v += fmt.Sprintf("&%s=%s", urlReleaseFrequencyScaleKey, q.ReleaseFrequencyScale)
		}
		return v
	}
	return ""
}

func buildReviews(q *SearchQuery) string {
	if q != nil && q.Reviews > 0 {
		v := fmt.Sprintf("%s=%d", urlReviewsKey, q.Reviews)
		if q.ReviewsScale == "" {
			v += fmt.Sprintf("&%s=%s", urlReviewsScaleKey, ScaleMin)
		} else {
			v += fmt.Sprintf("&%s=%s", urlReviewsScaleKey, q.ReviewsScale)
		}
		return v
	}
	return ""
}

func buildSort(q *SearchQuery) string {
	if q != nil && q.SortBy != "" {
		return fmt.Sprintf("%s=%s", urlSortKey, q.SortBy)
	}
	return fmt.Sprintf("%s=%s", urlSortKey, SortByLastUpdated)
}

func buildOrder(q *SearchQuery) string {
	if q != nil && q.OrderBy != "" {
		return fmt.Sprintf("%s=%s", urlOrderKey, q.OrderBy)
	}
	return fmt.Sprintf("%s=%s", urlOrderKey, OrderDescending)
}

func buildSearchStringFromQuery(q *SearchQuery) string {
	var params []string
	params = append(params, fmt.Sprintf("%s=1", urlSeriesFinderKey))
	params = appendNonEmpty(params, buildNovelType(q))
	params = appendNonEmpty(params, buildLanguage(q))
	params = appendNonEmpty(params, buildChapters(q))
	params = appendNonEmpty(params, buildReleaseFrequency(q))
	params = appendNonEmpty(params, buildReviews(q))
	params = appendNonEmpty(params, buildSort(q))
	params = appendNonEmpty(params, buildOrder(q))
	return strings.Join(params, "&")
}
