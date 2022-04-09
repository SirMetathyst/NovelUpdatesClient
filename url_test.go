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
				NovelType: NovelTypeList{
					NovelTypeLightNovel,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlNovelTypeKey:    {string(NovelTypeLightNovel)},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				NovelType: NovelTypeList{
					NovelTypeLightNovel,
					NovelTypeWebNovel,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlNovelTypeKey:    {string(NovelTypeLightNovel), string(NovelTypeWebNovel)},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlNovelTypeKey:    {string(NovelTypeLightNovel), string(NovelTypeWebNovel), string(NovelTypePublishedNovel)},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Language: LanguageList{
					LanguageChinese,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlLanguageKey:     {string(LanguageChinese)},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Language: LanguageList{
					LanguageChinese,
					LanguageFilipino,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlLanguageKey:     {string(LanguageChinese), string(LanguageFilipino)},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlLanguageKey:     {string(LanguageChinese), string(LanguageFilipino), string(LanguageIndonesian)},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlLanguageKey:     {string(LanguageChinese), string(LanguageFilipino), string(LanguageIndonesian), string(LanguageJapanese)},
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
		{
			Query: SearchQuery{
				Rating: 5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlRatingScaleKey:  {ScaleMin},
				urlRatingKey:       {"5"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				RatingScale: ScaleMin,
				Rating:      5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlRatingScaleKey:  {ScaleMin},
				urlRatingKey:       {"5"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				RatingScale: ScaleMax,
				Rating:      5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlRatingScaleKey:  {ScaleMax},
				urlRatingKey:       {"5"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				NumberOfRatings: 5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:         {SeriesFinderEnabled},
				urlNumberOfRatingsScaleKey: {ScaleMin},
				urlNumberOfRatingsKey:      {"5"},
				urlSortKey:                 {SortByLastUpdated},
				urlOrderKey:                {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				NumberOfRatingsScale: ScaleMin,
				NumberOfRatings:      5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:         {SeriesFinderEnabled},
				urlNumberOfRatingsScaleKey: {ScaleMin},
				urlNumberOfRatingsKey:      {"5"},
				urlSortKey:                 {SortByLastUpdated},
				urlOrderKey:                {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				NumberOfRatingsScale: ScaleMax,
				NumberOfRatings:      5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:         {SeriesFinderEnabled},
				urlNumberOfRatingsScaleKey: {ScaleMax},
				urlNumberOfRatingsKey:      {"5"},
				urlSortKey:                 {SortByLastUpdated},
				urlOrderKey:                {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				Readers: 5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlReadersScaleKey: {ScaleMin},
				urlReadersKey:      {"5"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ReadersScale: ScaleMin,
				Readers:      5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlReadersScaleKey: {ScaleMin},
				urlReadersKey:      {"5"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				ReadersScale: ScaleMax,
				Readers:      5,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlReadersScaleKey: {ScaleMax},
				urlReadersKey:      {"5"},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				FirstReleaseDate: "04/04/2022",
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:          {SeriesFinderEnabled},
				urlFirstReleaseDateScaleKey: {ScaleMin},
				urlFirstReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                  {SortByLastUpdated},
				urlOrderKey:                 {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				FirstReleaseDateScale: ScaleMin,
				FirstReleaseDate:      "04/04/2022",
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:          {SeriesFinderEnabled},
				urlFirstReleaseDateScaleKey: {ScaleMin},
				urlFirstReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                  {SortByLastUpdated},
				urlOrderKey:                 {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				FirstReleaseDateScale: ScaleMax,
				FirstReleaseDate:      "04/04/2022",
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:          {SeriesFinderEnabled},
				urlFirstReleaseDateScaleKey: {ScaleMax},
				urlFirstReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                  {SortByLastUpdated},
				urlOrderKey:                 {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				LastReleaseDate: "04/04/2022",
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:         {SeriesFinderEnabled},
				urlLastReleaseDateScaleKey: {ScaleMin},
				urlLastReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                 {SortByLastUpdated},
				urlOrderKey:                {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				LastReleaseDateScale: ScaleMin,
				LastReleaseDate:      "04/04/2022",
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:         {SeriesFinderEnabled},
				urlLastReleaseDateScaleKey: {ScaleMin},
				urlLastReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                 {SortByLastUpdated},
				urlOrderKey:                {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				LastReleaseDateScale: ScaleMax,
				LastReleaseDate:      "04/04/2022",
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:         {SeriesFinderEnabled},
				urlLastReleaseDateScaleKey: {ScaleMax},
				urlLastReleaseDateKey:      {"04/04/2022"},
				urlSortKey:                 {SortByLastUpdated},
				urlOrderKey:                {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreInclude: GenreList{
					GenreAction,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAction)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreIncludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreExclude: GenreList{
					GenreAction,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {string(GenreAction)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreFantasy)},
				urlGenreExcludeKey:  {string(GenreAction)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAdult), string(GenreDrama), string(GenreHarem), string(GenreHistorical)},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {string(GenreAdult), string(GenreDrama), string(GenreHarem), string(GenreHistorical)},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
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
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreIncludeKey:  {string(GenreAdult), string(GenreDrama), string(GenreHarem), string(GenreHistorical)},
				urlGenreExcludeKey:  {string(GenreAction), string(GenreAdventure), string(GenreComedy), string(GenreFantasy)},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagInclude: []string{
					TagBlacksmith,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagBlacksmith},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagInclude: []string{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagArtificialIntelligence, TagBetrayal, TagAliens, TagApocalypse},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorOr,
				TagInclude: []string{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagArtificialIntelligence, TagBetrayal, TagAliens, TagApocalypse},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorAnd,
				TagInclude: []string{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorAnd},
				urlTagIncludeKey:   {TagArtificialIntelligence, TagBetrayal, TagAliens, TagApocalypse},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		///////////////////////////////////////////////
		{
			Query: SearchQuery{
				TagExclude: []string{
					TagBlacksmith,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagExcludeKey:   {TagBlacksmith},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagExclude: []string{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagExcludeKey:   {TagArtificialIntelligence, TagBetrayal, TagAliens, TagApocalypse},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorOr,
				TagExclude: []string{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagExcludeKey:   {TagArtificialIntelligence, TagBetrayal, TagAliens, TagApocalypse},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorAnd,
				TagExclude: []string{
					TagArtificialIntelligence,
					TagBetrayal,
					TagAliens,
					TagApocalypse,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorAnd},
				urlTagExcludeKey:   {TagArtificialIntelligence, TagBetrayal, TagAliens, TagApocalypse},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagInclude: []string{
					TagBlacksmith,
				},
				TagExclude: []string{
					TagArtists,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagBlacksmith},
				urlTagExcludeKey:   {TagArtists},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagInclude: []string{
					TagBlacksmith,
					TagArtificialIntelligence,
					TagAbsentParents,
					TagAcademy,
				},
				TagExclude: []string{
					TagArtists,
					TagDestiny,
					TagColdProtagonist,
					TagDemonicCultivationTechnique,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagBlacksmith, TagArtificialIntelligence, TagAbsentParents, TagAcademy},
				urlTagExcludeKey:   {TagArtists, TagDestiny, TagColdProtagonist, TagDemonicCultivationTechnique},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorOr,
				TagInclude: []string{
					TagBlacksmith,
					TagArtificialIntelligence,
					TagAbsentParents,
					TagAcademy,
				},
				TagExclude: []string{
					TagArtists,
					TagDestiny,
					TagColdProtagonist,
					TagDemonicCultivationTechnique,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorOr},
				urlTagIncludeKey:   {TagBlacksmith, TagArtificialIntelligence, TagAbsentParents, TagAcademy},
				urlTagExcludeKey:   {TagArtists, TagDestiny, TagColdProtagonist, TagDemonicCultivationTechnique},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				TagOperator: OperatorAnd,
				TagInclude: []string{
					TagBlacksmith,
					TagArtificialIntelligence,
					TagAbsentParents,
					TagAcademy,
				},
				TagExclude: []string{
					TagArtists,
					TagDestiny,
					TagColdProtagonist,
					TagDemonicCultivationTechnique,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlTagOperatorKey:  {OperatorAnd},
				urlTagIncludeKey:   {TagBlacksmith, TagArtificialIntelligence, TagAbsentParents, TagAcademy},
				urlTagExcludeKey:   {TagArtists, TagDestiny, TagColdProtagonist, TagDemonicCultivationTechnique},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				StoryStatus: StoryStatusAll,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlStoryStatusKey:  {StoryStatusAll},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				StoryStatus: StoryStatusCompleted,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlStoryStatusKey:  {StoryStatusCompleted},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				StoryStatus: StoryStatusOngoing,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlStoryStatusKey:  {StoryStatusOngoing},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				StoryStatus: StoryStatusHiatus,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlStoryStatusKey:  {StoryStatusHiatus},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				SortBy: SortByReviews,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlSortKey:         {SortByReviews},
				urlOrderKey:        {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				OrderBy: OrderAscending,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderAscending},
			},
		},
		{
			Query: SearchQuery{
				Page: 1,
			},
			Expected: map[string][]string{
				urlSeriesFinderKey: {SeriesFinderEnabled},
				urlSortKey:         {SortByLastUpdated},
				urlOrderKey:        {OrderDescending},
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
