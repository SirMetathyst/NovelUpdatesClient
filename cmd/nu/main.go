package main

import (
	"encoding/json"
	"fmt"
	"github.com/SirMetathyst/NovelUpdatesClient"
	"github.com/sahilm/fuzzy"
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

func searchGenre(genreSourceName []string, genreSourceID []string) (genreList NovelUpdatesClient.GenreList) {

	var data []string
	for k, _ := range NovelUpdatesClient.SlugToGenre {
		data = append(data, k)
	}
	for _, gn := range genreSourceName {
		pattern := strings.Trim(gn, " \n\t")
		matches := fuzzy.Find(pattern, data)
		if len(matches) > 0 {
			if genre, ok := NovelUpdatesClient.SlugToGenre[matches[0].Str]; ok {
				genreList = append(genreList, genre)
			}
		}
	}
	for _, gid := range genreSourceID {
		if _, ok := NovelUpdatesClient.GenreToName[NovelUpdatesClient.Genre(gid)]; ok {
			genreList = append(genreList, NovelUpdatesClient.Genre(gid))
		}
	}
	return
}

func main() {

	//const bbold = "\033[1m%s\033[0m"
	//
	//var data []string
	//for k, _ := range NovelUpdatesClient.NameToTag {
	//	data = append(data, k)
	//}
	//
	//reader := bufio.NewReader(os.Stdin)
	//text := ""
	//
	//for text != "exit" {
	//	fmt.Print("Enter Genre: ")
	//	text, _ = reader.ReadString('\n')
	//
	//	//pattern := "magic"
	//	pattern := strings.Trim(text, " \n\t")
	//	matches := fuzzy.Find(pattern, data)
	//
	//	if len(matches) > 0 {
	//		fmt.Println("Found: ", matches[0].Str)
	//	}
	//
	//	//
	//	//for _, match := range matches {
	//	//	for i := 0; i < len(match.Str); i++ {
	//	//		if contains(i, match.MatchedIndexes) {
	//	//			fmt.Print(fmt.Sprintf(bbold, string(match.Str[i])))
	//	//		} else {
	//	//			fmt.Print(string(match.Str[i]))
	//	//		}
	//	//	}
	//	//	fmt.Println()
	//	//}
	//}
	//
	//return

	app := &cli.App{
		Name:    "nu",
		Version: "v0.1",
		Usage:   "NovelUpdate Client",
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			{
				Name:                   "search",
				UseShortOptionHandling: true,
				Flags: []cli.Flag{

					// QUERY
					&cli.StringSliceFlag{Name: "novel-type", Aliases: []string{"nt"}},
					&cli.StringSliceFlag{Name: "language", Aliases: []string{"l"}},
					&cli.StringSliceFlag{Name: "genre-include-name", Aliases: []string{"gin"}},
					&cli.StringSliceFlag{Name: "genre-exclude-name", Aliases: []string{"gen"}},
					&cli.StringSliceFlag{Name: "genre-include-id", Aliases: []string{"gii"}},
					&cli.StringSliceFlag{Name: "genre-exclude-id", Aliases: []string{"gei"}},
					&cli.StringSliceFlag{Name: "tag-include-name", Aliases: []string{"tin"}},
					&cli.StringSliceFlag{Name: "tag-exclude-name", Aliases: []string{"ten"}},
					&cli.StringSliceFlag{Name: "tag-include-id", Aliases: []string{"tii"}},
					&cli.StringSliceFlag{Name: "tag-exclude-id", Aliases: []string{"tei"}},

					// PAGINATION
					&cli.IntFlag{Name: "page", Aliases: []string{"p"}},
					//&cli.IntFlag{Name: "page-from"},
					//&cli.IntFlag{Name: "page-to"},
					//&cli.IntFlag{Name: "items-start"},
					//&cli.IntFlag{Name: "items-end"},

					// JSON
					&cli.BoolFlag{Name: "output-json"},
					&cli.BoolFlag{Name: "output-json-indent"},
				},
				Action: func(context *cli.Context) error {

					/////////////// Collect flags
					novelTypeList, err := NovelUpdatesClient.SlugListToNovelTypeList(context.StringSlice("novel-type"))
					if err != nil {
						return err
					}
					languageList, err := NovelUpdatesClient.SlugListToLanguageList(context.StringSlice("language"))
					if err != nil {
						return err
					}

					genreIncludeName := context.StringSlice("genre-include-name")
					genreIncludeID := context.StringSlice("genre-include-id")
					genreExcludeName := context.StringSlice("genre-exclude-name")
					genreExludeID := context.StringSlice("genre-exclude-id")

					/////////////// Do request
					response, err := NovelUpdatesClient.DoSearchRequestWith(&NovelUpdatesClient.BasicRequester{}, &NovelUpdatesClient.SearchQuery{
						NovelType:    novelTypeList,
						Language:     languageList,
						GenreInclude: searchGenre(genreIncludeName, genreIncludeID),
						GenreExclude: searchGenre(genreExcludeName, genreExludeID),
						Page:         context.Int("page"),
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
