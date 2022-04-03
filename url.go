package NovelUpdatesClient

import (
	"fmt"
	"strings"
)

type OfString interface {
	[]NovelType
}

func toStringSlice[T OfString](input T) (r []string) {
	for _, n := range input {
		r = append(r, string(n))
	}
	return
}

func appendNonEmpty(dst []string, src string) []string {
	if src != "" {
		dst = append(dst, src)
	}
	return dst
}

func buildNovelType(q *SearchQuery) string {
	// nt=2443,26874,2444
	if q != nil && len(q.NovelType) > 0 {
		v := toStringSlice(q.NovelType)
		return fmt.Sprintf("nt=%s", strings.Join(v, ","))
	}
	return ""
}

func buildSort(q *SearchQuery) string {
	// sort=sdate
	if q != nil && q.SortBy != "" {
		return fmt.Sprintf("sort=%s", q.SortBy)
	}
	return fmt.Sprintf("sort=%s", SortByDefault)
}

func buildOrder(q *SearchQuery) string {
	// sort=sdate
	if q != nil && q.OrderBy != "" {
		return fmt.Sprintf("order=%s", q.OrderBy)
	}
	return fmt.Sprintf("order=%s", OrderDefault)
}

func BuildSearchStringFromQuery(q *SearchQuery) string {
	var params []string
	params = append(params, "?sf=1")
	params = appendNonEmpty(params, buildNovelType(q))
	params = appendNonEmpty(params, buildSort(q))
	params = appendNonEmpty(params, buildOrder(q))
	return strings.Join(params, "&")
}
