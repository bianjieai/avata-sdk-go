package models

const (
	CreateUsers = "/v2/users" // 创建用户
	EditUsers   = "/v2/users" // 更新用户
	QueryUsers  = "/v2/users" // 查询用户
)

//创建用户
type CreateUsersReq struct {
	UserType           int    `json:"user_type"`           //用户类型， 对于创建的钱包用户的类型属性 1：个人 2：企业
	Name               string `json:"name"`                //用户的真实姓名， 支持汉字以及大小写字母、空格
	Region             int    `json:"region"`              //所属 国家/地区 1：其他 2：中国大陆（默认） 3：中国香港 4：中国台湾
	CertificateType    int    `json:"certificate_type"`    //用户证件类型 1：身份证（默认） 2：护照
	CertificateNum     string `json:"certificate_num"`     //用户证件号码
	PhoneNum           string `json:"phone_num"`           //联系人或授权人手机号
	RegistrationRegion int    `json:"registration_region"` //企业注册地址 1：其他 2：中国大陆（默认） 3：中国香港 4：中国台湾
	RegistrationNum    string `json:"registration_num"`    //企业的统一社会信用代码或机构注册号
	BusinessLicense    string `json:"business_license"`    //营业执照或认证授权人声明书
	Email              string `json:"email"`               //企业邮箱，不支持汉字以及非规范性特殊字符。
}

//创建用户返回结果
type CreateUsersRes struct {
	Data struct {
		UserID string `json:"user_id"` //用户唯一标识
		DID    string `json:"did"`     //用户身份标识(预留字段)
	} `json:"data"`
}

//更新用户
type EditUsersReq struct {
	UserID   string `json:"user_id"`   //用户唯一标识
	PhoneNum string `json:"phone_num"` //联系人或授权人手机号
}

//查询用户信息
type QueryUsersReq struct {
	UserType string `json:"user_type"` //用户认证时，填写的用户类型。"1"：个人，"2"：企业
	Code     string `json:"code"`      //user_type为 "1" 时，请填写提交认证时对应的个人信息（身份证或护照号）user_type为 "2" 时，请填写提交认证时对应的企业信息（企业的统一社会信用代码或机构注册号）参数需要进行hash操作，hash算法为：sha-256
}

//查询用户信息返回结果
type QueryUsersRes struct {
	Data struct {
		UserID string `json:"user_id"` //用户唯一标识
		DID    string `json:"did"`     //联系人或授权人手机号
	} `json:"data"`
}
