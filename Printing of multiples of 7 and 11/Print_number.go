package main

func main() {
	for i := 0; i <= 150; i++ {

		if (i%7) == 0 && (i%11) == 0 {
			println("Múltiplo de 7 e 11")
		} else if (i % 7) == 0 {
			println("Múltiplo de 7")
		} else if (i % 11) == 0 {
			println("Múltiplo de 11")
		} else {
			println(i)
		}
	}
}
