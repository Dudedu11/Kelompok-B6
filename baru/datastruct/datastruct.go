package datastruct

type MLikeRequest struct {
	NAME string `json:"name"`
}

type MLikeResponse struct {
	MESSAGE string `json:"message"`
}

type guidelinesActive struct {
	GuidelinesName string
	GuidelinesDesc string
	GuidelinesType string
	GuidelinesLink string
}
