package models

const (
	CreateRecord = "/v1beta1/record/records" // 数字作品存证接口
)

// CreateRecordReq 数字作品存证接口请求参数
type CreateRecordReq struct {
	IdentityType int    `json:"identity_type,omitempty"`
	IdentityName string `json:"identity_name,omitempty"`
	IdentityNum  string `json:"identity_num,omitempty"`
	Type         int    `json:"type"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Hash         string `json:"hash"`
	HashType     int    `json:"hash_type"`
	OperationId  string `json:"operation_id"`
}
