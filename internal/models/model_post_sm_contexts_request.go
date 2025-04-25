package models

type PostSmContextsRequest struct {
	JsonData *SmContextCreateData `json:"jsonData,omitempty" multipart:"contentType:application/json"`
}
