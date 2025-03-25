package models

import "project/dao"

// TODO Module
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreatTodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

func FindAllTodo() (todo []*Todo, err error) {
	err = dao.DB.Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// 这里你声明了 todo *Todo，
// 但是并没有初始化 todo，所以 todo 是 nil。当 First(todo) 试图把查询结果填充到 nil 指针时，
// 程序会 panic 或者报 unsupported destination, should be slice or struct。
func GetATodoByID(id string) (todo *Todo, err error) {
	todo = new(Todo) // ✅ 初始化 todo，避免 nil 指针错误
	err = dao.DB.Where("id = ?", id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateTodoByID(id string) (err error) {
	todo, err := GetATodoByID(id)
	if err != nil {
		return err
	}
	// 取反status
	todo.Status = !todo.Status
	err = dao.DB.Save(todo).Error
	return err
}

func DeleteTodoByID(id string) (err error) {
	todo, err := GetATodoByID(id)
	if err != nil {
		return err
	}
	err = dao.DB.Delete(todo).Error
	return err
}
