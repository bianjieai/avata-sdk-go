package models

const (
	CreateUsers = "/v3/users" // 创建钱包用户
	EditUsers   = "/v3/users" // 更新钱包用户
	QueryUsers  = "/v3/users" // 查询钱包用户信息

)

type CreateUsersReq struct {
	UserType           int    `json:"user_type"`
	Name               string `json:"name"`
	PhoneNum           string `json:"phone_num"`
	Region             int    `json:"region"`
	CertificateType    int    `json:"certificate_type"`
	CertificateNum     string `json:"certificate_num"`
	RegistrationRegion int    `json:"registration_region"`
	RegistrationNum    string `json:"registration_num"`
	BusinessLicense    string `json:"business_license"`
	Email              string `json:"email"`
}

type CreateUsersRes struct {
	Data struct {
		UserID string `json:"user_id"`
		Did    string `json:"did"`
	} `json:"data"`
}

type EditUsersReq struct {
	UserID   string `json:"user_id"`
	PhoneNum string `json:"phone_num"`
}

type QueryUsersReq struct {
	UserType string `json:"user_type"`
	Code     string `json:"code"`
}

type QueryUsersRes struct {
	Data struct {
		UserID string `json:"user_id"`
		Did    string `json:"did"`
	} `json:"data"`
}
