package main
import (
    "fmt"
    "os"
    "log"
    "net/http"
)

func runHelpCommand() {
    fmt.Println("\nHelp::\t\t\t\t.counter --help")
    fmt.Println("Create products database:\t.counter --createdb")
    fmt.Println("Start server:\t\t\t.counter --start\n")
}

func runCreateDBCommand() {
	if _, err := os.Stat("./monitors.txt"); os.IsNotExist(err) {
		log.Println("ERROR! File \"monitors.txt\" does not exist!")
		return
	}

	if _, err := os.Stat("./products.db"); err == nil {
		if os.Remove("./products.db") != nil {
			log.Println(err)
			return
		}
	}

	setupDatabaseAndImportMonitors("./monitors.txt")
	fmt.Println("OK. File products.db is created!")
}


func runStartCommand() {
	port := ":8090"

	http.HandleFunc(MonitorsRoutePath, getAllMonitorsHandler)
	http.HandleFunc(MonitorBaseRoutePath, getMonitorStatsHandler)
	http.HandleFunc(MonitorClickRouterPath, updateMonitorClickCountHandler)

	fmt.Println("The server is running at port " + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Failed to start server!", err)
	}
}
