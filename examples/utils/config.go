package utils

import (
	"encoding/json"
	"io/ioutil"
)

// APIConfig contains all information needed to interact with the Macstadium API
// and make non-administrative requests.
type APIConfig struct {
	// ServerURL is the Orka API url. This is usually a string like "http://10.221.188.100"
	ServerURL string `json:"server_url"`

	// BearerToken is the authentication token corresponding to the user that will be
	// performing these API operations.
	BearerToken string `json:"bearer_token"`
}

func ParseAPIConfig(configPath string) (APIConfig, error) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return APIConfig{}, err
	}

	config := APIConfig{}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return APIConfig{}, err
	}

	return config, nil
}

// UserConfig contains all information needed to create a new user and generate a token.
type UserConfig struct {
	ServerURL string `json:"server_url"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func ParseUserConfig(path string) (UserConfig, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return UserConfig{}, err
	}

	config := UserConfig{}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return UserConfig{}, err
	}

	return config, nil
}
