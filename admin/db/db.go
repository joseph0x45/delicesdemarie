package db

import (
	"admin/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const dbSchema = `
  create table if not exists users (
    id integer not null primary key autoincrement,
    username text not null unique,
    password text not null
  );

  insert or ignore into users (username, password)
  values (
    'admin', '$2a$12$zRjZjn3jP1RhIBPIePwQv.n5EH0CZlczVe92BdOPsiQ/RAk7M1pMC'
  );

  create table if not exists sessions (
    id text not null primary key
  );
`

type Conn struct {
	db *sqlx.DB
}

func NewConn(dbPath string, refreshDB bool) (*Conn, error) {
	if refreshDB {
		os.Remove(dbPath)
	}
	db, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to database: %w", err)
	}
	_, err = db.Exec("PRAGMA foreign_keys = on;")
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Error while connecting to database: %w", err)
	}
	_, err = db.Exec(dbSchema)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Error while connecting to database: %w", err)
	}
	log.Println("Connected to database at ", dbPath)
	return &Conn{
		db: db,
	}, nil
}

func (c *Conn) GetUserByUsername(username string) (*models.User, error) {
	const query = `
    select * from users where username=$1
  `
	user := &models.User{}
	err := c.db.Get(user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error while getting user by username: %w", err)
	}
	return user, nil
}

func (c *Conn) InsertSession(id string) error {
	const query = `
    insert into sessions(
      id
    )
    values (
      $1
    )
  `
	_, err := c.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error while inserting new session: %w", err)
	}
	return nil
}

func (c *Conn) GetSessionByID(id string) (*models.Session, error) {
	const query = `
    select * from sessions where id=$1
  `
	session := &models.Session{}
	err := c.db.Get(session, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error while getting session: %w", err)
	}
	return session, nil
}
