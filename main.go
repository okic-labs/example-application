// I am currently using the main package...
// I wonder how I can import sub packages and create packages...

package main

import "fmt"
import "log"
import "net/http"

func main() {
    helperConnector()
}

func helperConnector() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request ) {
        fmt.Fprintln(writer, "Hello Git!")
    })
    helperLoggerGuy()
}

func helperLoggerGuy() {
    log.Fatalf("error: %s", http.ListenAndServe(":8080", nil))
}
