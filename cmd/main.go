package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/schlucht/mnrNaters/internal/driver"
	"github.com/schlucht/mnrNaters/internal/models"
)

var session *scs.SessionManager

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

const (
	MsgTypeSuccess = iota
	MsgTypeError
	MsgTypeWarning
	MsgTypeInfo
)

type message struct {
	message string
	msgType int
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorlog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	Session       *scs.SessionManager
	DB            models.DBModel
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Printf("Server run on Port: %v in mode: %s\r", app.config.port, app.config.env)
	return srv.ListenAndServe()
}

func main() {

	var cfg config
	tc := make(map[string]*template.Template)

	flag.IntVar(&cfg.port, "port", 4444, "Server to port listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application enviroment { develompen | production}")
	flag.StringVar(&cfg.db.dsn, "dsn", "schmidschluch7:3903Schlucht@tcp(db8.hostpark.net)/schmidschluch7?parseTime=true", "DB connect String")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session := scs.New()
	session.Lifetime = 24 * time.Hour

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Println(err)
		errorLog.Fatal(err)
	}

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorlog:      errorLog,
		templateCache: tc,
		Session:       session,
		version:       "0.0.0.1",
		DB:            models.DBModel{DB: conn},
	}

	err = app.serve()
	if err != nil {
		app.errorlog.Println(err)
		log.Fatal(err)
	}
}
