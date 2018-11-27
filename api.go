package pocket2rm

// APIOriginPocket contains the destination
const APIOriginPocket = "pocket"

// APIOriginMercury is the origin host for the Mercury API
const APIOriginMercury = "mercury"

// APIOriginRemarkable is the origin host for the reMarkable API
const APIOriginRemarkable = "remarkable"

// APIOrigin maps each origin host name
var APIOrigin = map[string]string{
	APIOriginPocket : "https://getpocket.com",
	APIOriginMercury : "https://mercury.postlight.com",
	APIOriginRemarkable : "",
}
