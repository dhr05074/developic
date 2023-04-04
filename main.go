package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sashabaranov/go-openai"
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
	Statement    string   `json:"statement"`
	Functions    []string `json:"functions"`
	Requirements []string `json:"requirements"`
	TestCases    []string `json:"test_cases"`
}

type Problem struct {
	ID         string   `json:"id"`
	Statement  string   `json:"statement"`
	Conditions []string `json:"conditions"`
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
			Statement: "Implement a search functionality in Go for the e-commerce platform. The function should take a search query string and return a list of products that match the query criteria. The search should take into account various attributes such as name, description, price, and category.",
			Functions: []string{"func searchProducts(query string, products []Product) ([]Product, error)"},
			Requirements: []string{
				"The search should be case-insensitive.",
				"The search should match partial words as well as whole words.",
				"The search should prioritize exact matches over partial matches.",
				"The search should allow for filtering by category.",
				"The search results should be sorted by relevance, with exact matches appearing first.",
			},
			TestCases: []string{
				"When given a search query \"shoe\", the function should return a list of products that have \"shoe\" in their name or description, sorted by relevance.",
				"When given a search query \"Nike running shoes\", the function should return a list of products that have \"Nike\" and \"running\" and \"shoes\" in their name or description, sorted by relevance.",
				"When given a search query \"dress\" and a category filter of \"women's clothing\", the function should return a list of women's dresses that have \"dress\" in their name or description, sorted by relevance.",
			},
		}
		sampleJSON, _ := json.Marshal(sample)

		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Make a real world scenario briefly for a developer. The complexity should be %d out of 100. Exclude the AI technology.", body.Difficulty),
			},
		}

		_, err := gptClient.CompleteChatWithContext(c.Request().Context(), messages)
		if err != nil {
			return err
		}

		messages = []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Using the previous scenario, select one functionality and make a coding test to implement it in %s. It should evaluate the technical skills, problem solving and analytical skills of the developer. Don’t give the sample code, and hints. Give a problem statement, a function signature, five requirements and three test cases. The difficulty is %d out of 100. You must give the result as a JSON format like following example.", body.Language, body.Difficulty),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Example should be like this:",
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
