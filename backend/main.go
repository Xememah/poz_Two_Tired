package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"os"

	"encoding/json"

	"github.com/AllegroTechDays/poz_Two_Tired/backend/model"
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
	app := &App{
		Store:  &store.Local{},
		Logger: logger,
	}
	err := app.Store.Init()
	if err != nil {
		logger.Fatalln(err.Error())
	}
	app.Listen()
}

func (app *App) serveData(rw http.ResponseWriter, req *http.Request) {
	options := store.Query{}
	query := req.URL.Query()
	if query.Get("timestamp") != "" {
		ts, err := strconv.ParseInt(query.Get("timestamp"), 10, 64)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		timestamp := time.Unix(ts, 0)
		options.Timestamp = &timestamp
		duration, err := strconv.Atoi(query.Get("duration"))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		options.Offset = time.Duration(duration) * time.Second
	}

	if query.Get("latitude") != "" && query.Get("longitude") != "" {
		lat, err := strconv.ParseFloat(query.Get("latitude"), 64)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		lng, err := strconv.ParseFloat(query.Get("longitude"), 64)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		dist, err := strconv.ParseFloat(query.Get("distance"), 32)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		options.Position = &model.Position{
			Latitude:  lat,
			Longitude: lng,
		}
		options.Distance = dist
	}
	encoder := json.NewEncoder(rw)
	data, err := app.Store.Narrowed(options)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := encoder.Encode(data); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
