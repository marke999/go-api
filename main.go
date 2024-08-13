package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

var jokes = []string{
    "Why don't scientists trust atoms? Because they make up everything!",
    "Why did the scarecrow win an award? Because he was outstanding in his field!",
    "What do you call fake spaghetti? An impasta!",
}

var names = []string{"Mark", "Richard", "James", "Cassandra"}
var ages = []string{"21", "24", "29", "27"}
var occupations = []string{"nurse", "engineer", "developer", "musician"}
var devices = []string{"laptop", "mobile device", "smart TV", "watch"}
var bodyParts = []string{"left finger", "right arm", "left shoulder", "right foot"}
var moods = []string{"excited", "gloomy", "disgusted", "overwhelmed"}
var actions = []string{"playing music","narrating stories", "drawing", "researching events"}

func getRandomElement(arr []string) string {
    return arr[rand.Intn(len(arr))]
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
    joke := getRandomElement(jokes)
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(joke))
}

func madlibHandler(w http.ResponseWriter, r *http.Request) {
    name := getRandomElement(names)
    age := getRandomElement(ages)
    occupation := getRandomElement(occupations)
    device := getRandomElement(devices)
    bodyPart := getRandomElement(bodyParts)
    mood := getRandomElement(moods)
    action := getRandomElement(actions)

    madlib := fmt.Sprintf("%s is a %s-year old %s who has been struggling with a lot of job-related stress. "+
        "He/she decided to try a new application to relieve stress, which runs on a %s to help improve his/her mood.\n"+
        "The application senses his/her mood through a device he/she wears on his/her %s.\n"+
        "When the device senses that he/she is %s, it responds by %s.",
        name, age, occupation, device, bodyPart, mood, action)

    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(madlib))
}

func main() {
    rand.Seed(time.Now().UnixNano())

    http.HandleFunc("/joke", jokeHandler)
    http.HandleFunc("/madlib", madlibHandler)

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Server failed:", err)
    }
}
