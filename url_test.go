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
		Expected map[string][]interface{}
	}{
		{
			Query: SearchQuery{
				NovelType: NovelTypeList{
					NovelTypeLightNovel,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlNovelTypeKey:    {string(NovelTypeLightNovel)},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				NovelType: NovelTypeList{
					NovelTypeLightNovel,
					NovelTypeWebNovel,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlNovelTypeKey:    {string(NovelTypeLightNovel), string(NovelTypeWebNovel)},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				NovelType: NovelTypeList{
					NovelTypeLightNovel,
					NovelTypeWebNovel,
					NovelTypePublishedNovel,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlNovelTypeKey:    {string(NovelTypeLightNovel), string(NovelTypeWebNovel), string(NovelTypePublishedNovel)},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				Language: LanguageList{
					LanguageChinese,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlLanguageKey:     {string(LanguageChinese)},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				Language: LanguageList{
					LanguageChinese,
					LanguageFilipino,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlLanguageKey:     {string(LanguageChinese), string(LanguageFilipino)},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				Language: LanguageList{
					LanguageChinese,
					LanguageFilipino,
					LanguageIndonesian,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlLanguageKey:     {string(LanguageChinese), string(LanguageFilipino), string(LanguageIndonesian)},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				Language: LanguageList{
					LanguageChinese,
					LanguageFilipino,
					LanguageIndonesian,
					LanguageJapanese,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlLanguageKey:     {string(LanguageChinese), string(LanguageFilipino), string(LanguageIndonesian), string(LanguageJapanese)},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				Chapters: 10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlChaptersScaleKey: {ScaleMin},
				urlChaptersKey:      {"10"},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ChaptersScale: ScaleMin,
				Chapters:      10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlChaptersScaleKey: {ScaleMin},
				urlChaptersKey:      {"10"},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ChaptersScale: ScaleMax,
				Chapters:      10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlChaptersScaleKey: {ScaleMax},
				urlChaptersKey:      {"10"},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ReleaseFrequency: 10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:          {urlSeriesFinderEnabled},
				urlReleaseFrequencyScaleKey: {ScaleMax},
				urlReleaseFrequencyKey:      {"10"},
				urlSortKey:                  {SortByLastUpdated.String()},
				urlOrderKey:                 {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ReleaseFrequencyScale: ScaleMin,
				ReleaseFrequency:      10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:          {urlSeriesFinderEnabled},
				urlReleaseFrequencyScaleKey: {ScaleMin},
				urlReleaseFrequencyKey:      {"10"},
				urlSortKey:                  {SortByLastUpdated.String()},
				urlOrderKey:                 {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ReleaseFrequencyScale: ScaleMax,
				ReleaseFrequency:      10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:          {urlSeriesFinderEnabled},
				urlReleaseFrequencyScaleKey: {ScaleMax},
				urlReleaseFrequencyKey:      {"10"},
				urlSortKey:                  {SortByLastUpdated.String()},
				urlOrderKey:                 {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				Reviews: 10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlReviewsScaleKey: {ScaleMin},
				urlReviewsKey:      {"10"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ReviewsScale: ScaleMin,
				Reviews:      10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlReviewsScaleKey: {ScaleMin},
				urlReviewsKey:      {"10"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ReviewsScale: ScaleMax,
				Reviews:      10,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlReviewsScaleKey: {ScaleMax},
				urlReviewsKey:      {"10"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				Rating: 5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlRatingScaleKey:  {ScaleMin},
				urlRatingKey:       {"5"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				RatingScale: ScaleMin,
				Rating:      5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlRatingScaleKey:  {ScaleMin},
				urlRatingKey:       {"5"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				RatingScale: ScaleMax,
				Rating:      5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlRatingScaleKey:  {ScaleMax},
				urlRatingKey:       {"5"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				NumberOfRatings: 5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:         {urlSeriesFinderEnabled},
				urlNumberOfRatingsScaleKey: {ScaleMin},
				urlNumberOfRatingsKey:      {"5"},
				urlSortKey:                 {SortByLastUpdated.String()},
				urlOrderKey:                {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				NumberOfRatingsScale: ScaleMin,
				NumberOfRatings:      5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:         {urlSeriesFinderEnabled},
				urlNumberOfRatingsScaleKey: {ScaleMin},
				urlNumberOfRatingsKey:      {"5"},
				urlSortKey:                 {SortByLastUpdated.String()},
				urlOrderKey:                {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				NumberOfRatingsScale: ScaleMax,
				NumberOfRatings:      5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:         {urlSeriesFinderEnabled},
				urlNumberOfRatingsScaleKey: {ScaleMax},
				urlNumberOfRatingsKey:      {"5"},
				urlSortKey:                 {SortByLastUpdated.String()},
				urlOrderKey:                {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				Readers: 5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlReadersScaleKey: {ScaleMin},
				urlReadersKey:      {"5"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ReadersScale: ScaleMin,
				Readers:      5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlReadersScaleKey: {ScaleMin},
				urlReadersKey:      {"5"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				ReadersScale: ScaleMax,
				Readers:      5,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlReadersScaleKey: {ScaleMax},
				urlReadersKey:      {"5"},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				FirstReleaseDate: "04/04/2022",
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:          {urlSeriesFinderEnabled},
				urlFirstReleaseDateScaleKey: {ScaleMin},
				urlFirstReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                  {SortByLastUpdated.String()},
				urlOrderKey:                 {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				FirstReleaseDateScale: ScaleMin,
				FirstReleaseDate:      "04/04/2022",
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:          {urlSeriesFinderEnabled},
				urlFirstReleaseDateScaleKey: {ScaleMin},
				urlFirstReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                  {SortByLastUpdated.String()},
				urlOrderKey:                 {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				FirstReleaseDateScale: ScaleMax,
				FirstReleaseDate:      "04/04/2022",
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:          {urlSeriesFinderEnabled},
				urlFirstReleaseDateScaleKey: {ScaleMax},
				urlFirstReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                  {SortByLastUpdated.String()},
				urlOrderKey:                 {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				LastReleaseDate: "04/04/2022",
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:         {urlSeriesFinderEnabled},
				urlLastReleaseDateScaleKey: {ScaleMin},
				urlLastReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                 {SortByLastUpdated.String()},
				urlOrderKey:                {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				LastReleaseDateScale: ScaleMin,
				LastReleaseDate:      "04/04/2022",
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:         {urlSeriesFinderEnabled},
				urlLastReleaseDateScaleKey: {ScaleMin},
				urlLastReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                 {SortByLastUpdated.String()},
				urlOrderKey:                {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				LastReleaseDateScale: ScaleMax,
				LastReleaseDate:      "04/04/2022",
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:         {urlSeriesFinderEnabled},
				urlLastReleaseDateScaleKey: {ScaleMax},
				urlLastReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                 {SortByLastUpdated.String()},
				urlOrderKey:                {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreInclude: GenreList{
					GenreAction,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAction)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreInclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorAnd,
				GenreInclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorOr,
				GenreInclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreIncludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreExclude: GenreList{
					GenreAction,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {string(GenreAction)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreExclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorAnd,
				GenreExclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorOr,
				GenreExclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreInclude: GenreList{
					GenreFantasy,
				},
				GenreExclude: GenreList{
					GenreAction,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreFantasy)},
				urlGenreExcludeKey:  {string(GenreAction)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreInclude: GenreList{
					GenreAdult,
					GenreDrama,
					GenreHarem,
					GenreHistorical,
				},
				GenreExclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAdult), string(GenreDrama), string(GenreHarem), string(GenreHistorical)},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorAnd,
				GenreInclude: GenreList{
					GenreAdult,
					GenreDrama,
					GenreHarem,
					GenreHistorical,
				},
				GenreExclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAdult), string(GenreDrama), string(GenreHarem), string(GenreHistorical)},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorOr,
				GenreInclude: GenreList{
					GenreAdult,
					GenreDrama,
					GenreHarem,
					GenreHistorical,
				},
				GenreExclude: GenreList{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey:  {urlSeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreIncludeKey:  {string(GenreAdult), string(GenreDrama), string(GenreHarem), string(GenreHistorical)},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated.String()},
				urlOrderKey:         {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagInclude: TagList{
					TagBlacksmith,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagBlacksmith.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagInclude: TagList{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagArtificialIntelligence.String(), TagBetrayal.String(), TagAliens.String(), TagApocalypse.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorOr,
				TagInclude: TagList{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagArtificialIntelligence.String(), TagBetrayal.String(), TagAliens.String(), TagApocalypse.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorAnd,
				TagInclude: TagList{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorAnd},
				urlTagIncludeKey:   {TagArtificialIntelligence.String(), TagBetrayal.String(), TagAliens.String(), TagApocalypse.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagExclude: TagList{
					TagBlacksmith,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagExcludeKey:   {TagBlacksmith.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagExclude: TagList{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagExcludeKey:   {TagArtificialIntelligence.String(), TagBetrayal.String(), TagAliens.String(), TagApocalypse.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorOr,
				TagExclude: TagList{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagExcludeKey:   {TagArtificialIntelligence.String(), TagBetrayal.String(), TagAliens.String(), TagApocalypse.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorAnd,
				TagExclude: TagList{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorAnd},
				urlTagExcludeKey:   {TagArtificialIntelligence.String(), TagBetrayal.String(), TagAliens.String(), TagApocalypse.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagInclude: TagList{
					TagBlacksmith,
				},
				TagExclude: TagList{
					TagArtists,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagBlacksmith.String()},
				urlTagExcludeKey:   {TagArtists.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagInclude: TagList{
					TagBlacksmith,
					TagArtificialIntelligence,
					TagAbsentParents,
					TagAcademy,
				},
				TagExclude: TagList{
					TagArtists,
					TagDestiny,
					TagColdProtagonist,
					TagDemonicCultivationTechnique,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagBlacksmith.String(), TagArtificialIntelligence.String(), TagAbsentParents.String(), TagAcademy.String()},
				urlTagExcludeKey:   {TagArtists.String(), TagDestiny.String(), TagColdProtagonist.String(), TagDemonicCultivationTechnique.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorOr,
				TagInclude: TagList{
					TagBlacksmith,
					TagArtificialIntelligence,
					TagAbsentParents,
					TagAcademy,
				},
				TagExclude: TagList{
					TagArtists,
					TagDestiny,
					TagColdProtagonist,
					TagDemonicCultivationTechnique,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagBlacksmith.String(), TagArtificialIntelligence.String(), TagAbsentParents.String(), TagAcademy.String()},
				urlTagExcludeKey:   {TagArtists.String(), TagDestiny.String(), TagColdProtagonist.String(), TagDemonicCultivationTechnique.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorAnd,
				TagInclude: TagList{
					TagBlacksmith,
					TagArtificialIntelligence,
					TagAbsentParents,
					TagAcademy,
				},
				TagExclude: TagList{
					TagArtists,
					TagDestiny,
					TagColdProtagonist,
					TagDemonicCultivationTechnique,
				},
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorAnd},
				urlTagIncludeKey:   {TagBlacksmith.String(), TagArtificialIntelligence.String(), TagAbsentParents.String(), TagAcademy.String()},
				urlTagExcludeKey:   {TagArtists.String(), TagDestiny.String(), TagColdProtagonist.String(), TagDemonicCultivationTechnique.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				StoryStatus: StoryStatusAll,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlStoryStatusKey:  {StoryStatusAll.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				StoryStatus: StoryStatusCompleted,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlStoryStatusKey:  {StoryStatusCompleted.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				StoryStatus: StoryStatusOngoing,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlStoryStatusKey:  {StoryStatusOngoing.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				StoryStatus: StoryStatusHiatus,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlStoryStatusKey:  {StoryStatusHiatus.String()},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				SortBy: SortByReviews,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlSortKey:         {SortByReviews.String()},
				urlOrderKey:        {OrderByDescending.String()},
			},
		},
		{
			Query: SearchQuery{
				OrderBy: OrderByAscending,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByAscending.String()},
			},
		},
		{
			Query: SearchQuery{
				Page: 1,
			},
			Expected: map[string][]interface{}{
				urlSeriesFinderKey: {urlSeriesFinderEnabled},
				urlSortKey:         {SortByLastUpdated.String()},
				urlOrderKey:        {OrderByDescending.String()},
				urlPageKey:         {"1"},
			},
		},
	}

	for _, n := range data {
		urlString := buildSearchStringFromQuery(&n.Query)
		parsedUrl, _ := url.ParseQuery(urlString)
		t.Logf("QUERY: %s", urlString)

		for k, v := range n.Expected {
			actual := parsedUrl.Get(k)
			actualSplit := strings.Split(actual, ",")

			for _, vv := range v {
				assert.Containsf(t, actualSplit, vv, "for query param: %s", k)
			}
		}
	}
}
