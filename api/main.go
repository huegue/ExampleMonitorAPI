package main
import (
    "os"
    "strings"
)

func main() {
    if len(os.Args) <= 1 {
        runHelpCommand()
        return
    }

    command := strings.ToLower(os.Args[1])

    switch command {
    case "--help":
        runHelpCommand()
    case "--createdb":
        runCreateDBCommand()
    case "--start":
        runStartCommand()
    default:
        runHelpCommand()
    }
}
