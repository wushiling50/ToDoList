package service

import (
	"main/ToDoList/model"
	"main/ToDoList/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` // 0 是未做, 1是已做
}

type ShowTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type FinishTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UnfinishTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Status int `json:"status" form:"status"` // 0 是未做, 1是已做
}

type ChangeTaskService struct {
	Status int `json:"status" form:"status"` // 0 是未做, 1是已做
}

type DeleteTaskService struct {
}

type DeleteFinishTaskService struct {
}

type DeleteUnfinishTaskService struct {
}

type DeleteAllTaskService struct {
}

// 新增一条备忘录
func (service *CreateTaskService) Create(uid uint) serializer.Response {
	var user model.User
	code := 200
	model.DB.First(&user, uid)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
	}
}

// 展示一条备忘录
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}

}

// 列表返回所有备忘录
func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

// 列表返回已完成事项
func (service *FinishTaskService) FindFinish(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("status = ?", 1).
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "查询失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

// 列表返回未完成事项
func (service *UnfinishTaskService) FindUnfinish(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("status = ?", 0).
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "查询失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

// 查询
func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}

	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "查询失败",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// 更新
func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Status = service.Status
	err := model.DB.Save(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    " 更新备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新成功",
	}
}

// 更新一批
func (service *ChangeTaskService) Change(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.Status == 0 {
		model.DB.Model(&model.Task{}).Where("uid=?", uid).Where("status = ?", 1).Update("status", 0)
	} else if service.Status == 1 {
		model.DB.Model(&model.Task{}).Where("uid=?", uid).Where("status = ?", 0).Update("status", 1)
	}
	if service.Status == 0 {
		err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("status = ?", 0).
			Count(&count).Find(&tasks).Error
		if err != nil {
			return serializer.Response{
				Status: 400,
				Msg:    "查询失败",
				Error:  err.Error(),
			}
		}
	} else if service.Status == 1 {
		err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("status = ?", 1).
			Count(&count).Find(&tasks).Error
		if err != nil {
			return serializer.Response{
				Status: 400,
				Msg:    "查询失败",
				Error:  err.Error(),
			}
		}
	}
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

// 删除一条
func (service *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task

	err := model.DB.Delete(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

// 删除所有已经完成
func (service *DeleteFinishTaskService) DeleteFinish(uid uint) serializer.Response {
	var task model.Task
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).
		Where("status = ?", 1).Delete(&task).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

// 删除所有代办事项
func (service *DeleteUnfinishTaskService) DeleteUnfinish(uid uint) serializer.Response {
	var task model.Task
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).
		Where("status = ?", 0).Delete(&task).Error

	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

// 删除全部数据
func (service *DeleteAllTaskService) DeleteAll(uid uint) serializer.Response {
	var task model.Task
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Delete(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,

		Msg: "删除成功",
	}
}
