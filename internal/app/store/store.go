package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
	"github.com/sirupsen/logrus"
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
	logger *logrus.Logger
}

// New ...
func New(config *Config, logger *logrus.Logger) *Store {
	return &Store{
		config: config,
		logger: logger,
	}
}

// Open ...
func (s *Store) Open() error {
	s.logger.Debug("Open db method")

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

// Close ...
func (s *Store) Close() {
	s.logger.Debug("Close db method")
	s.db.Close()
}
