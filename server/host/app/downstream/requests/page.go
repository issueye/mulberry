package requests

import (
	"carambola/host/app/downstream/model"
	commonModel "carambola/host/common/model"
	"encoding/json"
)

type CreatePage struct {
	model.PageInfo
}

func NewCreatePage() *CreatePage {
	return &CreatePage{
		PageInfo: model.PageInfo{},
	}
}

type UpdatePage struct {
	model.PageInfo
}

func NewUpdatePage() *UpdatePage {
	return &UpdatePage{
		PageInfo: model.PageInfo{},
	}
}

func (req *CreatePage) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryPage struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryPage() *commonModel.PageQuery[*QueryPage] {
	return commonModel.NewPageQuery(0, 0, &QueryPage{})
}
