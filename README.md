# migrate-MySQL

migrate-MySQL is a command-line tool that allows you to migrate a large MySQL
database to a new host. It reads the table names from a CSV file and then
migrates each table by creating a new table with the same structure as the
old table, and then inserting all of the data from the old table into the
new table.

To use the tool, you will need to provide the following command-line flags:
```
- old-host: The hostname of the old MySQL server
 - old-user: The username to use to connect to the old MySQL server
 - old-password: The password to use to connect to the old MySQL server
 - old-dbname: The name of the database to migrate from the old MySQL server
 - new-host: The hostname of the new MySQL server
 - new-user: The username to use to connect to the new MySQL server
 - new-password: The password to use to connect to the new MySQL server
 - new-dbname: The name of the database to migrate to on the new MySQL server
 - tables: The path to the CSV file containing the table names to migrate
 ```
 
 
 
 ## Development
 
 To install the required packages, you will need to have the Go toolchain
installed on your system. Once you have done so, you can use the following
commands to install the required packages:
 go get -u github.com/golang-migrate/migrate/v4
 go get -u github.com/go-sql-driver/mysql
