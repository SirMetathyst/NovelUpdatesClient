package NovelUpdatesClient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoSearchRequest_DoesNotReturn_Nil(t *testing.T) {

	resp, err := DoSearchRequest(nil)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestDoSearchRequest_Returns_ResultsCount_As_25(t *testing.T) {

	resp, _ := DoSearchRequest(nil)

	assert.Len(t, resp, 25)
}

func TestDoSearchRequest_Returns_Title(t *testing.T) {

	resp, _ := DoSearchRequest(&SearchQuery{OrderBy: OrderAscending, SortBy: SortByLastUpdated})

	expectedTitles := []string{
		".hack//AI Buster",
		".hack//Cell",
		".hack//G.U.",
		"“Bokuchu” Us vs the Police: 700-Day War",
	}

	for i, expectedTitle := range expectedTitles {
		if resp[i].Title != expectedTitle {
			t.Errorf("expected title: `%s`, got `%s`", expectedTitle, resp[i].Title)
		}
	}
}

func TestDoSearchRequest_Returns_ChaptersCount(t *testing.T) {

	resp, _ := DoSearchRequest(&SearchQuery{OrderBy: OrderAscending, SortBy: SortByLastUpdated, Chapters: 3})

	expectedTitles := []struct {
		Title         string
		ChaptersCount int
	}{
		{Title: "Reverend Insanity", ChaptersCount: 25},
		{Title: "Kidou Senshi Z Gundam", ChaptersCount: 3},
		{Title: "Demon Child", ChaptersCount: 13},
		{Title: "Saiunkoku Monogatari: Kouryou no Yume", ChaptersCount: 10},
	}

	for i, expectedTitle := range expectedTitles {
		if resp[i].Title != expectedTitle.Title {
			t.Errorf("expected title: `%s`, got `%s`", expectedTitle.Title, resp[i].Title)
		}
		if resp[i].Chapters != expectedTitle.ChaptersCount {
			t.Errorf("expected chapter count: `%d`, got `%d`", expectedTitle.ChaptersCount, resp[i].Chapters)
		}
	}
}

func TestDoSearchRequest_Returns_UpdateFrequency(t *testing.T) {

	resp, _ := DoSearchRequest(&SearchQuery{OrderBy: OrderAscending, SortBy: SortByLastUpdated, ReleaseFrequency: 300})

	expectedTitles := []struct {
		Title            string
		ReleaseFrequency float64
	}{
		{Title: ".hack//AI Buster", ReleaseFrequency: 2000},
		{Title: ".hack//Cell", ReleaseFrequency: 8030},
		{Title: ".hack//G.U.", ReleaseFrequency: 8030},
		{Title: `“Bokuchu” Us vs the Police: 700-Day War`, ReleaseFrequency: 8030},
	}

	for i, expectedTitle := range expectedTitles {
		if resp[i].Title != expectedTitle.Title {
			t.Errorf("expected title: `%s`, got `%s`", expectedTitle.Title, resp[i].Title)
		}
		if resp[i].ReleaseFrequencyInDays != expectedTitle.ReleaseFrequency {
			t.Errorf("expected release frequency: `%f`, got `%f`", expectedTitle.ReleaseFrequency, resp[i].ReleaseFrequencyInDays)
		}
	}
}

func TestDoSearchRequest_Returns_Readers(t *testing.T) {

	resp, _ := DoSearchRequest(&SearchQuery{OrderBy: OrderAscending, SortBy: SortByLastUpdated, ReleaseFrequency: 300})

	expectedTitles := []struct {
		Title   string
		Readers int
	}{
		{Title: ".hack//AI Buster", Readers: 145},
		{Title: ".hack//Cell", Readers: 96},
		{Title: ".hack//G.U.", Readers: 102},
		{Title: `“Bokuchu” Us vs the Police: 700-Day War`, Readers: 56},
	}

	for i, expectedTitle := range expectedTitles {
		if resp[i].Title != expectedTitle.Title {
			t.Errorf("expected title: `%s`, got `%s`", expectedTitle.Title, resp[i].Title)
		}
		if resp[i].Readers != expectedTitle.Readers {
			t.Errorf("expected readers: `%d`, got `%d`", expectedTitle.Readers, resp[i].Readers)
		}
	}
}

func TestDoSearchRequest_Returns_Reviews(t *testing.T) {

	resp, _ := DoSearchRequest(&SearchQuery{OrderBy: OrderAscending, SortBy: SortByLastUpdated, ReleaseFrequency: 300, Reviews: 3})

	expectedTitles := []struct {
		Title   string
		Reviews int
	}{
		{Title: "Ai No Kusabi", Reviews: 8},
		{Title: "Back to Twelve", Reviews: 3},
		{Title: "Cold Fever", Reviews: 3},
		{Title: `Dungeon ni Deai o Motomeru no wa Machigatte Iru Darou ka`, Reviews: 5},
	}

	for i, expectedTitle := range expectedTitles {
		if resp[i].Title != expectedTitle.Title {
			t.Errorf("expected title: `%s`, got `%s`", expectedTitle.Title, resp[i].Title)
		}
		if resp[i].Reviews != expectedTitle.Reviews {
			t.Errorf("expected reviews: `%d`, got `%d`", expectedTitle.Reviews, resp[i].Reviews)
		}
	}
}
