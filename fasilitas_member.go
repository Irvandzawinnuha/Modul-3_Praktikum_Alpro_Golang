package main

import "fmt"

func main() {
    var total_belanja int
    var status_membership string

    fmt.Scanln(&total_belanja, &status_membership)
    if status_membership == "Yes" {
        if total_belanja > 100000 {
            fmt.Println("Anda mendapat diskon 5%, tambahan poin 10, dan free gift")
        } else if total_belanja > 75000 {
            fmt.Println("Anda mendapat diskon 5% dan tambahan poin 5")
        } else {
            fmt.Println("Anda mendapat tambahan poin 5")
        }
    }
}