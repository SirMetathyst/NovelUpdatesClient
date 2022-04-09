package main

import (
	"encoding/json"
	"fmt"
	"github.com/SirMetathyst/NovelUpdatesClient"
	"github.com/sahilm/fuzzy"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

func main() {
	app := &cli.App{
		Name:    "nu",
		Version: "v0.1",
		Usage:   "NovelUpdate Client",
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			{
				Name: "search",
				Flags: []cli.Flag{

					// QUERY
					&cli.StringSliceFlag{Name: "novel-type", Aliases: []string{"nt"}},
					&cli.StringSliceFlag{Name: "language", Aliases: []string{"l"}},
					&cli.StringFlag{Name: "chapters-scale", Aliases: []string{"cs"}, Usage: "min|max", DefaultText: "min", Value: "min"},
					&cli.IntFlag{Name: "chapters", Aliases: []string{"c"}},
					&cli.StringFlag{Name: "release-frequency-scale", Aliases: []string{"rfs"}, Usage: "min|max", DefaultText: "max", Value: "max"},
					&cli.IntFlag{Name: "release-frequency", Aliases: []string{"rf"}},
					&cli.StringFlag{Name: "reviews-scale", Aliases: []string{"rs"}, Usage: "min|max", DefaultText: "min", Value: "min"},
					&cli.IntFlag{Name: "reviews", Aliases: []string{"r"}},
					&cli.StringFlag{Name: "rating-scale", Aliases: []string{"rts"}, Usage: "min|max", DefaultText: "min", Value: "min"},
					&cli.IntFlag{Name: "rating", Aliases: []string{"rt"}},
					&cli.StringFlag{Name: "number-of-ratings-scale", Aliases: []string{"nrts"}, Usage: "min|max", DefaultText: "min", Value: "min"},
					&cli.IntFlag{Name: "number-of-ratings", Aliases: []string{"nrt"}},
					&cli.StringFlag{Name: "readers-scale", Aliases: []string{"rds"}, Usage: "min|max", DefaultText: "min", Value: "min"},
					&cli.IntFlag{Name: "readers", Aliases: []string{"rd"}},
					&cli.StringFlag{Name: "first-release-date-scale", Aliases: []string{"frds"}, Usage: "min|max", DefaultText: "min", Value: "min"},
					&cli.StringFlag{Name: "first-release-date", Aliases: []string{"frd"}},
					&cli.StringFlag{Name: "last-release-date-scale", Aliases: []string{"lrds"}, Usage: "min|max", DefaultText: "min", Value: "min"},
					&cli.StringFlag{Name: "last-release-date", Aliases: []string{"lrd"}},
					&cli.StringFlag{Name: "genre-operator", Aliases: []string{"gop"}, Usage: "and|or", DefaultText: "and", Value: "or"},
					&cli.StringSliceFlag{Name: "genre-include-name", Aliases: []string{"gin"}},
					&cli.StringSliceFlag{Name: "genre-exclude-name", Aliases: []string{"gen"}},
					&cli.StringSliceFlag{Name: "genre-include-id", Aliases: []string{"gii"}},
					&cli.StringSliceFlag{Name: "genre-exclude-id", Aliases: []string{"gei"}},
					&cli.StringFlag{Name: "tag-operator", Aliases: []string{"top"}, Usage: "and|or", DefaultText: "or", Value: "or"},
					&cli.StringSliceFlag{Name: "tag-include-name", Aliases: []string{"tin"}},
					&cli.StringSliceFlag{Name: "tag-exclude-name", Aliases: []string{"ten"}},
					&cli.StringSliceFlag{Name: "tag-include-id", Aliases: []string{"tii"}},
					&cli.StringSliceFlag{Name: "tag-exclude-id", Aliases: []string{"tei"}},
					&cli.StringFlag{Name: "story-status", Aliases: []string{"ss"}, Usage: "all|completed|ongoing|hiatus", DefaultText: NovelUpdatesClient.StoryStatusDefault.SlugString(), Value: NovelUpdatesClient.StoryStatusDefault.SlugString()},
					&cli.StringFlag{Name: "sort-by", Aliases: []string{"sb"}, Usage: "chapters|frequency|rank|rating|readers|reviews|title|last-updated", DefaultText: NovelUpdatesClient.SortByDefault.SlugString(), Value: NovelUpdatesClient.SortByDefault.SlugString()},
					&cli.StringFlag{Name: "order-by", Aliases: []string{"ob"}, Usage: "asc|desc", DefaultText: NovelUpdatesClient.OrderByDefault.SlugString(), Value: NovelUpdatesClient.OrderByDefault.SlugString()},

					// PAGINATION
					&cli.IntFlag{Name: "page", Aliases: []string{"p"}, DefaultText: "1"},
					//&cli.IntFlag{Name: "page-from"},
					//&cli.IntFlag{Name: "page-to"},
					//&cli.IntFlag{Name: "items-start"},
					//&cli.IntFlag{Name: "items-end"},

					// EXCLUDE FIELD(S)
					&cli.BoolFlag{Name: "output-exclude-id", Aliases: []string{"oes"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-url", Aliases: []string{"oesu"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-title", Aliases: []string{"oet"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-chapters", Aliases: []string{"oec"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-release-frequency", Aliases: []string{"oerf"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-readers", Aliases: []string{"oerd"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-reviews", Aliases: []string{"oer"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-last-updated", Aliases: []string{"oelu"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-genre", Aliases: []string{"oeg"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-description", Aliases: []string{"oed"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-short-language", Aliases: []string{"oesl"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-exclude-rating", Aliases: []string{"oert"}, Usage: "true|false", DefaultText: "false", Value: false},

					// JSON
					&cli.BoolFlag{Name: "output-json", Aliases: []string{"oj"}, Usage: "true|false", DefaultText: "false", Value: false},
					&cli.BoolFlag{Name: "output-json-indent", Aliases: []string{"oji"}, Usage: "true|false", DefaultText: "false", Value: false},
				},
				Action: func(context *cli.Context) error {

					/////////////// Collect flags
					novelTypeList := context.StringSlice("novel-type")
					languageList := context.StringSlice("language")
					genreIncludeNameList := context.StringSlice("genre-include-name")
					genreIncludeIDList := context.StringSlice("genre-include-id")
					genreExcludeNameList := context.StringSlice("genre-exclude-name")
					genreExludeIDList := context.StringSlice("genre-exclude-id")
					tagIncludeNameList := context.StringSlice("tag-include-name")
					tagIncludeIDList := context.StringSlice("tag-include-id")
					tagExcludeNameList := context.StringSlice("tag-exclude-name")
					tagExludeIDList := context.StringSlice("tag-exclude-id")
					storyStatus := context.String("story-status")
					sortBy := context.String("order-by")
					orderBy := context.String("order-by")
					page := context.Int("page")

					/////////////// Do request
					response, err := NovelUpdatesClient.DoSearchRequestWith(&NovelUpdatesClient.BasicRequester{}, &NovelUpdatesClient.SearchQuery{
						NovelType: searchNovelType(novelTypeList),
						Language:  searchLanguage(languageList),

						GenreInclude: searchGenre(genreIncludeNameList, genreIncludeIDList),
						GenreExclude: searchGenre(genreExcludeNameList, genreExludeIDList),
						TagInclude:   searchTag(tagIncludeNameList, tagIncludeIDList),
						TagExclude:   searchTag(tagExcludeNameList, tagExludeIDList),
						StoryStatus:  NovelUpdatesClient.SlugStringToStoryStatusOrDefault(storyStatus),
						SortBy:       NovelUpdatesClient.SlugStringToSortByOrDefault(sortBy),
						OrderBy:      NovelUpdatesClient.SlugStringToOrderByOrDefault(orderBy),
						Page:         page,
					})

					/////////////// Output response
					if err != nil {
						return err
					}
					if context.Bool("output-json") {
						if !context.Bool("output-json-indent") {
							v, err := json.Marshal(response)
							if err != nil {
								return err
							}
							fmt.Println(string(v))
						} else {
							v, err := json.MarshalIndent(response, "", "\t")
							if err != nil {
								return err
							}
							fmt.Println(string(v))
						}
					}

					///////////////
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func searchNovelType(novelTypeSourceName []string) (novelTypeList NovelUpdatesClient.NovelTypeList) {

	var data []string
	for k, _ := range NovelUpdatesClient.SlugStringToNovelType {
		data = append(data, k)
	}
	for _, gn := range novelTypeSourceName {
		pattern := strings.Trim(gn, " \n\t")
		matches := fuzzy.Find(pattern, data)
		if len(matches) > 0 {
			if novelType, ok := NovelUpdatesClient.SlugStringToNovelType[matches[0].Str]; ok {
				novelTypeList = append(novelTypeList, novelType)
			}
		}
	}
	return
}

func searchLanguage(languageSourceName []string) (languageList NovelUpdatesClient.LanguageList) {

	var data []string
	for k, _ := range NovelUpdatesClient.SlugStringToLanguage {
		data = append(data, k)
	}
	for _, gn := range languageSourceName {
		pattern := strings.Trim(gn, " \n\t")
		matches := fuzzy.Find(pattern, data)
		if len(matches) > 0 {
			if language, ok := NovelUpdatesClient.SlugStringToLanguage[matches[0].Str]; ok {
				languageList = append(languageList, language)
			}
		}
	}
	return
}

func searchGenre(sourceName []string, sourceID []string) (list NovelUpdatesClient.GenreList) {

	var data []string
	for k, _ := range NovelUpdatesClient.SlugStringToGenre {
		data = append(data, k)
	}
	for _, gn := range sourceName {
		pattern := strings.Trim(gn, " \n\t")
		matches := fuzzy.Find(pattern, data)
		if len(matches) > 0 {
			if genre, ok := NovelUpdatesClient.SlugStringToGenre[matches[0].Str]; ok {
				list = append(list, genre)
			}
		}
	}
	for _, gid := range sourceID {
		if _, ok := NovelUpdatesClient.GenreToDisplayString[NovelUpdatesClient.Genre(gid)]; ok {
			list = append(list, NovelUpdatesClient.Genre(gid))
		}
	}
	return
}

func searchTag(sourceName []string, sourceID []string) (list NovelUpdatesClient.TagList) {

	var data []string
	for k, _ := range NovelUpdatesClient.SlugStringToTag {
		data = append(data, k)
	}
	for _, gn := range sourceName {
		pattern := strings.Trim(gn, " \n\t")
		matches := fuzzy.Find(pattern, data)
		if len(matches) > 0 {
			if tag, ok := NovelUpdatesClient.SlugStringToTag[matches[0].Str]; ok {
				list = append(list, tag)
			}
		}
	}
	for _, gid := range sourceID {
		if _, ok := NovelUpdatesClient.TagToDisplayString[NovelUpdatesClient.Tag(gid)]; ok {
			list = append(list, NovelUpdatesClient.Tag(gid))
		}
	}
	return
}
