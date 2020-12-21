package main

import (
    "log"
    "net/http"
    "os"
    "strings"
    "flag"
    "fmt"
    "github.com/joho/godotenv"
)

var version = "master"

func showVersion(w http.ResponseWriter, r *http.Request) {
    log.Println(version)
    w.Write([]byte(version))
}

func showEnv(w http.ResponseWriter, r *http.Request){
    godotenv.Load()

    ABC := os.Getenv("ABC")

    log.Println(ABC)
    w.Write([]byte(ABC))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
    message := r.URL.Path
    message = strings.TrimPrefix(message, "/")
    message = "Hello, btanch1 drone got the message: " + message
    log.Println(message)
    w.Write([]byte(message))
}

func main() {

    wordPtr := flag.String("word", "world", "a string")
    vers := flag.Int("vers", 1, "a int")

    flag.Parse()

    fmt.Println(*wordPtr)
    fmt.Println(*vers)

    // use PORT environment variable, or default to 8080
    port := "8080"
    if fromEnv := os.Getenv("PORT"); fromEnv != "" {
        port = fromEnv
    }
    http.HandleFunc("/version", showVersion)
    http.HandleFunc("/env", showEnv)
    http.HandleFunc("/", sayHello)

    log.Println("Listen server on " + port + " port for " + *wordPtr)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}

