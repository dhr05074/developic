package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strings"
)

type Body struct {
	Statement string   `json:"statement"`
	Points    []string `json:"points"`
	Code      string   `json:"code"`
}

type ProblemBody struct {
	Factors  []string `json:"factors"`
	Language string   `json:"language"`
}

type ProblemResponse struct {
	Statement   string    `json:"statement"`
	SubProblems []Problem `json:"sub_problems"`
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

	app.POST("/problem", func(c echo.Context) error {
		var body ProblemBody
		if err := c.Bind(&body); err != nil {
			return err
		}

		sampleResponse := ProblemResponse{
			Statement: "You should create an internet shopping mall service",
			SubProblems: []Problem{
				{
					Statement: "Design and create a database schema to support an internet shopping mall service. The schema should be able to handle the following requirements",
					Conditions: []string{
						"Each product can have multiple categories.",
						"Each product has a unique ID, name, description, image, price, and inventory count.",
						"Users can view products and add them to their shopping cart.",
						"Users can create an account and log in to the service.",
						"Users can view their order history.",
						"Users can place orders and receive an order confirmation with a unique order ID.",
					},
				},
			},
		}
		sampleJSON, _ := json.Marshal(sampleResponse)

		result, err := cli.CreateChatCompletion(c.Request().Context(), openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo0301,
			Temperature: 1,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "You should create a coding test which is composed of four sub-problems which has a statement and conditions to evaluate the characteristics of the programmer based on the following conditions and program language and give a test in a JSON format.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Example is like this:",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: string(sampleJSON),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Conditions: \n%s", strings.Join(body.Factors, "\n")),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Language: %s\n", body.Language),
				},
			},
		})
		if err != nil {
			return err
		}

		log.Println(result.Choices[0].Message.Content)

		var response ProblemResponse
		if err := json.Unmarshal([]byte(result.Choices[0].Message.Content), &response); err != nil {
			log.Println(err)
			return err
		}

		for i := range response.SubProblems {
			response.SubProblems[i].ID = gonanoid.Must(7)
		}

		return c.JSON(200, response)
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
