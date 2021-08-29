package ProductDb

import(
	"database/sql"
	"fmt"
    "log"
    "time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(){
	db, err := sql.Open("mysql", "root:Sql16180339887@(127.0.0.1:3306)/teste?parseTime=true")

	if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

	// Insert a new user
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
    
}