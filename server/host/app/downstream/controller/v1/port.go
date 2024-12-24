package v1

import (
	"carambola/host/app/downstream/logic"
	"carambola/host/app/downstream/requests"
	"carambola/host/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePort doc
//
//	@tags			端口号管理
//	@Summary		添加端口号信息
//	@Description	添加端口号信息
//	@Produce		json
//	@Param			body	body		requests.CreatePort	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/create [post]
//	@Security		ApiKeyAuth
func CreatePort(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreatePort()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreatePort(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdatePort doc
//
//	@tags			端口号管理
//	@Summary		修改端口号信息
//	@Description	修改端口号信息
//	@Produce		json
//	@Param			body	body		requests.UpdatePort	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/update [put]
//	@Security		ApiKeyAuth
func UpdatePort(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdatePort()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdatePort(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeletePort doc
//
//	@tags			端口号管理
//	@Summary		删除端口号信息
//	@Description	删除端口号信息
//	@Produce		json
//	@Param			id		path	int	true	"端口号id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/delete [delete]
//	@Security		ApiKeyAuth
func DeletePort(c *gin.Context) {
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

	err = logic.DeletePort(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// PortList doc
//
//	@tags			端口号管理
//	@Summary		端口号列表
//	@Description	端口号列表
//	@Produce		json
//	@Param			body	body		requests.QueryPort	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/list [post]
//	@Security		ApiKeyAuth
func PortList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryPort()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.PortList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetPort doc
//
//	@tags			端口号管理
//	@Summary		端口号详情
//	@Description	端口号详情
//	@Produce		json
//	@Param			id		path	int	true	"端口号id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/downstream/get [get]
//	@Security		ApiKeyAuth
func GetPort(c *gin.Context) {
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

	info, err := logic.GetPort(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}
