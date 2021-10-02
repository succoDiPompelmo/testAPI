package main

func main() {
	a := App{}
	a.Init()
	// port := os.Getenv("APP_PORT")
	a.Run(":8010")
}