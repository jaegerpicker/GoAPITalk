package main

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"net/http"
    "strconv"
    "encoding/json"
    "io/ioutil"
    "fmt"
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

func UsersAdd(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    defer r.Body.Close()
    var dat map[string]interface{}
    requestBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return http.StatusBadRequest, "Invalid user in request"
    }
    err = json.Unmarshal(requestBody, &dat)
    if err != nil {
        return http.StatusBadRequest, "Invalid user"
    }
    user, err := AddUser(db, fmt.Sprintf("%v", dat["username"]))
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(user)
    return http.StatusOK, ret
}

func UsersShow(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    id, err := strconv.Atoi(params["id"])
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

func UsersUpdate(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    defer r.Body.Close()
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid user id"
    }
    var dat map[string]interface{}
    requestBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return http.StatusBadRequest, "Invalid user in request"
    }
    err = json.Unmarshal(requestBody, &dat)
    if err != nil {
        return http.StatusBadRequest, "Invalid user"
    }
    user, err := UpdateUser(db, int64(id), fmt.Sprintf("%v",dat["username"]))
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    ret, _ := enc.Encode(user)
    return http.StatusOK, ret
}

func UsersAddTodos(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    defer r.Body.Close()
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid user id"
    }
    var dat map[string]interface{}
    requestBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return http.StatusBadRequest, "Invalid todo in request"
    }
    err = json.Unmarshal(requestBody, &dat)
    if err != nil {
        return http.StatusBadRequest, "Invalid todo"
    }
    todo, err := AddTodo(db, fmt.Sprintf("%v",dat["Todo"]), int64(id))
    if err != nil {
        return http.StatusInternalServerError, "Could not create new todo"
    }
    LogWrite(fmt.Sprintf("%v", todo), "INFO")
    user, err := GetUser(db, int64(id))
    if err != nil {
        return http.StatusInternalServerError, "Could not assign todo to user"
    }
    ret, _ := enc.Encode(user)
    return http.StatusOK, ret
}

func UsersShowTodos(r *http.Request, enc Encoder, db *gorm.DB, params martini.Params) (int, string) {
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid user id"
    }
    todos, err := GetTodosForUser(db, int64(id))
    if err != nil {
        return http.StatusBadRequest, "Could not find user"
    }
    ret, _ := enc.Encode(todos)
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
    defer r.Body.Close()
    var dat map[string]interface{}
    requestBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return http.StatusBadRequest, "Invalid todo in request"
    }
    err = json.Unmarshal(requestBody, &dat)
    if err != nil {
        return http.StatusBadRequest, "Invalid todo"
    }
    uid, _ := strconv.Atoi(fmt.Sprintf("%v", dat["UserId"]))
    todo, err := AddTodo(db, fmt.Sprintf("%v",dat["Todo"]), int64(uid))
    if err != nil {
        return http.StatusInternalServerError, "Could not create new todo"
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
    id, err := strconv.Atoi(params["todo_id"])
    if err != nil {
        return http.StatusBadRequest, "Invalid id"
    }
    ret, err := DeleteTodo(db, int64(id))
    if err != nil {
        return http.StatusInternalServerError, ""
    }
    if !ret {
        return http.StatusConflict, "Todo not deleted!"
    }
    return http.StatusAccepted, "Todo Deleted!"
}
