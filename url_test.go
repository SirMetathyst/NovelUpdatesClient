package NovelUpdatesClient_test

import (
	"github.com/SirMetathyst/NovelUpdatesClient"
	"testing"
)

func TestBuildSearchStringFromQuery_NovelType_Empty(t *testing.T) {

	expected := "?sf=1&sort=sdate&order=desc"
	if str := NovelUpdatesClient.BuildSearchStringFromQuery(nil); str != expected {
		t.Errorf("expected `%s`, got `%s`", expected, str)
	}
}

func TestBuildSearchStringFromQuery_NovelType_One(t *testing.T) {

	q := &NovelUpdatesClient.SearchQuery{
		NovelType: []NovelUpdatesClient.NovelType{
			NovelUpdatesClient.NovelTypeWebNovel,
		},
	}

	expected := "?sf=1&nt=2444&sort=sdate&order=desc"
	if str := NovelUpdatesClient.BuildSearchStringFromQuery(q); str != expected {
		t.Errorf("expected `%s`, got `%s`", expected, str)
	}
}

func TestBuildSearchStringFromQuery_NovelType_Two(t *testing.T) {

	q := &NovelUpdatesClient.SearchQuery{
		NovelType: []NovelUpdatesClient.NovelType{
			NovelUpdatesClient.NovelTypeWebNovel,
			NovelUpdatesClient.NovelTypePublishedNovel,
		},
	}

	expected := "?sf=1&nt=2444,26874&sort=sdate&order=desc"
	if str := NovelUpdatesClient.BuildSearchStringFromQuery(q); str != expected {
		t.Errorf("expected `%s`, got `%s`", expected, str)
	}
}

func TestBuildSearchStringFromQuery_NovelType_Three(t *testing.T) {

	q := &NovelUpdatesClient.SearchQuery{
		NovelType: []NovelUpdatesClient.NovelType{
			NovelUpdatesClient.NovelTypeWebNovel,
			NovelUpdatesClient.NovelTypePublishedNovel,
			NovelUpdatesClient.NovelTypePublishedNovel,
		},
	}

	expected := "?sf=1&nt=2444,26874,26874&sort=sdate&order=desc"
	if str := NovelUpdatesClient.BuildSearchStringFromQuery(q); str != expected {
		t.Errorf("expected `%s`, got `%s`", expected, str)
	}
}
