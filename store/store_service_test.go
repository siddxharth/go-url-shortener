package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreStorage = &StorageService{}

func init() {
	testStoreStorage = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreStorage.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.linkedin.com/in/siddxharth.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	SaveUrlMapping(shortURL, initialLink, userUUId)

	// Retrieve initial URL
	retrievedUrl := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
