package request

//https://github.com/go-playground/validator

type AddVideoClass struct {
	Name string `form:"name" json:"name" binding:"required"`
	Sort int    `form:"sort" json:"sort" binding:"required"`
}

func (AddVideoClass) GetMessages() ValidateErrorMessages {
	return ValidateErrorMessages{
		"Name.required": "分类名称不能为空",
		"Sort.required": "排序不能为空",
	}
}
