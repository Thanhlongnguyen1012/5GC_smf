package models

type ProblemDetails struct {
	Type          string         `json:"type, omitempty"`
	Title         string         `json:"title, omitempty"`
	Status        int32          `json:"status, omitempty"`
	Detail        string         `json:"detail, omitempty"`
	Instance      string         `json:"instance, omitempty"`
	Cause         string         `json:"cause, omitempty"`
	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`
}
