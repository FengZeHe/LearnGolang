package main

func main() {

	app := InitializeApp()
	server := app.server
	err := server.Run(":8088")
	if err != nil {
		return
	}

}
