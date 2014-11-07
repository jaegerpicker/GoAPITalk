package main

import (
	"fmt"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
    "net/http"
    "regexp"
    "time"
    "sync"
    "strings"
    )

var settings Settings

func main() {
    m := martini.Classic()
    var s Settings
    var err error
	settings, err = s.GetSettings()
	if err != nil {
		// LogWrite is my version of django's log.info etc... Works about the same way "DEBUG", "INFO", "WARN", and
		// "ERROR" are valid choices
		LogWrite(fmt.Sprintf("GetSettings failed err: %v", err), "ERROR")
		return
	}
    db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", settings.DbUser, settings.DbPass, settings.DbHost, settings.DbPort, settings.DbName))
	if err != nil {
		LogWrite(fmt.Sprintf("Failed to open Db: %v", err), "ERROR")
	}
    AutoMigrateModels(&db)
    m.Use(martini.Recovery())
	m.Use(martini.Logger())
    m.Use(MapEncoder)
    m.Map(&db)
    db.LogMode(true)
    r := martini.NewRouter()
    r.Get("/version/", func() string {
		return "alive"
	})
	r.Get("/users", UsersList)
    r.Post("/users", UsersAdd)
    r.Get("/users/:id", UsersShow)
    r.Put("/users/:id", UsersUpdate)
    r.Delete("/users/:id", UsersDelete)
    r.Get("/todos", GetTodos)
	r.Post("/todos", TodosAdd)
    r.Get("/todos/:id", TodosShow)
    r.Put("/todos/:id", TodosUpdate)
    r.Delete("/todos/:id", TodosDelete)

	m.Action(r.Handle)
    // go uses channels to communitcate amgonst async go routines
    cs := make(chan bool)
    var wg sync.WaitGroup
    wg.Add(1)
    go func(cs chan bool) {
        defer wg.Done()
        b := <-cs
        fmt.Println(fmt.Sprintf("Channel passed: %v", b))
        if err = http.ListenAndServe(":"+settings.Port, m); err != nil {
            LogWrite(fmt.Sprintf("Error starting webserver: %v", err), "ERROR")
        }

    }(cs)
    wg.Add(1)
    go CheckOnFile(&wg, cs)

    // You can use this to wait on the go routines to exit, in the http server case
    // it hopefully never exits.
    wg.Wait()

    //for {
    //    duration := time.Second
    //    time.Sleep(duration * 5)
    //    fmt.Println("Loop to keep application alive, sleeping: %v", (duration * 5))
    //}

}

func CheckOnFile(wg *sync.WaitGroup, cs chan bool) {
    defer wg.Done()
    fmt.Println("This is just a example function. In here you could easliy find a file and use a channel to pass the contents to another goroutine")
    cs <- true
    for {
        duration := time.Second
        time.Sleep(duration * 5)
        fmt.Println("Loop checking on a file....")
    }

}

var rxExt = regexp.MustCompile(`(\.(?:xml|text|json))\/?$`)


// Function: MapEncoder is a martini middleware func to match the request type to the proper encoding
func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	// Get the format extension
	matches := rxExt.FindStringSubmatch(r.URL.Path)
	ft := ".json"
	if len(matches) > 1 {
		// Rewrite the URL without the format extension
		l := len(r.URL.Path) - len(matches[1])
		if strings.HasSuffix(r.URL.Path, "/") {
			l--
		}
		r.URL.Path = r.URL.Path[:l]
		ft = matches[1]
	}
	// Inject the requested encoder
	LogWrite(fmt.Sprintf("\nencoder type:\n\n%s\n\n", ft), "INFO")
	switch ft {
	case ".xml":
		c.MapTo(XmlEncoder{}, (*Encoder)(nil))
		w.Header().Set("Content-Type", "text/xml")
	case ".text":
		c.MapTo(TextEncoder{}, (*Encoder)(nil))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	case ".json":
		c.MapTo(JsonEncoder{}, (*Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	default:
		c.MapTo(JsonEncoder{}, (*Encoder)(nil))
        w.Header().Set("Content-Type", "application/json")
	}
}
