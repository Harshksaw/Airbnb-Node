package db
import (
	"database/sql"
)

type Storage struct { //facilitate dependency injection
	
	UserRepository UserRepository

}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		UserRepository: &UserRepositoryImpl{},
	}
}




