package myjson

const (
	LAYER_SUBMIT_NORMAL = iota
	LAYER_SUBMIT_FIRST
	LAYER_SUBMIT_SECOND
)

type Element struct {
	RadioChkboxVal string `json:"radioChkboxVal"`
	Type           string `json:"type"`
	Label          string `json:"label"`
	Validate       string `json:"validate"`
	Unique         string `json:"unique"`
	AllowEmpty     int    `json:"allowEmpty"`
	ID             int64  `json:"id"`
}

//easyjson:json
type AdvFormData struct {
	ID     int64 `json:"id"`
	FormID int64 `json:"form_id"`
	AdID   int64 `json:"ad_id"`
	SiteID int64 `json:"site_id"`
	Status int   `json:"status"`
	Data   struct {
		Elements map[string]string `json:"elements"`
	} `json:"data"`
	DataMap       map[string]string `json:"datamap"`
	CreateTime    string            `json:"create_time"`
	DeviceID      int64             `json:"device_id"`
	UserID        int64             `json:"user_id"`
	ReqID         string            `json:"req_id"`
	ValidateCheck int               `json:"validate_check"`
	SpamReason    int               `json:"spam_reason"`
}

//easyjson:json
type AdvForm struct {
	ID         int64  `json:"id"`
	AdvID      int64  `json:"adv_id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	ModifyTime string `json:"modify_time"`
	Status     int    `json:"status"`
	Data       struct {
		ExtraElements []Element `json:"extraElements"`
		Elements      []Element `json:"elements"`
		Caculator     struct {
			Active bool `json:"active"`
		} `json:"caculator"`
	} `json:"data"`
	Permission string `json:"permission"`
	SmartCheck int    `json:"smart_check"`
	Version    int    `json:"version"`
}

//easyjson:json
type FormDataMQ struct {
	Type        string      `json:"type"`
	AdvFormData AdvFormData `json:"advformdata"`
	AdvForm     AdvForm     `json:"advform"`
	Spamed      int         `json:"spamed"`
	Export      string      `json:"export"`
	LogID       string      `json:"log_id"`
	PageID      string      `json:"page_id"`
	PhoneNumber string      `json:"phone_number"`
	CustomSms   string      `json:"custom_sms"`
	Layer       int         `json:"is_layer"`
}

type ExportItem struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}
