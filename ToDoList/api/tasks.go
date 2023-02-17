package api

import (
	"main/ToDoList/pkg/error"
	"main/ToDoList/pkg/utils"
	"main/ToDoList/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	if err := c.ShouldBind(&showTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	}
}

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	}
}

func FinishTask(c *gin.Context) {
	var finishTask service.FinishTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&finishTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := finishTask.FindFinish(claim.Id)
		c.JSON(200, res)
	}
}

func UnfinishTask(c *gin.Context) {
	var unfinishTask service.UnfinishTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&unfinishTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := unfinishTask.FindUnfinish(claim.Id)
		c.JSON(200, res)
	}
}

func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	}
}

func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	if err := c.ShouldBind(&updateTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

func ChangeTask(c *gin.Context) {
	var changeTask service.ChangeTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&changeTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := changeTask.Change(claim.Id)
		c.JSON(200, res)
	}
}

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	if err := c.ShouldBind(&deleteTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	}
}

func DeleteFinishTask(c *gin.Context) {
	var deleteFinishTask service.DeleteFinishTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFinishTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := deleteFinishTask.DeleteFinish(claim.Id)
		c.JSON(200, res)
	}
}

func DeleteUnfinishTask(c *gin.Context) {
	var deleteUnfinishTask service.DeleteUnfinishTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteUnfinishTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := deleteUnfinishTask.DeleteUnfinish(claim.Id)
		c.JSON(200, res)
	}
}

func DeleteAllTask(c *gin.Context) {
	var deleteAllTask service.DeleteAllTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteAllTask); err != nil {
		logging.Error(err)
		c.JSON(400, error.ErrorResponse(err))
	} else {
		res := deleteAllTask.DeleteAll(claim.Id)
		c.JSON(200, res)
	}
}
