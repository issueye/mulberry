package v1

import (
	"mulberry/internal/host/app/common/controller"
	"mulberry/internal/host/app/task/logic"
	"mulberry/internal/host/app/task/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTask doc
//
//	@tags			定时任务管理
//	@Summary		添加定时任务信息
//	@Description	添加定时任务信息
//	@Produce		json
//	@Param			body	body		requests.CreateTask	true	"body"
//	@Success		200		{object}	controller.Response	"code: 200 成功"
//	@Failure		500		{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/create [post]
//	@Security		ApiKeyAuth
func CreateTask(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateTask()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateTask(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateTask doc
//
//	@tags			定时任务管理
//	@Summary		修改定时任务信息
//	@Description	修改定时任务信息
//	@Produce		json
//	@Param			body	body		requests.UpdateTask	true	"body"
//	@Success		200		{object}	controller.Response	"code: 200 成功"
//	@Failure		500		{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/update [put]
//	@Security		ApiKeyAuth
func UpdateTask(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateTask()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateTask(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteTask doc
//
//	@tags			定时任务管理
//	@Summary		删除定时任务信息
//	@Description	删除定时任务信息
//	@Produce		json
//	@Param			id	path		int					true	"定时任务id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/delete [delete]
//	@Security		ApiKeyAuth
func DeleteTask(c *gin.Context) {
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

	err = logic.DeleteTask(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// TaskList doc
//
//	@tags			定时任务管理
//	@Summary		定时任务列表
//	@Description	定时任务列表
//	@Produce		json
//	@Param			body	body		requests.QueryTask	true	"body"
//	@Success		200		{object}	controller.Response	"code: 200 成功"
//	@Failure		500		{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/list [post]
//	@Security		ApiKeyAuth
func TaskList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryTask()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.TaskList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetTask doc
//
//	@tags			定时任务管理
//	@Summary		定时任务详情
//	@Description	定时任务详情
//	@Produce		json
//	@Param			id	path		int					true	"定时任务id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/get [get]
//	@Security		ApiKeyAuth
func GetTask(c *gin.Context) {
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

	info, err := logic.GetTask(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}

// SaveCode doc
//
//	@tags			定时任务管理
//	@Summary		保存脚本代码
//	@Description	保存脚本代码
//	@Produce		json
//	@Param			body	body		requests.SaveCode	true	"body"
//	@Success		200		{object}	controller.Response	"code: 200 成功"
//	@Failure		500		{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/save_code [put]
//	@Security		ApiKeyAuth
func SaveCode(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewSaveCode()
	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.SaveCode(req.ID, req.Code)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateTaskStatus doc
//
//	@tags			定时任务管理
//	@Summary		更新任务状态
//	@Description	更新任务状态
//	@Produce		json
//	@Param			id	path		int					true	"定时任务id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/updateStatus/{id} [put]
//	@Security		ApiKeyAuth
func UpdateTaskStatus(c *gin.Context) {
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

	err = logic.UpdateTaskStatus(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// RunTask doc
//
//	@tags			定时任务管理
//	@Summary		立即执行任务
//	@Description	立即执行任务
//	@Produce		json
//	@Param			id	path		int					true	"定时任务id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response	"错误返回内容"
//	@Router			/api/v1/task/run/{id} [get]
//	@Security		ApiKeyAuth
func RunTask(c *gin.Context) {
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

	err = logic.RunTask(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
