package model

type AdvanceFilterPayload struct {
	ModelType  string                 `json:"modelType"`
	Conditions map[string]interface{} `json:"conditions"`
	Page       int                    `json:"page"`
	PageSize   int                    `json:"pageSize"`
}
