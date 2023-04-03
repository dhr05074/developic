package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type Body struct {
	Code string `json:"code"`
}

type ProblemBody struct {
	Factors  []string `json:"factors"`
	Language string   `json:"language"`
}

func main() {
	apiKey := os.Getenv("CHATGPT_API_KEY")
	cli := openai.NewClient(apiKey)
	app := echo.New()

	app.POST("/analyze", func(c echo.Context) error {
		var body Body
		if err := c.Bind(&body); err != nil {
			return err
		}

		decoded, err := base64.StdEncoding.DecodeString(body.Code)
		if err != nil {
			return err
		}

		prompt := `
Infer the programmer's five strengths and five weaknesses by the following code and convert the result to JSON format.

For example:
{
	"strengths": [
		"Readable and maintainable code",
		"Good understanding of Go language"	
	],	
	"weaknesses": [	
		"Complex code",	
		"Not secure"
	],
	"score": 80
}
		`

		content := fmt.Sprintf("%s\n\n%s", prompt, string(decoded))

		// Analyze the code
		response, err := cli.CreateChatCompletion(c.Request().Context(), openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
			Messages: []openai.ChatCompletionMessage{{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			}},
		})
		if err != nil {
			return err
		}

		var responseMap map[string]interface{}
		if err := json.Unmarshal([]byte(response.Choices[0].Message.Content), &responseMap); err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(200, responseMap)
	})

	if err := app.Start(":8080"); err != nil {
		panic(err)
	}
}
