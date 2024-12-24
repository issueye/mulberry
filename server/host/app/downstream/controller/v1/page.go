package v1

import (
	"carambola/host/app/downstream/logic"
	"carambola/host/app/downstream/requests"
	"carambola/host/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePage doc
//
//	@tags			页面管理
//	@Summary		添加页面信息
//	@Description	添加页面信息
//	@Produce		json
//	@Param			body	body		requests.CreatePage	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/create [post]
//	@Security		ApiKeyAuth
func CreatePage(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreatePage()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreatePage(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdatePage doc
//
//	@tags			页面管理
//	@Summary		修改页面信息
//	@Description	修改页面信息
//	@Produce		json
//	@Param			body	body		requests.UpdatePage	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/update [put]
//	@Security		ApiKeyAuth
func UpdatePage(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdatePage()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdatePage(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeletePage doc
//
//	@tags			页面管理
//	@Summary		删除页面信息
//	@Description	删除页面信息
//	@Produce		json
//	@Param			id		path	int	true	"页面id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/delete [delete]
//	@Security		ApiKeyAuth
func DeletePage(c *gin.Context) {
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

	err = logic.DeletePage(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// PageList doc
//
//	@tags			页面管理
//	@Summary		页面列表
//	@Description	页面列表
//	@Produce		json
//	@Param			body	body		requests.QueryPage	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/list [post]
//	@Security		ApiKeyAuth
func PageList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryPage()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.PageList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetPage doc
//
//	@tags			页面管理
//	@Summary		页面详情
//	@Description	页面详情
//	@Produce		json
//	@Param			id		path	int	true	"页面id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/get [get]
//	@Security		ApiKeyAuth
func GetPage(c *gin.Context) {
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

	info, err := logic.GetPage(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}
