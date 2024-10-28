package main

import (
    "fmt"
    "net/http"
    "strings"
)

const (
    MonitorsRoutePath = "/api/category/monitors"
    MonitorBaseRoutePath  = "/api/category/monitor/"
    MonitorClickRouterPath = "/api/category/monitor_click/"
)

var getAllMonitorsHandler = createHandler(http.MethodGet, func(w http.ResponseWriter, request *http.Request) {
    var monitorLines []string
    monitors := getAllMonitors()

    for _, monitor := range monitors {
        monitorLines = append(monitorLines, fmt.Sprintf("[%d, %s]", monitor.ID, monitor.Name))
    }

    fmt.Fprintf(w, "{ \"monitors\": [%s] }", strings.Join(monitorLines, ", "))
})

var getMonitorStatsHandler = createHandler(http.MethodGet, func(w http.ResponseWriter, request *http.Request) {
    monitorId := strings.TrimPrefix(request.URL.Path, MonitorBaseRoutePath)
    countValue := getMonitorClickCount(monitorId)
    fmt.Fprintf(w, "{ \"id\": %s, \"count\": %d }", monitorId, countValue)
})

var updateMonitorClickCountHandler= createHandler(http.MethodPost, func(w http.ResponseWriter, request *http.Request) {
    monitorId := strings.TrimPrefix(request.URL.Path, MonitorClickRouterPath)
    updateMonitorClickCount(monitorId)
})

func createHandler(expectedMethod string, handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, request *http.Request) {
        if err := validateRequest(w, request, expectedMethod); err == nil {
            w.Header().Set("Access-Control-Allow-Origin", "*") 
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type") 
            w.Header().Set("Content-Type", "application/json")
            if request.Method == http.MethodOptions {
                w.WriteHeader(http.StatusOK)
                return
            }
            handler(w, request) 
        }
    }
}

func validateRequest(w http.ResponseWriter, request *http.Request, expectedMethod string) error {
    if request.Method != expectedMethod {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return fmt.Errorf("Method Not Allowed: " + request.Method)
    }
    return nil  
}
