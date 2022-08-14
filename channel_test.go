package go_routines

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// channel hanya bisa menampung 1 DATA
// channel bisa digunakan untuk mengirim data dan menerima data

// func TestCreateChannel(t *testing.T) {
// 	channel := make(chan string)
// 	defer close(channel) // wajib close channel

// 	// mengirim data ke channel
// 	channel <- "Damar"

// 	// menerima data dari channel
// 	data := <-channel
// 	fmt.Println("CHANNEL", channel)
// 	fmt.Println("DATA", data)
// }

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // wajib close channel

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Raden Ario Damar" // program stop there, becuase no one getting this channel
		fmt.Println("Selesai Mengirim data ke Channel")
	}()

	data := <-channel // will error deadlock! kalau channel tidak ada data nya
	fmt.Println("DATA", data)

	time.Sleep(5 * time.Second)
}

// CHANNEL AS PARAMETER, no need POINTER!
func GiveMeResponse(channel chan string) {
	fmt.Println("give me response channel XD")
	channel <- "Raden Ario Damar"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // MUST be closing the channel

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("DATA Channel:", data)

	time.Sleep(5 * time.Second)
}

// ./CHANNEL AS PARAMETER, no need POINTER!

// CHANNEL In & Out
func OnlyIn(channel chan<- string) { // hanya untuk mengirim chan<-
	fmt.Println("give me response channel XD")
	// data := <-channel // WILL Throw error invalid operation: cannot receive from send-only channel channel (variable of type chan<- string)
	channel <- "Raden Ario Damar"
}

func OnlyOut(channel <-chan string) { // hanya untuk menerima <-chan
	data := <-channel
	// channel <- "Raden Ario Damar" // WILL Throw error invalid operation: cannot send to receive-only channel channel (variable of type <-chan string)
	fmt.Println("DATA Channel:", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // MUST be closing the channel

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// ./CHANNEL In & Out

// BUFFERED CHANNEL
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // buffer 3
	defer close(channel)            // wajib close channel

	// WITH GO ROUTINE
	go func() {
		// mengirim data ke channel
		channel <- "Raden"
		channel <- "Ario"
		channel <- "Damar"
	}()

	go func() {
		// menerima data dari channel
		fmt.Println(<-channel) // 1
		fmt.Println(<-channel) // 2
		fmt.Println(<-channel) // 3
	}()

	// WITHOUT GO ROUTINE
	// // mengirim data ke channel
	// channel <- "Raden"
	// channel <- "Ario"
	// channel <- "Damar"

	// // menerima data dari channel
	// fmt.Println(<-channel) // 1
	// fmt.Println(<-channel) // 2
	// fmt.Println(<-channel) // 3

	fmt.Println("END!!!")
}

// ./BUFFERED CHANNEL

// RANGE CHANNEL
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Loop - " + strconv.Itoa(i)
		}
		// close channel here
		close(channel)

	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("END !!!")
}

// ./RANGE CHANNEL

// SELECT CHANNEL
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// ./SELECT CHANNEL

// DEFAULT SELECT CHANNEL
func TestDefaulSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data . . . ")
		}
		if counter == 2 {
			break
		}
	}
}

// ./SELECT CHANNEL

// RACE CONDITION Using Mutex
func TestRaceCondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex // FOR LOCKING to Handle Race Condition

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1 // <--- Race Condition, for make sure the counter result is 100000 or expected what u want
				mutex.Unlock()

			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", x)
}

// ./RACE CONDITION  Using Mutex

// RACE CONDITION using RWMutex
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance", account.GetBalance())
}

// ./RACE CONDITION  Using RWMutex

// DEADLOCK Simulation
type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Tranfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Damar",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Benny",
		Balance: 1000000,
	}

	go Tranfer(&user1, &user2, 100000)
	go Tranfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}

// ./DEADLOCK Simulation
