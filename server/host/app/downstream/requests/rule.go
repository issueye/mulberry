package requests

import (
	"carambola/host/app/downstream/model"
	commonModel "carambola/host/common/model"
	"encoding/json"
)

type CreateRule struct {
	model.RuleInfo
}

func NewCreateRule() *CreateRule {
	return &CreateRule{
		RuleInfo: model.RuleInfo{},
	}
}

type UpdateRule struct {
	model.RuleInfo
}

func NewUpdateRule() *UpdateRule {
	return &UpdateRule{
		RuleInfo: model.RuleInfo{},
	}
}

func (req *CreateRule) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryRule struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryRule() *commonModel.PageQuery[*QueryRule] {
	return commonModel.NewPageQuery(0, 0, &QueryRule{})
}
