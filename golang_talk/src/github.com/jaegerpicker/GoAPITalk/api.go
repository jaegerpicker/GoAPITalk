package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"net/http"
    "strconv"
)


func UsersList(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    users, err := GetUsers(db)
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    if len(users) <= 0 {
        return http.StatusNotFound, "No users in db"
    }
    ret, _ := enc.Encode(users)
    return http.StatusOK, ret
}

func UsersAdd(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params, todos []Task) (int, string) {
    user, err := AddUser(db, params["username"])
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(user)
    return http.StatusOK, ret
}

func UsersShow(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    id, err := strconv.Atoi(params["user_id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid user id"
    }
    user, err := GetUser(db, int64(id))
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(user)
    return http.StatusOK, ret
}

func UsersUpdate(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params, todos []Task) (int, string) {
    id, err := strconv.Atoi(params["user_id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid user id"
    }
    user, err := UpdateUser(db, int64(id), params["username"], todos)
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(user)
    return http.StatusOK, ret
}

func UsersDelete(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid user id"
    }
    ret, err := DeleteUser(db, int64(id))
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    if !ret {
        return http.StatusConflict, "User not deleted!"
    }
    return http.StatusOK, "user deleted"
}

func TodosList(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    todos, err := GetTodos(db)
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(todos)
    return http.StatusOK, ret
}

func TodosAdd(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    todo, err := AddTodo(db, params["task"])
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(todo)
    return http.StatusOK, ret
}

func TodosShow(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid id"
    }
    todo, err := GetTodo(db, int64(id))
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(todo)
    return http.StatusOK, ret
}

func TodosUpdate(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid id"
    }
    todo, err := UpdateTodo(db, int64(id), params["task"])
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(todo)
    return http.StatusOK, ret
}

func TodosDelete(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    ret, err := DeleteTodo(db, params["todo_id"])
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    if !ret {
        return http.StatusConflict, "Todo not deleted!"
    }
    return http.StatusAccepted, "Todo Deleted!"
}
