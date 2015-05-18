package xero

import (
	"fmt"
	"strings"
)

type APIException struct {
	Type        string                 `json:"Type",xml:"Type"`
	ErrorNumber int                    `json:"ErrorNumber",xml:"ErrorNumber"`
	Message     string                 `json:"Message",xml:"Message"`
	Elements    []APIValidationElement `xml:"Elements>DataContractBase"`
	HTTPStatus  int
}

type APIValidationElement struct {
	Errors   []string `xml:"ValidationErrors>ValidationError>Message"`
	Warnings []string `xml:"Warnings>Warning>Message"`
}
type APIValidationError struct {
	Message string `json:"Message"`
}

func (e *APIException) Error() string {
	return fmt.Sprintf("Xero Exception %d: (%s) %s", e.ErrorNumber, e.Type, e.Message)
}

func (e *APIException) GetMessages() []string {
	messages := map[string]bool{}
	if e.Elements != nil {
		for _, element := range e.Elements {
			for _, txt := range element.Errors {
				messages[txt] = true
			}
			for _, txt := range element.Warnings {
				messages[txt] = true
			}
		}
	}

	var messageArray = make([]string, 0, len(messages))
	for m, _ := range messages {
		messageArray = append(messageArray, m)
	}
	return messageArray
}
func (e *APIException) GetUserDescription() string {
	if e.Elements != nil && len(e.Elements) > 0 && len(e.Elements[0].Errors) > 0 {
		messages := map[string]bool{}
		for _, m := range e.Elements[0].Errors {
			messages[m] = true
		}
		messageArray := []string{"Object was invalid"}
		for m, _ := range messages {
			messageArray = append(messageArray, m)
		}
		return strings.Join(messageArray, ", ")
	}

	return e.Message
}

func (e *APIException) GetHTTPStatus() int {
	return e.HTTPStatus
}
