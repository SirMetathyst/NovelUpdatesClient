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
				GenreInclude: []string{
					GenreAction,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {GenreAction},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreInclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorAnd,
				GenreInclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorOr,
				GenreInclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreIncludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreExclude: []string{
					GenreAction,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {GenreAction},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreExclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorAnd,
				GenreExclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreExcludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorOr,
				GenreExclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreExcludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreInclude: []string{
					GenreFantasy,
				},
				GenreExclude: []string{
					GenreAction,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {GenreFantasy},
				urlGenreExcludeKey:  {GenreAction},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreInclude: []string{
					GenreAdult,
					GenreDrama,
					GenreHarem,
					GenreHistorical,
				},
				GenreExclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {GenreAdult, GenreDrama, GenreHarem, GenreHistorical},
				urlGenreExcludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorAnd,
				GenreInclude: []string{
					GenreAdult,
					GenreDrama,
					GenreHarem,
					GenreHistorical,
				},
				GenreExclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorAnd},
				urlGenreIncludeKey:  {GenreAdult, GenreDrama, GenreHarem, GenreHistorical},
				urlGenreExcludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
				urlSortKey:          {SortByLastUpdated},
				urlOrderKey:         {OrderDescending},
			},
		},
		{
			Query: SearchQuery{
				GenreOperator: OperatorOr,
				GenreInclude: []string{
					GenreAdult,
					GenreDrama,
					GenreHarem,
					GenreHistorical,
				},
				GenreExclude: []string{
					GenreAction,
					GenreAdventure,
					GenreComedy,
					GenreFantasy,
				},
			},
			Expected: map[string][]string{
				urlSeriesFinderKey:  {SeriesFinderEnabled},
				urlGenreOperatorKey: {OperatorOr},
				urlGenreIncludeKey:  {GenreAdult, GenreDrama, GenreHarem, GenreHistorical},
				urlGenreExcludeKey:  {GenreAction, GenreAdventure, GenreComedy, GenreFantasy},
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
