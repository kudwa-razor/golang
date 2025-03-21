package main
import "fmt"

func main(){
	messages := make(chan string)
	signals := make(chan bool)

	// here the channel is received
	select{
	case msg := <-messages:
		fmt.Println("Message received", msg)
	default:
		fmt.Println("Not received!")
	}

	//here channel sent
	msg := "hi"
	select{
		case messages <- msg:
			fmt.Println("Mess sent", msg)
		default:
			fmt.Println("not sent!")

	}

	//this is an exmaple of multiple channels, but both are receiving.
	select{
	case msg := <-messages:
		fmt.Println("Received message!", msg)
	
	case sig := <-signals:
		fmt.Println("Sent message!", sig)
	default:
		fmt.Println("both receive and send did not happen!")
	}
}