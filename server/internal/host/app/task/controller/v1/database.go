package v1

import (
	"mulberry/internal/host/app/common/controller"
	"mulberry/internal/host/app/task/logic"
	"mulberry/internal/host/app/task/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateDatabaseInfo doc
//
//	@tags			数据库信息管理
//	@Summary		添加数据库信息信息
//	@Description	添加数据库信息信息
//	@Produce		json
//	@Param			body	body		requests.CreateDatabase	true	"body"
//	@Success		200		{object}	controller.Response		"code: 200 成功"
//	@Failure		500		{object}	controller.Response		"错误返回内容"
//	@Router			/api/v1/task/create [post]
//	@Security		ApiKeyAuth
func CreateDatabaseInfo(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateDatabase()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateDatabaseInfo(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateDatabaseInfo doc
//
//	@tags			数据库信息管理
//	@Summary		修改数据库信息信息
//	@Description	修改数据库信息信息
//	@Produce		json
//	@Param			body	body		requests.UpdateDatabaseInfo	true	"body"
//	@Success		200		{object}	controller.Response			"code: 200 成功"
//	@Failure		500		{object}	controller.Response			"错误返回内容"
//	@Router			/api/v1/task/update [put]
//	@Security		ApiKeyAuth
func UpdateDatabaseInfo(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateDatabaseInfo()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateDatabaseInfo(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteDatabaseInfo doc
//
//	@tags			数据库信息管理
//	@Summary		删除数据库信息信息
//	@Description	删除数据库信息信息
//	@Produce		json
//	@Param			id	path		int					true	"数据库信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/delete [delete]
//	@Security		ApiKeyAuth
func DeleteDatabaseInfo(c *gin.Context) {
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

	err = logic.DeleteDatabaseInfo(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DatabaseInfoList doc
//
//	@tags			数据库信息管理
//	@Summary		数据库信息列表
//	@Description	数据库信息列表
//	@Produce		json
//	@Param			body	body		requests.QueryDatabaseInfo	true	"body"
//	@Success		200		{object}	controller.Response			"code: 200 成功"
//	@Failure		500		{object}	controller.Response			"错误返回内容"
//	@Router			/api/v1/task/list [post]
//	@Security		ApiKeyAuth
func DatabaseInfoList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryDatabaseInfo()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.DatabaseInfoList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetDatabaseInfo doc
//
//	@tags			数据库信息管理
//	@Summary		数据库信息详情
//	@Description	数据库信息详情
//	@Produce		json
//	@Param			id	path		int					true	"数据库信息id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/get [get]
//	@Security		ApiKeyAuth
func GetDatabaseInfo(c *gin.Context) {
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

	info, err := logic.GetDatabaseInfo(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}
