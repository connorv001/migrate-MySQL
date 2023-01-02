package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	// "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func readTablesFromCSV(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tables []string
	scanner := csv.NewReader(file)
	for {
		record, err := scanner.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tables = append(tables, record[0])
	}
	return tables, nil
}

func main() {
	// Parse command-line flags
	oldHost := flag.String("old-host", "", "The hostname of the old MySQL server")
	oldUser := flag.String("old-user", "", "The username to use to connect to the old MySQL server")
	oldPassword := flag.String("old-password", "", "The password to use to connect to the old MySQL server")
	oldDBName := flag.String("old-dbname", "", "The name of the database to migrate from the old MySQL server")
	newHost := flag.String("new-host", "", "The hostname of the new MySQL server")
	newUser := flag.String("new-user", "", "The username to use to connect to the new MySQL server")
	newPassword := flag.String("new-password", "", "The password to use to connect to the new MySQL server")
	newDBName := flag.String("new-dbname", "", "The name of the database to migrate to on the new MySQL server")
	tablesFile := flag.String("tables", "", "The path to the CSV file containing the table names to migrate")
	flag.Parse()

	// Validate command-line flags
	if *oldHost == "" || *oldUser == "" || *oldPassword == "" || *oldDBName == "" || *newHost == "" || *newUser == "" || *newPassword == "" || *newDBName == "" || *tablesFile == "" {
		flag.Usage()
		os.Exit(1)
	}
	// Read the table names from the CSV file
	tables, err := readTablesFromCSV(*tablesFile)
	if err != nil {
		log.Fatal(err)
	}
	// Connect to the old MySQL database
	oldDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", *oldUser, *oldPassword, *oldHost, *oldDBName)
	oldDB, err := sql.Open("mysql", oldDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer oldDB.Close()
	// Connect to the new MySQL database
	newDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", *newUser, *newPassword, *newHost, *newDBName)
	newDB, err := sql.Open("mysql", newDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer newDB.Close()
	// Migrate the specified tables
	for _, table := range tables {
		query := fmt.Sprintf("CREATE TABLE %s LIKE %s.%s", table, *oldDBName, table)
		if _, err := newDB.Exec(query); err != nil {
			log.Fatal(err)
		}
		query = fmt.Sprintf("INSERT INTO %s SELECT * FROM %s.%s", table, *oldDBName, table)
		if _, err := newDB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Migration complete!")
}




