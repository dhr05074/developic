package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

type Body struct {
	Statement string   `json:"statement"`
	Points    []string `json:"points"`
	Code      string   `json:"code"`
}

type ProblemRequestBody struct {
	Difficulty int32  `json:"difficulty"`
	Language   string `json:"language"`
}

type ProblemResponseBody struct {
	Statement   string    `json:"statement"`
	Description string    `json:"description"`
	Examples    []Example `json:"examples"`
}

type Example struct {
	Input  any `json:"input"`
	Output any `json:"output"`
}

func main() {
	apiKey := os.Getenv("CHATGPT_API_KEY")
	cli := openai.NewClient(apiKey)
	app := echo.New()
	app.Use(middleware.CORS())

	gptClient := ChatGPTClient{cli: cli}

	app.POST("/problem", func(c echo.Context) error {
		var body ProblemRequestBody
		if err := c.Bind(&body); err != nil {
			return err
		}

		if body.Difficulty == 0 {
			body.Difficulty = 50
		}

		if body.Language == "" {
			body.Language = "Go"
		}

		sample := ProblemResponseBody{
			Statement:   "Powers of Two",
			Description: "Have the function `PowersofTwo(num)` take the num parameter being passed which will be an integer and return the string true if it's a power of two. If it's not return the string false. For example if the input is 16 then your program should return the string true but if the input is 22 then the output should be the string false.",
			Examples: []Example{
				{
					Input:  "4",
					Output: "true",
				},
				{
					Input:  "124",
					Output: "false",
				},
			},
		}
		sampleJSON, _ := json.MarshalIndent(sample, "", "  ")

		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Make a real world scenario briefly for a developer. The complexity should be %d out of 100. Complexity is a measure of how complex a problem situation is. The higher it is, the more complicated the situation. Exclude the AI technology.", body.Difficulty),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Using the previous scenario, select one functionality and make only one coding test to implement it in %s. It should evaluate the technical skills, problem solving and analytical skills of the developer. The difficulty is %d out of 100. You should give at least two examples.", body.Language, body.Difficulty),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Don't use the markdown syntax. You can just use the JSON syntax. Don't include your words.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Convert the result to JSON like this:",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: string(sampleJSON),
			},
		}

		msg, err := gptClient.CompleteChatWithContext(c.Request().Context(), messages)
		if err != nil {
			return err
		}

		//messages = []openai.ChatCompletionMessage{
		//
		//}
		//
		//msg, err := gptClient.CompleteChatWithContext(c.Request().Context(), messages)
		//if err != nil {
		//	return err
		//}

		log.Println(msg)

		var response ProblemResponseBody
		if err := json.Unmarshal([]byte(msg), &response); err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(201, response)
	})

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
Please evaluate the characteristics of the programmer based on the following code and give the result in a JSON format.
You must consider the following statement and points when evaluating the code.
`

		// Analyze the code
		response, err := cli.CreateChatCompletion(c.Request().Context(), openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo0301,
			Temperature: 1,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: string(decoded),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("The statement is %s and the points are %v.", body.Statement, body.Points),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Convert the result to JSON format has four keys",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "The first key is keyword which is array of string type that has 5 short keywords which describe the developer's characteristics well.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "The second key is score which is integer type that has the score of the code out of 100.0",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "The third key is strengths which is array of string type that has 5 elements which have the strengths of the code",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "The fourth key is weaknesses which is array of string type that has 5 elements which have the weaknesses of the code",
				},
			},
		})
		if err != nil {
			return err
		}

		log.Println(response.Choices[0].Message.Content)

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

type ChatGPTClient struct {
	cli      *openai.Client
	messages []openai.ChatCompletionMessage
}

func (c *ChatGPTClient) CompleteChatWithContext(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	c.messages = append(c.messages, messages...)

	response, err := c.cli.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo0301,
		Temperature: 1,
		Messages:    c.messages,
	})
	if err != nil {
		return "", err
	}

	c.messages = append(c.messages, response.Choices[0].Message)

	return response.Choices[0].Message.Content, nil
}
