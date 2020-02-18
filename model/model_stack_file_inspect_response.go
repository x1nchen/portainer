package model

type StackFileInspectResponse struct {
	// Content of the Stack file.
	StackFileContent string `json:"StackFileContent,omitempty"`
}
