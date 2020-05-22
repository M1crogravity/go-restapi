package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

//Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

//New ...
func New(c *Config) *Store {
	return &Store{
		config: c,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}
