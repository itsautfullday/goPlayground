package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)
func main() {
    // Get a greeting message and print it.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    names := []string{"Gladys", "Samantha", "Darrin"}
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(messages)
}
