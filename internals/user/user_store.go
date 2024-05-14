package user

import (
	"context"
	"database/sql"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{
		db,
	}
}

func (s *Store) createUser(ctx context.Context, user *User) (*User,error){
	var userId int
	query := `INSERT INTO users(username,email,password) VALUES ($1, $2, $3) RETURNING id`

	err := s.db.QueryRowContext(ctx,query,user.Username,user.Email,user.Password).Scan(&userId)

	if err != nil {
		log.Fatal("something wrong with the query")
		return nil,err
	}
	user.ID = int64(userId)
	return user,nil
}

func (s *Store) getUserByEmail(ctx context.Context,email string) (*User,error) {
	u :=User{}

	query:=`SELECT id,username,email,password FROM users WHERE email = $1`

	err := s.db.QueryRowContext(ctx,query,email).Scan(&u.ID,&u.Username,&u.Email,&u.Password)

	if err != nil {
		log.Fatal("something wrong with the query")
		return nil,err
	}
	return &u,nil
}