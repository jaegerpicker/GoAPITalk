package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "fmt"
    )

type User struct {
    Id              int64
    Username        string              `sql:not null;"size:50"`
    Tasks           []Task
}

type Task struct {
    Id              int64
    Todo            string              `sql:not null;"size:255"`
}

func GetUsers(db *gorm.DB) (users []User, err error) {
    err = db.Find(&users).Error
    if err != nil && fmt.Sprintf("%v",err) != "Record Not Found" {
        LogWrite(fmt.Sprintf("Error trying to get all users %v", err), "ERROR")
        return users, err
    }
    return users, nil
}

func AddUser(db *gorm.DB, username string) (user User, err error) {
    user.Username = username
    db.Create(&user)
    err = db.Save(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to add user %s, err %v", username, err), "ERROR")
        return user, err
    }
    return user, nil
}

func GetUser(db *gorm.DB, id int64) (user User, err error) {
    err = db.Where("id = ?", id).First(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find user: %v, err: %v", id, err), "ERROR")
        return user, err
    }
    return user, nil
}

func UpdateUser(db *gorm.DB, id int64, username string, todos []Task) (user User, err error) {
    err = db.Where("id = ?", id).First(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find user: %v, err: %v", id, err), "ERROR")
        return user, err
    }
    user.Username = username
    for _, tp := range todos {
        var t Task
        err := db.Where("id = ?", tp.Id).First(&t).Error
        if err != nil {
            LogWrite(fmt.Sprintf("Error trying to find todo: %s, err: %v", string(tp.Id), err), "ERROR")
            return user, err
        }
        user.Tasks = append(user.Tasks, t)
    }
    err = db.Save(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to save User: %s err: %v", username, err), "ERROR")
        return user, err
    }
    return user, nil
}

func DeleteUser(db *gorm.DB, id int64) (success bool, err error) {
    var user User
    err = db.Where("id = ?", id).First(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error, trying to find user: %v, err: %v", id, err), "ERROR")
        return false, err
    }
    err = db.Delete(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to delete user: %v, err: %v", id, err), "ERROR")
        return false, err
    }
    return true, nil
}

func GetTodos(db *gorm.DB) (todos []Task, err error) {
    err = db.Find(&todos).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to get all todos %v", err), "ERROR")
        return todos, err
    }
    return todos, nil
}

func AddTodo(db *gorm.DB, task string) (todo Task, err error) {
    todo.Todo = task
    db.Create(&todo)
    err = db.Save(&todo).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to add user %s, err %v", task, err), "ERROR")
        return todo, err
    }
    return todo, nil
}

func GetTodo(db *gorm.DB, id int64) (todo Task, err error) {
    err = db.Where("id = ?", id).First(&todo).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find todo: %v, err: %v", id, err), "ERROR")
        return todo, err
    }
    return todo, nil
}

func UpdateTodo(db *gorm.DB, id int64, task string) (todo Task, err error) {
    err = db.Where("id = ?", id).First(&todo).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find todo: %v, err: %v", id, err), "ERROR")
        return todo, err
    }
    todo.Todo = task
    err = db.Save(&todo).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to update todo: %v, error: %v", id, task), "ERROR")
        return todo, err
    }
    return todo, nil
}

func DeleteTodo(db *gorm.DB, id int64) (success bool, err error) {
    var todo Task
    err = db.Where("id = ?", id).First(&todo).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error, trying to find todo: %v, err: %v", id, err), "ERROR")
        return false, err
    }
    err = db.Delete(&todo).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to delete todo: %v, err: %v", id, err), "ERROR")
        return false, err
    }
    return true, nil
}
