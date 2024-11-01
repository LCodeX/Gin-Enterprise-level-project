package models

type AppConfig struct {
	CustomerServicePhone string `json:"customer_service_phone"`
	WechatId             string `json:"wechat_id"`
	PcUrl                string `json:"pc_url"`
}

func (AppConfig) TableName() string {
	return "configurations"
}
