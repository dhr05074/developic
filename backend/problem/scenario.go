package problem

import (
	"code-connect/problem/ent"
	"code-connect/problem/ent/scenario"
	"context"
	"encoding/json"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"regexp"
)

type (
	CreateScenarioRequest struct {
		Qualification string `json:"qualification"`
		Preferences   string `json:"preferences"`
	}

	CreateScenarioResponse struct {
		RequestID string `json:"request_id"`
	}

	Scenario struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
)

func (h *Handler) CreateScenario(_ context.Context, req CreateScenarioRequest) (*CreateScenarioResponse, error) {
	requestID := h.generateID()

	go func(entClient *ent.Client) {
		newCtx := context.TODO()
		newGPTClient := h.gptClient.NewClientWithEmptyContext()

		prompts := make([]string, 0, 2)
		prompts = append(prompts, "Read the following preferences and deduce the characteristics of the developer this company wants and make it a bullet point list.")
		prompts = append(prompts, req.Preferences)

		_, err := newGPTClient.CompleteWithContext(newCtx, prompts)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		prompts = make([]string, 0, 2)
		prompts = append(prompts, "Read the following qualification and deduce the abilities of the developer this company wants and make it a bullet point list.")
		prompts = append(prompts, req.Qualification)

		_, err = newGPTClient.CompleteWithContext(newCtx, prompts)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		prompts = make([]string, 0, 2)
		prompts = append(prompts, "Considering the aforementioned qualifications and preferential conditions, create 5 coding assignment topics.")
		prompts = append(prompts, "Arrange them in the form of a JSON array with the following format: [{\"title\": \"title\", \"content\": \"content\"}, ...]")

		answer, err := newGPTClient.CompleteWithContext(newCtx, prompts)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		// Regular expression to extract only JSON code from GPT's natural language (Markdown) answers
		re := regexp.MustCompile("\\[\\n([\\s\\S]*?)\\n]")
		matches := re.FindString(answer)

		scenarios := make([]*Scenario, 0, 5)
		err = json.Unmarshal([]byte(matches), &scenarios)
		if err != nil {
			l.Errorw("error while unmarshalling", "error", err)
			return
		}

		extractedScenarios := make([]*ent.ScenarioCreate, 0, 5)
		for _, s := range scenarios {
			extractedScenarios = append(extractedScenarios, h.entClient.Scenario.Create().SetUUID(h.generateID()).SetTitle(s.Title).SetContent(s.Content).SetRequestID(requestID))
		}

		tx, err := entClient.Tx(newCtx)
		if err != nil {
			l.Errorw("error while creating transaction", "error", err)
			return
		}

		defer func() {
			if err == nil {
				return
			}

			if err := tx.Rollback(); err != nil {
				l.Errorw("error while rolling back transaction", "error", err)
			}
		}()

		_, err = tx.Scenario.CreateBulk(extractedScenarios...).Save(newCtx)
		if err != nil {
			l.Errorw("error while saving scenarios", "error", err)
			return
		}

		if err = tx.Commit(); err != nil {
			l.Errorw("error while committing transaction", "error", err)
			return
		}
	}(h.entClient)

	return &CreateScenarioResponse{
		RequestID: requestID,
	}, nil
}

func (h *Handler) generateID() string {
	return gonanoid.MustGenerate(CodeAlphabets, 8)
}

func (h *Handler) ScenariosByRequestID(ctx context.Context, requestID string) ([]*Scenario, error) {
	tx, err := h.entClient.Tx(ctx)
	if err != nil {
		return nil, err
	}

	entScenarios, err := tx.Scenario.Query().ForUpdate().Where(scenario.RequestID(requestID)).All(ctx)
	if err != nil {
		return nil, err
	}

	ss := make([]*Scenario, 0, 5)
	for _, sc := range entScenarios {
		ss = append(ss, &Scenario{
			ID:      sc.UUID,
			Title:   sc.Title,
			Content: sc.Content,
		})
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ss, nil
}
