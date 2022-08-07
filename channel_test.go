package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // MUST BE CLOSED !!!

	// channel <- "DAMAR" // CHANNEL TO DATA
	// // "DAMAR" <- channel // DATA TO CHANNEL
	// fmt.Println(<-channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "RADEN ARIO DAMAR"
		fmt.Println("SELESAI MENGIRIM DATA KE CHANNEL")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
