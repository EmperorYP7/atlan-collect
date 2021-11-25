package v1

import (
	"collect/api/v1/structures"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type collectController struct {
	dbObject *sql.DB
}

func NewCollectController(connector *sql.DB) *collectController {
	return &collectController{
		dbObject: connector,
	}
}

func (tc *collectController) GetFormHandler(c *gin.Context) {
	statusCode := 400
	res := structures.GetFormResponse{
		Status: "Failed",
		Message: "Failed to load response",
	}

	queryString := `
		SELECT
			f.id,
			f.user_id,
			f.name,
			f.description,
			f.created_at,
			f.is_published,
			qb.version,
			f.question_bank_id
		FROM
			Form f
		LEFT JOIN
			QuestionBank qb
		ON
			qb.form_id = f.id AND f.question_bank_id = qb.id
		WHERE
			f.user_id = ?
	`
	rows, err := tc.dbObject.Query(queryString, c.Request)

	if err != nil {
		log.Fatal(err.Error())
		c.JSON(statusCode, res)
		return
	}

	defer rows.Close()

	for rows.Next() {
		rowData := &res.FormData
		err = rows.Scan(&rowData.FormId, &rowData.UserId, &rowData.Name,
			&rowData.Description, &rowData.CreatedAt, &rowData.IsPublished,
			&rowData.Version)
		
		if err != nil {
			log.Fatal(err.Error())
			c.JSON(statusCode, res)
			return
		}
	}

	queryString = `
		SELECT
			q.id,
			q.text,
			q.type,
			q.is_published,
		FROM
			Question q
		WHERE
			q.form_id = ?
	`

	rows2, err := tc.dbObject.Query(queryString, res.FormData.FormId)

	if err != nil {
		log.Fatal(err.Error())
		c.JSON(statusCode, res)
		return
	}

	defer rows2.Close()

	for rows2.Next() {
		questionData := structures.QuestionData{}

		rows2.Scan(&questionData.FormId, &questionData.QuestionText,
		&questionData.Type, &questionData.IsPublished)

		if err != nil {
			log.Fatal(err.Error())
			c.JSON(statusCode, res)
			return
		}

		res.FormData.QuestionBank = append(res.FormData.QuestionBank, questionData)
	}

	res.Status = "Success"
	res.Message = "Data fetched successfully"
	statusCode = 200

	c.JSON(statusCode, res)
}

func (tc *collectController) GetQuestionHandler(c *gin.Context) {
	statusCode := 400
	res := structures.GetQuestionResponse{
		Status: "Failed",
		Message: "Failed to load response",
	}

	queryString := `
		SELECT
			q.id,
			q.text,
			q.type,
			q.is_published,
		FROM
			Question q
		WHERE
			q.id = ?
	`

	questionId, isExist := c.Params.Get("question_id")

	if !isExist {
		log.Fatal("Question ID not provided")
		c.JSON(statusCode, res)
		return
	}

	rows, err := tc.dbObject.Query(queryString, questionId)

	if err != nil {
		log.Fatal(err.Error())
		c.JSON(statusCode, res)
		return
	}

	defer rows.Close()

	for rows.Next() {
		questionData := res.QuestionData

		rows.Scan(&questionData.FormId, &questionData.QuestionText,
		&questionData.Type, &questionData.IsPublished)

		if err != nil {
			log.Fatal(err.Error())
			c.JSON(statusCode, res)
			return
		}
	}

	res.Status = "Success"
	res.Message = "Data fetched successfully"
	statusCode = 200

	c.JSON(statusCode, res)
}

func (tc *collectController) GetResponseHandler(c *gin.Context) {
	statusCode := 400
	res := structures.GetResponseResponse{
		Status: "Failed",
		Message: "Couldn't fetch data",
	}

	queryString := `
		SELECT
			rb.id,
			rb.user_id,
			rb.question_id,
			rb.is_published
		FROM
			ResponseBank rb
		WHERE
			rb.form_id = ? AND rb.version = ?
	`

	form_id, isExist := c.Params.Get("form_id")
	version, isExist2 := c.Params.Get("version")

	if !(isExist && isExist2) {
		log.Fatal("Form ID not provided")
		c.JSON(statusCode, res)
		return
	}

	rows, err := tc.dbObject.Query(queryString, form_id, version)

	if err != nil {
		log.Fatal("Error in query")
		c.JSON(statusCode, res)
		return
	}

	defer rows.Close()

	for rows.Next() {
		resBank := structures.ResponseBankData{}
		err := rows.Scan(&resBank.ResponseBankId, &resBank.UserId, &resBank.IsPublished)
		if err != nil {
			log.Fatal("Error in scanning")
			c.JSON(statusCode, res)
			return
		}

		subQueryString := `
			SELECT
				r.id,
				r.question_id,
				r.user_id,
				r.type
			FROM
				Response r
			WHERE
				r.form_id = ? AND r.response_bank_id = ?
		`

		rows2, err := tc.dbObject.Query(subQueryString, form_id, resBank.ResponseBankId)
		if err != nil {
			log.Fatal("Error in query")
			c.JSON(statusCode, res)
			return
		}

		for rows2.Next() {
			resResponse := structures.ResponseData{}
			err := rows2.Scan(&resResponse.ResponseId, &resResponse.QuestionId, &resResponse.Response)

			if err != nil {
				log.Fatal("Error in scanning")
				c.JSON(statusCode, res)
				return
			}

			resBank.Responses = append(resBank.Responses, resResponse)
		}

		res.ResponseBankData = append(res.ResponseBankData, resBank)
	}

	res.Status = "Success"
	res.Message = "Data fetched successfully"
	statusCode = 200

	c.JSON(statusCode, res)
}

func (tc *collectController) GetResponseBankHandler(c *gin.Context) {
	
}
