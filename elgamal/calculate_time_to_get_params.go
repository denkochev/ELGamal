package elgamal

import (
	"fmt"
	"time"
)

func TimeToGenerateParams() {
	start := time.Now()
	GenerateBigSafePrime(true)

	duration := time.Since(start)
	fmt.Println("Час знаходження валідних початкових параметрів: ", duration)
}
