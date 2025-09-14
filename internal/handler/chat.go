package handler

import (
	"encoding/json"
	"go-web-crud/internal/models"
	"os"

	"github.com/gofiber/fiber/v2"
)

func GetSimpleChatResponse(c *fiber.Ctx) error {
	openai_key := string(os.Getenv("OPENAI_API_KEY"))

	input := new(models.SimpleChatRequest)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON:" + err.Error(),
		})
	}

	requestBody := models.OpenAIRequest{
		Model: "qwen/qwen-plus-2025-07-28",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "user",
				Content: input.Message,
			},
		},
	}
	body_bytes, err := json.Marshal(requestBody)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON:" + err.Error(),
		})
	}
	agent := fiber.Post("https://openrouter.ai/api/v1/chat/completions")
	agent = agent.Add("Content-Type", "application/json")
	agent = agent.Add("Authorization", "Bearer "+openai_key)

	agent.Body(body_bytes)

	statusCode, responseBody, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "API request failed",
			"details": errs[0].Error(),
		})
	}

	if statusCode != fiber.StatusOK {
		return c.Status(statusCode).JSON(fiber.Map{
			"error":   "OpenAI API returned a non-200 status code",
			"details": string(responseBody),
		})
	}

	var openAIResp models.OpenAIResponse

	if err := json.Unmarshal(responseBody, &openAIResp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot parse OpenAI response: " + err.Error(),
		})
	}

	if len(openAIResp.Choices) > 0 {
		// Return the content of the first message choice to the client
		return c.JSON(fiber.Map{
			"response": openAIResp.Choices[0].Message.Content,
		})
	}

	return c.JSON(fiber.Map{
		"response": "No response content received from OpenAI.",
	})
}
