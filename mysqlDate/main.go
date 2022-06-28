package main

//connect mysql database
import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //导入database/sql的便捷包，用于查询各种mysql数据库
)

func main() {
	//配置数据库连接信息
	// dsn := "root:123456@tcp(localhost:3306)/sql_test?parseTime=true"
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/sql_test")
	if err != nil {
		log.Println("Verify mysql account failed,err:", err)
		return
	}
	//初始化数据库连接，尝试建立连接
	err = db.Ping()
	if err != nil {
		log.Println("connect mysql failed,err:", err)
	}
	fmt.Println("连接成功")
	{ // Create a new table
		query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Insert a new user
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

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
