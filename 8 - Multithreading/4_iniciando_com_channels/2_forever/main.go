package main

// Thread 1 (Principal)
func main() {
	forever := make(chan bool)

	go func() {
		for i := range 10 {
			println(i)
		}
		forever <- true
	}()

	<-forever
}
