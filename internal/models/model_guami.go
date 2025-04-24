package models

type Guami struct {
	PlmnId *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
	AmfId  string  `json:"amfId" yaml:"amfId" bson:"amfId" mapstructure:"AmfId"`
}
