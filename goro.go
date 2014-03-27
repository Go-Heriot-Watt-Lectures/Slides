func DoSomeHardWork(chan done) {
	...
	done <- true
}

func main() {
	// first call using named function
	done := make(chan bool)
	go DoSomeHardWork(done)
	
	// unnamed function literal
	go func(done) {
		...
		done <- true
	}(done)
	<- done
}
