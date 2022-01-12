package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sonerrtng/go-social-media/entity"
)

type MysqlRepository struct {
	DatabasePointer *sql.DB
}

func CreateMysqlRepository() MysqlRepository {
	dbPointer, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic("Error connect not the mysql database")
	}

	return MysqlRepository{DatabasePointer: dbPointer}
}

func (repository MysqlRepository) CreateUser(user entity.User) int64 {
	statment, err := repository.DatabasePointer.Prepare("INSERT INTO users(user_name,name,sur_name,password,age) VALUES" +
		"(?,?,?,?,?)")
	if err != nil {
		panic(err)
	}

	runStatment, err := statment.Exec(user.UserName, user.Name, user.SurName, user.Password, user.Age)
	if err != nil {
		panic(err)
	}

	response, err := runStatment.RowsAffected()

	if err != nil {
		panic("Error get RowsAffected methods")
	}

	return response
}

func (repository MysqlRepository) GetUser() []entity.User {
	rows, err := repository.DatabasePointer.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	results := []entity.User{}

	for rows.Next() {
		var userPointer entity.User
		err := rows.Scan(&userPointer.UserName, &userPointer.Name, &userPointer.SurName, &userPointer.Password, &userPointer.Age)
		if err != nil {
			panic(err)
		}
		results = append(results, userPointer)
	}
	return results
}
