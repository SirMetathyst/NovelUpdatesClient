package NovelUpdatesClient

type StoryStatus int

const (
	StoryStatusDefault   StoryStatus = 1
	StoryStatusAll       StoryStatus = 1
	StoryStatusCompleted StoryStatus = 2
	StoryStatusOngoing   StoryStatus = 3
	StoryStatusHiatus    StoryStatus = 4
)

type Sort string

const (
	SortByDefault     Sort = "sdate"
	SortByChapters    Sort = "srel"
	SortByFrequency   Sort = "sfrel"
	SortByRank        Sort = "srank"
	SortByRating      Sort = "srating"
	SortByReaders     Sort = "sread"
	SortByReviews     Sort = "sreview"
	SortByTitle       Sort = "abc"
	SortByLastUpdated Sort = "sdate"
)

type Order string

const (
	OrderDefault    Order = "desc"
	OrderAscending  Order = "asc"
	OrderDescending Order = "desc"
)

type ScaleType string

const (
	ScaleDefault ScaleType = "min"
	ScaleMin     ScaleType = "min"
	ScaleMax     ScaleType = "max"
)

type Operator string

const (
	OperatorDefault Operator = "and"
	OperatorAnd     Operator = "and"
	OperatorOr      Operator = "or"
)

type NovelTypeResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type LanguageResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GenreResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TagResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SearchQuery struct {
	NovelType              []NovelType `json:"novel_type"`
	Language               []Language  `json:"language"`
	ChaptersMinMax         ScaleType   `json:"chapters_min_max"`
	Chapters               int         `json:"chapters"`
	ReleaseFrequencyMinMax ScaleType   `json:"release_frequency_min_max"`
	ReleaseFrequency       int         `json:"release_frequency"`
	ReviewsMinMax          ScaleType   `json:"reviews_min_max"`
	Reviews                int         `json:"reviews"`
	RatingMinMax           ScaleType   `json:"rating_min_max"`
	Rating                 int         `json:"rating"`
	NumberOfRatingsMinMax  ScaleType   `json:"number_of_ratings_min_max"`
	NumberOfRatings        int         `json:"number_of_ratings"`
	ReadersMinMax          ScaleType   `json:"readers_min_max"`
	Readers                int         `json:"readers"`
	FirstReleaseDateMinMax ScaleType   `json:"first_release_date_min_max"`
	FirstReleaseDate       string      `json:"first_release_date"`
	LastReleaseDateMinMax  ScaleType   `json:"last_release_date_min_max"`
	LastReleaseDate        string      `json:"last_release_date"`
	GenreOperator          Operator    `json:"genre_operator"`
	GenreInclude           []Genre     `json:"genre_include"`
	GenreExclude           []Genre     `json:"genre_exclude"`
	TagOperator            Operator    `json:"tag_operator"`
	TagInclude             []Tag       `json:"tag_include"`
	TagExclude             []Tag       `json:"tag_exclude"`
	StoryStatus            StoryStatus `json:"story_status"`
	SortBy                 Sort        `json:"sort_by"`
	OrderBy                Order       `json:"order_by"`
}

type SearchResult struct {
	Title string `json:"title"`
}
