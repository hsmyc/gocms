package richtexteditor

import (
	"encoding/json"
)

func RichTextEditor(data string, specArray []string) []byte {
	result, err := json.Marshal(struct {
		Data      string   `json:"data"`
		SpecArray []string `json:"specArray"`
	}{
		Data:      data,
		SpecArray: specArray,
	})
	if err != nil {
		panic(err)
	}

	return result
}
