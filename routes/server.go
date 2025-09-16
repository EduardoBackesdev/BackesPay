package routes

func Server() {
	r := Router()
	r.Run("localhost:8080")
}
