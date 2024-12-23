package requests

import (
	"carambola/host/app/task/model"
	commonModel "carambola/host/common/model"
	"encoding/json"
)

type CreateClient struct {
	model.Client
}

func NewCreateClient() *CreateClient {
	return &CreateClient{
		Client: model.Client{},
	}
}

type UpdateClient struct {
	model.Client
}

func NewUpdateClient() *UpdateClient {
	return &UpdateClient{
		Client: model.Client{},
	}
}

func (req *CreateClient) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryClient struct {
	KeyWords string `json:"keywords" form:"keywords"`   // 关键词
	IsHidden int    `json:"is_hidden" form:"is_hidden"` // 0 不隐藏 1 隐藏
}

func NewQueryClient() *commonModel.PageQuery[*QueryClient] {
	return commonModel.NewPageQuery(0, 0, &QueryClient{})
}
