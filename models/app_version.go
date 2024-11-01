package models

type AppVersion struct {
	ID          uint   `json:"id"`
	Platform    string `json:"platform"`
	UpdateUrl   string `json:"update_url"`
	IsForced    int8   `json:"is_forced"`
	Description string `json:"description"`
}

func (AppVersion) TableName() string {
	return "app_version"
}
