package main
import (
    "log"
    "os"
    "bufio"
    "strings"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)


var openDB = createDBConnector()

func createDBConnector() func() *sql.DB {
    var DB *sql.DB

    return func() *sql.DB {
        if DB == nil || DB.Ping() != nil {
            db, err := sql.Open("sqlite3", "./data/products.db")
            if err != nil {
                log.Fatal("Failed to open the database:", err)
            }
            DB = db
        }
        return DB
    }
}

func createMonitorsTable(DB *sql.DB) {
    _, err := DB.Exec("create table monitors(id integer primary key, name varchar(255) not null, count integer)")
    if err != nil {
        log.Fatal(err)
    }   
}

func importMonitorsFromFile(DB *sql.DB, filePath string) {
    var file *os.File
    var err error

    if file, err = os.Open(filePath); err != nil {
        log.Fatal("Failed to open the file: ", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        arr := strings.Split(scanner.Text(), ",")
        id := arr[0]
        monitorName := arr[1]
        DB.Exec("insert into monitors(id, name, count) values($1, $2, 0)", id, monitorName)
    }
}

func setupDatabaseAndImportMonitors(importFilePath string) {
    DB := openDB()
    defer DB.Close()
    createMonitorsTable(DB)
    importMonitorsFromFile(DB,importFilePath)
}