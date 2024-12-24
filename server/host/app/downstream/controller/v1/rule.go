package v1

import (
	"carambola/host/app/downstream/logic"
	"carambola/host/app/downstream/requests"
	"carambola/host/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRule doc
//
//	@tags			规则管理
//	@Summary		添加规则信息
//	@Description	添加规则信息
//	@Produce		json
//	@Param			body	body		requests.CreateRule	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/create [post]
//	@Security		ApiKeyAuth
func CreateRule(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateRule()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateRule(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateRule doc
//
//	@tags			规则管理
//	@Summary		修改规则信息
//	@Description	修改规则信息
//	@Produce		json
//	@Param			body	body		requests.UpdateRule	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/update [put]
//	@Security		ApiKeyAuth
func UpdateRule(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateRule()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateRule(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteRule doc
//
//	@tags			规则管理
//	@Summary		删除规则信息
//	@Description	删除规则信息
//	@Produce		json
//	@Param			id		path	int	true	"规则id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/delete [delete]
//	@Security		ApiKeyAuth
func DeleteRule(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.DeleteRule(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// RuleList doc
//
//	@tags			规则管理
//	@Summary		规则列表
//	@Description	规则列表
//	@Produce		json
//	@Param			body	body		requests.QueryRule	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/list [post]
//	@Security		ApiKeyAuth
func RuleList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryRule()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.RuleList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetRule doc
//
//	@tags			规则管理
//	@Summary		规则详情
//	@Description	规则详情
//	@Produce		json
//	@Param			id		path	int	true	"规则id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/get [get]
//	@Security		ApiKeyAuth
func GetRule(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	info, err := logic.GetRule(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}
