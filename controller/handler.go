package controller

import (
	"io"
	"net/http"
	"os"
	"strings"

	"final/model"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("csvfile")
	if err != nil {
		c.String(http.StatusBadRequest, "Error retrieving the file")
		return
	}
	defer file.Close()

	csvData, err := io.ReadAll(file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading the file")
		return
	}

	table, err := model.CsvToSlice(string(csvData))
	if err != nil {
		c.String(http.StatusInternalServerError, "Error converting CSV to slice: "+err.Error())
		return
	}

	queries := c.PostForm("queries")
	queryList := strings.Split(queries, ";")

	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		c.String(http.StatusInternalServerError, "HUGGINGFACE_TOKEN is required")
		return
	}

	client := &http.Client{}
	connector := model.AIModelConnector{Client: client}

	var responses []model.Response
	for _, query := range queryList {
		query = strings.TrimSpace(query)
		input := model.Inputs{
			Table: table,
			Query: query,
		}

		response, err := connector.ConnectAIModel(input, token)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error connecting to AI model: "+err.Error())
			return
		}
		response.Recommendation = model.GenerateRecommendation(response)
		responses = append(responses, response)
	}

	c.HTML(http.StatusOK, "response.html", responses)
}
