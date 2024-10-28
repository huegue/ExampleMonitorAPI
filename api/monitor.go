package main
import (
    "strconv"
)

type Monitor struct {
    ID   int    
    Name string
}

func getAllMonitors() []Monitor {
    var monitors []Monitor
    var monitorID int
    var monitorName string

    DB := openDB()
    rows, _ := DB.Query("select id, name from monitors")
    defer rows.Close()

    for rows.Next() {
        rows.Scan(&monitorID, &monitorName)
        monitors = append(monitors, Monitor{ID: monitorID, Name: monitorName})
    }
    return monitors
}

func getMonitorClickCount(monitorId string) int {
    DB := openDB()
    rows, _ := DB.Query("select count from monitors where id=" + monitorId)
    defer rows.Close()
    
    countValue := 0
    for rows.Next() {
        rows.Scan(&countValue)
    }
    return countValue
}

func updateMonitorClickCount(monitorId string) {
    DB := openDB()
    rows, _ := DB.Query("select count from monitors where id=" + monitorId)
    defer rows.Close()

    countValue := 0
    for rows.Next() {
        rows.Scan(&countValue)
    }
    DB.Exec("update monitors set count=" + strconv.Itoa(countValue + 1) + " where id=" + monitorId)
}