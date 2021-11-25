package structures

import "time"

type MetaData struct {
}

type QuestionData struct {
	QuestionBankId   int64      `json:"question_bank_id"`
	FormId           int64      `json:"form_id"`
	UserId           int64      `json:"user_id"`
	IsPublished      bool       `json:"is_published"`
	QuestionText     string     `json:"question_text"`
	Type             string     `json:"type"`
	MetaData         MetaData   `json:"metadata"`
}

type FormData struct {
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	CreatedAt        time.Time       `json:"createdAt"`
	IsPublished      bool            `json:"is_published"`
	FormId           int64           `json:"form_id"`
	UserId           int64           `json:"user_id"`
	Version          string          `json:"version"`
	QuestionBank     []QuestionData  `json:"question_bank"`
}

type ResponseData struct {
	ResponseId           int64           `json:"response_id"`
	QuestionId           int64           `json:"question_id"`
	Type                 string          `json:"type"`
	Response             string          `json:"response"`
	MetaData             MetaData        `json:"metadata"`
}

type ResponseBankData struct {
	ResponseBankId           int64           `json:"response_bank_id"`
	UserId                   int64           `json:"user_id"`
	IsPublished              bool            `json:"is_published"`
	Responses                []ResponseData  `json:"responses"`
}
