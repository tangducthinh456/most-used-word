package model

type TopUsedReturn struct {
	Title  string        `json:"title"`
	Result []ItemTopWord `json:"result"` // list item contain word and number it occur
}

type ItemTopWord struct {
	Word        string `json:"word"`
	NumberOccur int    `json:"number_occur"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
