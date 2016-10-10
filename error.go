package anilist

import (
	"encoding/json"
	"fmt"
)

type errorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Status   int      `json:"status"`
	Messages []string `json:"messages"`
}

func (e Error) Error() string {
	return fmt.Sprintf("api response error: status: %d, messages: %s", e.Status, e.Messages)
}

func UnmarshalErrorResponse(response []byte) (result Error) {
	var (
		content errorResponse
	)

	err := json.Unmarshal(response, &content)
	if err != nil {
		panic(err)
	}

	result = content.Error
	return
}
