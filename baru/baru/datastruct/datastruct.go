package datastruct

type MLikeRequest struct {
	NAME string `json:"klik"`
}

type MLikeResponse struct {
	MESSAGE string `json:"like"`
}

type guidelinesActive struct {
	GuidelinesName string
	GuidelinesDesc string
	GuidelinesType string
	GuidelinesLink string
}
