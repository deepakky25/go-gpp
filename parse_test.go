package gpp

import (
	"testing"

	"github.com/prebid/go-gpp/constants"
	"github.com/stretchr/testify/assert"
)

type gppTestData struct {
	description string
	gppString   string
	expected    GppContainer
}

func TestParse(t *testing.T) {
	testData := []gppTestData{
		{
			description: "GPP string with EU TCF V2",
			gppString:   "DBABMA~CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA",
			expected: GppContainer{
				Version:        1,
				SectionStrings: []string{"DBABMA", "CPXxRfAPXxRfAAfKABENB-CgAAAAAAAAAAYgAAAAAAAA"},
				Sectiontypes:   []constants.SectionID{3, 2},
			},
		},
	}

	for _, test := range testData {
		result, err := Parse(test.gppString)

		assert.Nil(t, err)
		assert.Equal(t, test.expected, result)
	}
}