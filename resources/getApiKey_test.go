package resources

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApiKey_Success(t *testing.T) {
	cred := `{"apiKey": "test_key"}`

	tmpFile, err := os.CreateTemp("", "test_cred.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(cred)
	assert.NoError(t, err)
	tmpFile.Close()

	apiKey, err := GetApiKey(tmpFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "test_key", apiKey)
}

func TestGetApiKey_FileNotFound(t *testing.T) {
	apiKey, err := GetApiKey("non_existent_file.json")
	assert.Error(t, err)
	assert.Equal(t, "", apiKey)
}

func TestGetApiKey_InvalidJSON(t *testing.T) {
	incredFile, err := os.CreateTemp("", "invalid_test_cred_*.json")
	assert.NoError(t, err)
	defer os.Remove(incredFile.Name())

	_, err = incredFile.WriteString(`{"apiKey": `) // Invalid JSON
	assert.NoError(t, err)
	incredFile.Close()

	apiKey, err := GetApiKey(incredFile.Name())
	assert.Error(t, err)
	assert.Equal(t, "", apiKey)
}
