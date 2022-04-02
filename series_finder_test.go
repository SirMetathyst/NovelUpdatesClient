package NovelUpdatesClient_test

import (
	"NovelUpdatesClient"
	"testing"
)

func TestSeriesFinder_DoesNotReturn_Nil(t *testing.T) {

	resp, err := NovelUpdatesClient.SeriesFinder()

	if err != nil {
		t.Errorf("expected `nil` error, got `%v`", err)
	}

	if resp == nil {
		t.Error("expected non-nil response, got `nil`")
	}
}

func TestSeriesFinder_Returns_ResultsCount_As_25(t *testing.T) {

	resp, _ := NovelUpdatesClient.SeriesFinder()

	if len(resp) != 25 {
		t.Errorf("got `%d`, expected `25`", len(resp))
	}
}

func TestSeriesFinder_Returns_Title(t *testing.T) {

	resp, _ := NovelUpdatesClient.SeriesFinder()

	expectedTitles := []string{
		".hack//AI Buster",
		".hack//Cell",
		".hack//G.U.",
		"“Bokuchu” Us vs the Police: 700-Day War",
		"37°C",
		"Aa Megami-sama",
		"Adventurer Lindy’s Calamity",
		"Ai No Kusabi",
		"Ai Shika Iranee yo",
		"Aibiki",
		"Akihabara Pioneer Road",
		"Akikan!",
		"Akuma to Yobanaide",
		"All You Need Is Kill",
		"Ambang Batas",
		"Angel Beats! – Track Zero",
		"Another",
		"Baccano!",
		"Back to Twelve",
		"Baito de Wizard",
		"Bakemono No Ko",
		"Battle Royale",
		"Black Blood Brothers",
		"Black Bullet",
		"Black Lagoon",
	}

	for i, expectedTitle := range expectedTitles {
		if resp[i].Title != expectedTitle {
			t.Errorf("expected title: `%s`, got `%s`", expectedTitle, resp[i].Title)
		}
	}
}
