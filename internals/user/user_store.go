package user

import (
	"context"
	"database/sql"
	"fmt"
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

	fmt.Print("user detils", user)

	var userId int
	query := `INSERT INTO "users" (username, email, password) VALUES ($1, $2, $3) returning id`

	err := s.db.QueryRowContext(ctx,query, user.Username, user.Email , user.Password).Scan(&userId)


	if err != nil {
		log.Fatal(err.Error())
		return nil,err
	}


	user.ID = int64(userId)
	return user,nil
}

func (s *Store) getUserByEmail(ctx context.Context,email string) (*User,error) {
	u := User{}

	fmt.Print("Email " , email)

	query:=`SELECT * FROM "users" WHERE email = $1`

	err := s.db.QueryRowContext(ctx,query,email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)

	if err != nil {
		log.Fatal(err.Error())
		return nil,err
	}
	return &u,nil
}