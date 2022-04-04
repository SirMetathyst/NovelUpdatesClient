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
	urlRatingScaleKey           = "mrt"
	urlRatingKey                = "rt"
	urlNumberOfRatingsScaleKey  = "mrtc"
	urlNumberOfRatingsKey       = "rtc"
	urlReadersScaleKey          = "mrct"
	urlReadersKey               = "rct"
	urlFirstReleaseDateScaleKey = "mdtf"
	urlFirstReleaseDateKey      = "dtf"
	urlLastReleaseDateScaleKey  = "mdt"
	urlLastReleaseDateKey       = "dt"
	urlGenreOperatorKey         = "mgi"
	urlGenreIncludeKey          = "gi"
	urlGenreExcludeKey          = "ge"
	urlTagOperatorKey           = "mtgi"
	urlTagIncludeKey            = "tgi"
	urlTagExcludeKey            = "tge"
	urlStoryStatusKey           = "ss"
	urlSortKey                  = "sort"
	urlOrderKey                 = "order"
)

// TODO: this code is crap (but it works) so clean it up later

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

func buildRating(q *SearchQuery) string {
	if q != nil && q.Rating > 0 {
		v := fmt.Sprintf("%s=%d", urlRatingKey, q.Rating)
		if q.RatingScale == "" {
			v += fmt.Sprintf("&%s=%s", urlRatingScaleKey, ScaleMin)
		} else {
			v += fmt.Sprintf("&%s=%s", urlRatingScaleKey, q.RatingScale)
		}
		return v
	}
	return ""
}

func buildNumberOfRatings(q *SearchQuery) string {
	if q != nil && q.NumberOfRatings > 0 {
		v := fmt.Sprintf("%s=%d", urlNumberOfRatingsKey, q.NumberOfRatings)
		if q.NumberOfRatingsScale == "" {
			v += fmt.Sprintf("&%s=%s", urlNumberOfRatingsScaleKey, ScaleMin)
		} else {
			v += fmt.Sprintf("&%s=%s", urlNumberOfRatingsScaleKey, q.NumberOfRatingsScale)
		}
		return v
	}
	return ""
}

func buildReaders(q *SearchQuery) string {
	if q != nil && q.Readers > 0 {
		v := fmt.Sprintf("%s=%d", urlReadersKey, q.Readers)
		if q.ReadersScale == "" {
			v += fmt.Sprintf("&%s=%s", urlReadersScaleKey, ScaleMin)
		} else {
			v += fmt.Sprintf("&%s=%s", urlReadersScaleKey, q.ReadersScale)
		}
		return v
	}
	return ""
}

func buildFirstReleaseDate(q *SearchQuery) string {
	if q != nil && q.FirstReleaseDate != "" {
		v := fmt.Sprintf("%s=%s", urlFirstReleaseDateKey, q.FirstReleaseDate)
		if q.FirstReleaseDateScale == "" {
			v += fmt.Sprintf("&%s=%s", urlFirstReleaseDateScaleKey, ScaleMin)
		} else {
			v += fmt.Sprintf("&%s=%s", urlFirstReleaseDateScaleKey, q.FirstReleaseDateScale)
		}
		return v
	}
	return ""
}

func buildLastReleaseDate(q *SearchQuery) string {
	if q != nil && q.LastReleaseDate != "" {
		v := fmt.Sprintf("%s=%s", urlLastReleaseDateKey, q.LastReleaseDate)
		if q.LastReleaseDateScale == "" {
			v += fmt.Sprintf("&%s=%s", urlLastReleaseDateScaleKey, ScaleMin)
		} else {
			v += fmt.Sprintf("&%s=%s", urlLastReleaseDateScaleKey, q.LastReleaseDateScale)
		}
		return v
	}
	return ""
}

func buildGenre(q *SearchQuery) string {
	if q != nil && len(q.GenreInclude) > 0 || len(q.GenreExclude) > 0 {
		v := ""
		if len(q.GenreInclude) > 0 {
			v += fmt.Sprintf("%s=%s", urlGenreIncludeKey, strings.Join(q.GenreInclude, ","))
		}
		if len(q.GenreExclude) > 0 {
			v += fmt.Sprintf("&%s=%s", urlGenreExcludeKey, strings.Join(q.GenreExclude, ","))
		}
		if q.GenreOperator == "" {
			v += fmt.Sprintf("&%s=%s", urlGenreOperatorKey, OperatorAnd)
		} else {
			v += fmt.Sprintf("&%s=%s", urlGenreOperatorKey, q.GenreOperator)
		}
		return v
	}
	return ""
}

func buildTag(q *SearchQuery) string {
	if q != nil && len(q.TagInclude) > 0 || len(q.TagExclude) > 0 {
		v := ""
		if len(q.TagInclude) > 0 {
			v += fmt.Sprintf("%s=%s", urlTagIncludeKey, strings.Join(q.TagInclude, ","))
		}
		if len(q.TagExclude) > 0 {
			v += fmt.Sprintf("&%s=%s", urlTagExcludeKey, strings.Join(q.TagExclude, ","))
		}
		if q.TagOperator == "" {
			v += fmt.Sprintf("&%s=%s", urlTagOperatorKey, OperatorOr)
		} else {
			v += fmt.Sprintf("&%s=%s", urlTagOperatorKey, q.TagOperator)
		}
		return v
	}
	return ""
}

func buildStoryStatus(q *SearchQuery) string {
	if q != nil && q.StoryStatus != "" {
		return fmt.Sprintf("%s=%s", urlStoryStatusKey, q.StoryStatus)
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
	if q == nil {
		return fmt.Sprintf("%s=1", urlSeriesFinderKey)
	}
	var params []string
	params = append(params, fmt.Sprintf("%s=1", urlSeriesFinderKey))
	params = appendNonEmpty(params, buildNovelType(q))
	params = appendNonEmpty(params, buildLanguage(q))
	params = appendNonEmpty(params, buildChapters(q))
	params = appendNonEmpty(params, buildReleaseFrequency(q))
	params = appendNonEmpty(params, buildReviews(q))
	params = appendNonEmpty(params, buildRating(q))
	params = appendNonEmpty(params, buildNumberOfRatings(q))
	params = appendNonEmpty(params, buildReaders(q))
	params = appendNonEmpty(params, buildFirstReleaseDate(q))
	params = appendNonEmpty(params, buildLastReleaseDate(q))
	params = appendNonEmpty(params, buildGenre(q))
	params = appendNonEmpty(params, buildTag(q))
	params = appendNonEmpty(params, buildStoryStatus(q))
	params = appendNonEmpty(params, buildSort(q))
	params = appendNonEmpty(params, buildOrder(q))
	return strings.Join(params, "&")
}
