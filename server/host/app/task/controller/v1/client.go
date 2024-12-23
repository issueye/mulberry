package v1

import (
	"carambola/host/app/task/logic"
	"carambola/host/app/task/requests"
	"carambola/host/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateClient doc
//
//	@tags			客户端管理
//	@Summary		添加客户端信息
//	@Description	添加客户端信息
//	@Produce		json
//	@Param			body	body		requests.CreateTask	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/task/create [post]
//	@Security		ApiKeyAuth
func CreateClient(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateClient()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateClient(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateClient doc
//
//	@tags			客户端管理
//	@Summary		修改客户端信息
//	@Description	修改客户端信息
//	@Produce		json
//	@Param			body	body		requests.UpdateClient	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/task/update [put]
//	@Security		ApiKeyAuth
func UpdateClient(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateClient()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateClient(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteClient doc
//
//	@tags			客户端管理
//	@Summary		删除客户端信息
//	@Description	删除客户端信息
//	@Produce		json
//	@Param			id		path	int	true	"客户端id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/task/delete [delete]
//	@Security		ApiKeyAuth
func DeleteClient(c *gin.Context) {
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

	err = logic.DeleteClient(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// ClientList doc
//
//	@tags			客户端管理
//	@Summary		客户端列表
//	@Description	客户端列表
//	@Produce		json
//	@Param			body	body		requests.QueryClient	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/task/list [post]
//	@Security		ApiKeyAuth
func ClientList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryClient()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.ClientList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetClient doc
//
//	@tags			客户端管理
//	@Summary		客户端详情
//	@Description	客户端详情
//	@Produce		json
//	@Param			id		path	int	true	"客户端id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/task/get [get]
//	@Security		ApiKeyAuth
func GetClient(c *gin.Context) {
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

	info, err := logic.GetClient(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}

// CloseClient doc
//
//	@tags			客户端管理
//	@Summary		关闭客户端
//	@Description	关闭客户端
//	@Produce		json
//	@Param			id		path	int	true	"客户端id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/task/close [put]
//	@Security		ApiKeyAuth
func CloseClient(c *gin.Context) {
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

	err = logic.CloseClient(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// OpenClient doc
//
//	@tags			客户端管理
//	@Summary		开启客户端
//	@Description	开启客户端
//	@Produce		json
//	@Param			id		path	int	true	"客户端id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/task/open [put]
//	@Security		ApiKeyAuth
func OpenClient(c *gin.Context) {
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

	err = logic.RunClient(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
