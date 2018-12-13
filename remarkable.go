package pocket2rm

// RMToken is a reMarkable access token
type RMToken struct {
	Token string `json:"token"`
}

// GetRMToken will retrieve an access token from Pocket
func GetRMToken(code string, deviceDesc string, deviceID string) (*RMToken, error) {
	result := &RMToken{}
	err := PostJSON("/token/json/2/device/new",
		APIOriginRemarkable,
		map[string]string{
			"code":       code,
			"deviceDesc": deviceDesc,
			"deviceID":   deviceID,
		},
		result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
