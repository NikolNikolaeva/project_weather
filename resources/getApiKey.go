package resources

import (
	"encoding/json"
	"io"
	"os"
)

type Cred struct {
	ApiKey string `json:"apiKey"`
}

func GetApiKey(credFile string) (string, error) {
	jsonFile, err := os.OpenFile(credFile, os.O_RDONLY, os.ModePerm)

	if err != nil {
		return "", err
	}
	defer func() { _ = jsonFile.Close() }()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}

	var cred Cred
	err = json.Unmarshal(byteValue, &cred)

	if err != nil {
		return "", err
	}

	return cred.ApiKey, nil
}
