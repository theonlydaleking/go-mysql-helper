package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	chalk "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println()

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connected succesfully")

	getUsers(db)
	username := getUsername()
	password := getPassword()
	DbName := getDbName()
	fmt.Println()
	fmt.Println(chalk.Green("Your username will be: "), chalk.Cyan(username))
	fmt.Println(chalk.Green("Your db Name will be: "), chalk.Cyan(password))
	fmt.Println(chalk.Green("Your db Name will be: "), chalk.Cyan(DbName))

	// check the users
	// check the databases
	// Get some inputs

	defer db.Close()

}

func getUsers(db *sql.DB) {

	sqlStatement := `SELECT User FROM mysql.user;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		users := Users{}
		// up to here with https://flaviocopes.com/golang-sql-database/
	}
	fmt.Println(rows)
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

// func createDatabase() {

// }
