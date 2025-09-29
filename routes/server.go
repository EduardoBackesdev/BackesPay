package routes

func Server() {
	r := Router()
	r.Run(":8080")
}
