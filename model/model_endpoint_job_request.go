package model

type EndpointJobRequest struct {
	// Container image which will be used to execute the job
	Image string `json:"Image"`
	// Content of the job script
	FileContent string `json:"FileContent"`
}
