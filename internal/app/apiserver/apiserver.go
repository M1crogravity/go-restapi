package apiserver

import (
	"database/sql"
	"github.com/M1crogravity/go-resapi/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq" // ...
	"net/http"
)

func Start(config *Config) error {
	db, err := newDb(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	s := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, s)
}

func newDb(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
