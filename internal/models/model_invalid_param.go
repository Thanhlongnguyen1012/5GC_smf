package models

type InvalidParam struct {
	Param  string `json:"param,omitempty"`
	Reason string `josn:"reason,omitempty"`
}
