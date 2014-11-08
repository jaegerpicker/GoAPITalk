package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "fmt"
    )

type User struct {
    Id              int64
    Username        string              `sql:not null;"size:50"`
    Tasks           []Task              `json:"Todos"`
}

type Task struct {
    Id              int64
    Todo            string              `sql:not null;"size:255"`
    UserId          int64
}

func GetUsers(db *gorm.DB) (users []User, err error) {
    var tasks []Task
    err = db.Find(&users).Error
    if err != nil && fmt.Sprintf("%v",err) != "Record Not Found" {
        LogWrite(fmt.Sprintf("Error trying to get all users %v", err), "ERROR")
        return users, err
    }
    for id, user := range users {
        err = db.Model(&user).Related(&tasks).Error
        if err != nil {
            LogWrite(fmt.Sprintf("%v", err), "ERROR")
        }
        user.Tasks = tasks
        //LogWrite(fmt.Sprintf("%v", user), "ERROR")
        users[id] = user
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
    var tasks []Task
    err = db.Where("id = ?", id).First(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find user: %v, err: %v", id, err), "ERROR")
        return user, err
    }
    err = db.Model(&user).Related(&tasks).Error
    if err != nil {
        LogWrite(fmt.Sprintf("%v", err), "ERROR")
    }
    user.Tasks = tasks
    return user, nil
}

func GetTodosForUser(db *gorm.DB, id int64) (todos []Task, err error) {
    var user User
    err = db.Where("id = ?", id).First(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find user: %v, err: %v", id, err), "ERROR")
        return todos, err
    }
    return user.Tasks, nil
}

func UpdateUser(db *gorm.DB, id int64, username string) (user User, err error) {
    err = db.Where("id = ?", id).First(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find user: %v, err: %v", id, err), "ERROR")
        return user, err
    }
    user.Username = username
    err = db.Save(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to save User: %s err: %v", username, err), "ERROR")
        return user, err
    }
    return user, nil
}

func AddTodoToUser(db *gorm.DB, id int64, tid int64) (user User, err error) {
    err = db.Where("id = ?", id).First(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find user: %v, err: %v", id, err), "ERROR")
        return user, err
    }
    var t Task
    err = db.Where("id = ?", tid).First(&t).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to find task %v, err: %v", tid, err), "ERROR")
        return user, err
    }
    user.Tasks = append(user.Tasks, t)
    err = db.Save(&user).Error
    if err != nil {
        LogWrite(fmt.Sprintf("Error trying to save user %v, err: %v", user, err), "ERROR")
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

func AddTodo(db *gorm.DB, task string, uid int64) (todo Task, err error) {
    todo.Todo = task
    todo.UserId = uid
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
