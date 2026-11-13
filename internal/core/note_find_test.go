package core

import (
	"testing"

	"github.com/zk-org/zk/internal/util/test/assert"
)

func TestNoteSorterFromString(t *testing.T) {
	cases := []struct {
		input     string
		field     NoteSortField
		ascending bool
	}{
		{"c", NoteSortCreated, false},
		{"c+", NoteSortCreated, true},
		{"created", NoteSortCreated, false},
		{"created-", NoteSortCreated, false},
		{"created+", NoteSortCreated, true},

		{"m", NoteSortModified, false},
		{"modified", NoteSortModified, false},
		{"modified+", NoteSortModified, true},

		{"p", NoteSortPath, true},
		{"path", NoteSortPath, true},
		{"path-", NoteSortPath, false},

		{"t", NoteSortTitle, true},
		{"title", NoteSortTitle, true},
		{"title-", NoteSortTitle, false},

		{"r", NoteSortRandom, true},
		{"random", NoteSortRandom, true},
		{"random-", NoteSortRandom, false},

		{"wc", NoteSortWordCount, true},
		{"word-count", NoteSortWordCount, true},
		{"word-count-", NoteSortWordCount, false},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			actual, err := NoteSorterFromString(tc.input)
			assert.Nil(t, err)
			assert.Equal(t, actual, NoteSorter{Field: tc.field, Ascending: tc.ascending})
		})
	}

	_, err := NoteSorterFromString("foobar")
	assert.Err(t, err, "foobar: unknown sorting term")
}

func TestSortersFromStrings(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected []NoteSorter
	}{
		{"empty", []string{}, []NoteSorter{}},
		{"single", []string{"created"}, []NoteSorter{
			{Field: NoteSortCreated, Ascending: false},
		}},
		{"reverse order override", []string{"c+", "title", "random"}, []NoteSorter{
			{Field: NoteSortRandom, Ascending: true},
			{Field: NoteSortTitle, Ascending: true},
			{Field: NoteSortCreated, Ascending: true},
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NoteSortersFromStrings(tc.input)
			assert.Nil(t, err)
			assert.Equal(t, actual, tc.expected)
		})
	}

	_, err := NoteSortersFromStrings([]string{"c", "foobar"})
	assert.Err(t, err, "foobar: unknown sorting term")
}

func TestMatchStrategyFromString(t *testing.T) {
	cases := []struct {
		input    string
		expected MatchStrategy
	}{
		{"f", MatchStrategyFts},
		{"fts", MatchStrategyFts},
		{"r", MatchStrategyRe},
		{"re", MatchStrategyRe},
		{"grep", MatchStrategyRe},
		{"e", MatchStrategyExact},
		{"exact", MatchStrategyExact},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			actual, err := MatchStrategyFromString(tc.input)
			assert.Nil(t, err)
			assert.Equal(t, actual, tc.expected)
		})
	}

	_, err := MatchStrategyFromString("foobar")
	assert.Err(t, err, "foobar: unknown match strategy\ntry fts (full-text search), re (regular expression) or exact")
}
