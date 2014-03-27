
// instantiation

c := make(chan string)

// write to channel

c <- "Tom"

// read off channel into variable
value := <- c

// close channel to release reference
close(c)
