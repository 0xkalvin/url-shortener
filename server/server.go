package server

// Run app
func Run() {
	var port = ":3000"

	router := initializeRouter()

	router.Run(port)
}
