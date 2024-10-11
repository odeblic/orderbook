package main

func main() {
	connectToDatabase("127.0.0.1", 3301, "moneyboy", "secret1234")
	defer disconnectFromDatabase()

	printSpace()
	printTitle("Trading program")
	printMessage("Populate the database...")
	populateDatabase()
	printMessage("Show the details of the account...")
	checkAccount()
	printMessage("Enter the main loop...")

	for {
		consumeOrderQueue()
		printSpace()
		printMessage("Pause...")
		pause()
	}
}
