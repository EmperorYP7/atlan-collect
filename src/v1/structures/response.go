package structures

type GetFormResponse struct {
	Status       string         `json:"status"`
	Message      string         `json:"message"`
	FormData     FormData       `json:"form_data"`
}

type GetQuestionResponse struct {
	Status         string         `json:"status"`
	Message        string         `json:"message"`
	QuestionData   QuestionData   `json:"question_data"`
}

type GetResponseResponse struct {
	Status              string             `json:"status"`
	Message             string             `json:"message"`
	ResponseBankData    []ResponseBankData `json:"response_data"`
}

type GetResponseBankResponse struct {
	Status              string             `json:"status"`
	Message             string             `json:"message"`
	ResponseBankData    ResponseBankData   `json:"response_bank_data"`
}
