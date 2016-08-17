package nouns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = [][]string{
	[]string{"child", "children"},
	[]string{"echo", "echoes"},
	[]string{"embargo", "embargoes"},
	[]string{"hero", "heroes"},
	[]string{"potato", "potatoes"},
	[]string{"tomato", "tomatoes"},
	[]string{"torpedo", "torpedoes"},
	[]string{"veto", "vetoes"},
	[]string{"auto", "autos"},
	[]string{"kangaroo", "kangaroos"},
	[]string{"kilo", "kilos"},
	[]string{"memo", "memos"},
	[]string{"photo", "photos"},
	[]string{"piano", "pianos"},
	[]string{"pimento", "pimentos"},
	[]string{"pro", "pros"},
	[]string{"solo", "solos"},
	[]string{"soprano", "sopranos"},
	[]string{"studio", "studios"},
	[]string{"tattoo", "tattoos"},
	[]string{"video", "videos"},
	[]string{"zoo", "zoos"},
	[]string{"fish", "fish"},
	[]string{"sheep", "sheep"},
	[]string{"barrack", "barracks"},
	[]string{"foot", "feet"},
	[]string{"tooth", "teeth"},
	[]string{"goose", "geese"},
	[]string{"tooth", "teeth"},
	[]string{"goose", "geese"},
	[]string{"child", "children"},
	[]string{"man", "men"},
	[]string{"woman", "women"},
	[]string{"person", "people"},
	[]string{"mouse", "mice"},
	[]string{"analysis", "analyses"},
	[]string{"axe", "axes"},
	[]string{"basis", "bases"},
	[]string{"crisis", "crises"},
	[]string{"diagnosis", "diagnoses"},
	[]string{"hypothesis", "hypotheses"},
	[]string{"parenthesis", "parentheses"},
	[]string{"synopsis", "synopses"},
	[]string{"thesis", "theses"},
	[]string{"criterion", "criteria"},
	[]string{"phenomenon", "phenomena"},
	[]string{"automaton", "automata"},
	[]string{"category", "categories"},
	[]string{"test", "tests"},
}

func testTransform(t *testing.T, trans, expected int, transform func(string) (string, error)) {
	for _, testCase := range testCases {
		trans, expected := testCase[trans], testCase[expected]

		actual, err := transform(trans)

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	}
}

func TestPluralize(t *testing.T) {
	testTransform(t, 0, 1, func(trans string) (string, error) {
		return Pluralize(trans)
	})
}

func TestSingularize(t *testing.T) {
	testTransform(t, 1, 0, func(trans string) (string, error) {
		return Singularize(trans)
	})
}
