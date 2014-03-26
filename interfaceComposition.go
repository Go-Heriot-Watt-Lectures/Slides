type StreamReadWriter struct {
	io.ReadWriter
	...
}

func main() {
	d := new(StreamReadWriter)
	n, err := d.Write([]byte("Interface composition in action!")
	if err != nil {
		log.Println("An error occurred.")
	} else {
		log.Println(fmt.Sprintf("%i bytes were written",n))
	}
}
