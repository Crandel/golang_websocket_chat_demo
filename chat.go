package main

import (
	"app/controllers"
	"app/utils/db"
	"app/utils/server"
	"app/utils/session"
	"app/utils/settings"
	"encoding/json"
	"log"
	_ "net/http/pprof"
	"runtime"
)

var config = &configuration{}

// config struct
type configuration struct {
	Version  string                `json:"Version"`
	Database db.Database           `json:"Database"`
	Server   server.Server         `json:"Server"`
	Template controllers.Templates `json:"Template"`
	Session  session.Session       `json:"Session"`
}

// ParseJSON ...
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	settings.LoadConfig("config.json", config)
	db.LoadDb(&config.Database)
	controllers.LoadTemplates(&config.Template)
	session.InitSession(&config.Session, config.Server.Domain)
	controllers.NewHub()
	r := controllers.RouteInit()
	server.Run(r, config.Server)
}
