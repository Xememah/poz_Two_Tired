package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"os"

	"encoding/json"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/converter"
	"github.com/AllegroTechDays/poz_Two_Tired/backend/store"
)

type App struct {
	Store  store.Store
	Logger *log.Logger
}

func (app *App) Listen() {
	server := http.Server{
		Addr:         ":3000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	http.HandleFunc("/data", app.serveData)
	if err := server.ListenAndServe(); err != nil {
		app.Logger.Fatal(err.Error())
	}
}

func main() {
	logger := log.New(os.Stdout, "TwoTired", log.Ldate|log.Lshortfile)
	data, err := converter.Load("./_data")
	if err != nil {
		logger.Fatal(err.Error())
	}
	app := &App{
		Store:  &store.Local{Data: data},
		Logger: logger,
	}
	app.Listen()
}

func (app *App) serveData(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	ts, err := strconv.ParseInt(query.Get("timestamp"), 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	duration, err := strconv.Atoi(query.Get("duration"))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	encoder := json.NewEncoder(rw)
	data, err := app.Store.TimestampNarrowed(time.Unix(ts, 0), time.Duration(duration)*time.Second)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := encoder.Encode(data); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
