package app

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type App struct {
	Config  Config
	DB      *sql.DB
	Logger  *log.Logger
	Version string
}

func New(config Config, logger *log.Logger, version string) *App {
	return &App{
		Config:  config,
		Logger:  logger,
		Version: version,
	}
}

func (app *App) Init() error {
	db, err := app.openDB()
	if err != nil {
		return err
	}
	app.DB = db
	return nil
}

func (app *App) Uninit() {
	app.Logger.Print("Releasing db connections ...")
	app.DB.Close()
}

func (app *App) openDB() (*sql.DB, error) {

	// Create an empty connection pool.
	db, err := sql.Open("postgres", app.Config.Db.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(app.Config.Db.MaxOpenConns)
	db.SetMaxIdleConns(app.Config.Db.MaxIdleConns)
	duration, err := time.ParseDuration(app.Config.Db.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Establishing a database connection. If it couldn't be established
	// successfully within the 5 second deadline (as per provided context)
	// it will return an error.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	app.Logger.Println("Database connection established")
	return db, nil
}
