package requests

import (
	"carambola/host/app/downstream/model"
	commonModel "carambola/host/common/model"
	"encoding/json"
)

type CreatePort struct {
	model.PortInfo
}

func NewCreatePort() *CreatePort {
	return &CreatePort{
		PortInfo: model.PortInfo{},
	}
}

type UpdatePort struct {
	model.PortInfo
}

func NewUpdatePort() *UpdatePort {
	return &UpdatePort{
		PortInfo: model.PortInfo{},
	}
}

func (req *CreatePort) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryPort struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryPort() *commonModel.PageQuery[*QueryPort] {
	return commonModel.NewPageQuery(0, 0, &QueryPort{})
}