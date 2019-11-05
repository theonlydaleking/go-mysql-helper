package main

// Using https://flaviocopes.com/golang-sql-database/

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	chalk "github.com/logrusorgru/aurora"
)

type createDetails struct {
	username string
	password string
	DbName   string
}

func main() {
	fmt.Println()

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connected succesfully")

	getUsers(db)
	// ToDo: Get databases

	username := getUsername()
	password := getPassword()
	DbName := getDbName()
	fmt.Println()
	fmt.Println(chalk.Green("Your username will be: "), chalk.Cyan(username))
	fmt.Println(chalk.Green("Your db Name will be: "), chalk.Cyan(password))
	fmt.Println(chalk.Green("Your db Name will be: "), chalk.Cyan(DbName))

	details := createDetails{username: username, password: password, DbName: DbName}

	createDatabase(details, db)

	// check the users
	// check the databases
	// Get some inputs

	fmt.Printf("\nSuccessfully created database: %v and user %v", username, password)

	defer db.Close()

}

func createDatabase(details createDetails, db *sql.DB) {
	sqlStatement := fmt.Sprintf("CREATE DATABASE %v;", details.DbName)
	execDbQuery(sqlStatement, db)

	sqlStatement = fmt.Sprintf("CREATE USER '%v'@'localhost' IDENTIFIED BY '%v';", details.username, details.password)
	execDbQuery(sqlStatement, db)

	sqlStatement = fmt.Sprintf("GRANT ALL PRIVILEGES ON %v.* TO '%v'@'localhost';", details.DbName, details.username)
	execDbQuery(sqlStatement, db)

}

func execDbQuery(query string, db *sql.DB) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func getUsers(db *sql.DB) {

	type User struct {
		username string
	}

	sqlStatement := `SELECT User FROM mysql.user;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	fmt.Println(chalk.Green("\nCurrent MySQL Users:"))
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.username)
		if err != nil {
			panic(err)
		}
		fmt.Println(user.username)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func inputHandler(askFor string) string {

	buf := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a %s: ", askFor)
	answer, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}
	answerString := string(answer)
	answerString = strings.TrimSuffix(answerString, "\n")
	return answerString
}

// Asks user for a Username
func getUsername() string {

	username := inputHandler("username")
	return username
}

// Asks user for a Password
func getPassword() string {

	password := inputHandler("password")
	return password
}

// Asks user for the name of the database
func getDbName() string {

	dbName := inputHandler("database name")
	return dbName
}
