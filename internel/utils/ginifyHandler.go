package utils

import (
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/gin-gonic/gin"
)

func GinifyHandler(
	lambdaHandler func(events.APIGatewayProxyRequest, string, dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error),
	db dynamodbiface.DynamoDBAPI,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, _ := io.ReadAll(c.Request.Body)

		req := events.APIGatewayProxyRequest{
			HTTPMethod:            c.Request.Method,
			QueryStringParameters: map[string]string{},
			Body:                  string(bodyBytes),
		}

		// query params
		for k, v := range c.Request.URL.Query() {
			if len(v) > 0 {
				req.QueryStringParameters[k] = v[0]
			}
		}

		resp, err := lambdaHandler(req, "Users", db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(resp.StatusCode)
		if resp.Body != "" {
			c.Writer.Write([]byte(resp.Body))
		}
	}
}
