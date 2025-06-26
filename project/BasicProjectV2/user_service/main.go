package main

func main() {

	app := InitializeApp()
	err := app.Start()
	if err != nil {
		return
	}
}
