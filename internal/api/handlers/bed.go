package handlers

import (
	"fmt"
	"main/internal/api/services"
	"main/internal/repositories"

	"github.com/gin-gonic/gin"
)

func Bed(r *gin.Context) {

	var a repositories.BedRequest
	if err := r.ShouldBindJSON(&a); err != nil {
		r.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error json request: %v", err)})
		return
	}
	if err := services.Bed(a); err != nil {
		r.JSON(400, ErrorResponse{Message: fmt.Sprintf("Error with bed: %v", err)})
		return
	}
	return

}
