package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ditto "github.com/julioc98/anbima/goapi/ditto"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req Request) (Response, error) {
	var buf bytes.Buffer
	var body []byte
	var err error
	var in ditto.Ditto
	var out ditto.Ditto
	ok := true

	switch req.HTTPMethod {
	case "POST":
		in, err = ditto.ByteToMap([]byte(req.Body))
		out = in
		body, err = json.Marshal(out)
		if in["file"] == "fundo-1" {
			ok = false
		}
	case "GET":
		body, err = json.Marshal(ditto.Ditto{
			"message": "Go JC GET!",
		})
	default:
		body, err = json.Marshal(ditto.Ditto{
			"message": "Go JC Default!",
		})
	}

	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	respOK := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
			"X-MyCompany-Func-Reply":      "check-handler",
		},
	}

	respNotOK := Response{
		StatusCode:      400,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
			"X-MyCompany-Func-Reply":      "check-handler",
		},
	}

	switch ok {
	case true:
		return respOK, nil
	case false:
		return respNotOK, nil
	default:
		return respOK, nil
	}

}

func main() {
	lambda.Start(Handler)
}
