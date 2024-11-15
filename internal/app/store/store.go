package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	UserRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	var err error
	s.db, err = sql.Open("postgres", s.config.DatabaseURl)
	if err != nil {
		return err
	}
	return s.db.Ping()
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.UserRepository == nil {
		s.UserRepository = &UserRepository{store: s}
	}
	return s.UserRepository
}
