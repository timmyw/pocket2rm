package pocket2rm

// PocketCode contains the code returned from the first Pocket auth
// call.
type PocketCode struct {
	Code string `json:"code"`
}

// Authorise carries out an OAUTH call to Pocket to get a token
func Authorise(consumerKey string) (string, error) {
	// First step is to get a token from the Pocket servers
	url := "https://getpocket.com/v3/oauth/request"
	params := map[string]interface{}{
		"consumer_key": consumerKey,
		"redirect_uri": "http://localhost",
	}

	result := &PocketCode{}
	err := PostJSON(url, params, result)

	if err != nil {
		return "", err
	}

	return result.Code, nil
	
}
