package request

type VideoSetting struct {
	Key   string `form:"key" json:"key" binding:"required"`
	Value string `form:"value" json:"value" binding:"required"`
	Notes string `form:"notes" json:"notes"`
}

func (VideoSetting) GetMessages() ValidateErrorMessages {
	return ValidateErrorMessages{
		"Key.required":   "key不能为空",
		"Value.required": "value不能为空",
	}
}
