package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const userId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initalLink_1 := "https://www.linkedin.com/in/siddxharth.html"
	shortLink_1 := GenerateShortLink(initalLink_1, userId)

	initalLink_2 := "https://www.github.com/siddxharth"
	shortLink_2 := GenerateShortLink(initalLink_2, userId)

	initalLink_3 := "mailto:siddxharth@gmail.com"
	shortLink_3 := GenerateShortLink(initalLink_3, userId)

	assert.Equal(t, shortLink_1, "linkedin")
	assert.Equal(t, shortLink_2, "github")
	assert.Equal(t, shortLink_3, "email")
}
