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
    "Hari ini adalah kesempatan baru.",
    "Bersyukur atas apa yang kita miliki.",
}

func main() {
    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(sentences))
    randomSentence := sentences[randomIndex]
    fmt.Println(randomSentence)
}
