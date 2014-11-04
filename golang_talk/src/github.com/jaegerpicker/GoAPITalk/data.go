package main

type User struct {
    Username        string
    Tasks           []Task
}

type Task struct {
    Title           string
}
