package main

import (
    "fmt"
    "math/rand"
    "time"
)

var kalimat = []string{
    "Jangan lupa tersenyum hari ini!",
    "Teruslah berjuang dan jangan menyerah.",
    "Kebaikan sekecil apapun akan dihargai.",
    "Be your self",
    "Hari ini adalah kesempatan baru.",
    "Hidup sesuka hati",
    "Bersyukur atas apa yang kita miliki.",
}

func main() {
    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(sentences))
    randomSentence := sentences[randomIndex]
    fmt.Println(randomSentence)
}
