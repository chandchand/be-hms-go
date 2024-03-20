package controllers

import (
	"net/http"

	"github.com/chandchand/be-hms-go/database"
	"github.com/chandchand/be-hms-go/database/models"
	"github.com/chandchand/be-hms-go/services"
	"github.com/gin-gonic/gin"
)

func CreateBed(c *gin.Context) {
	var bed models.BedType
	if err := c.ShouldBindJSON(&bed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&bed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Created data success", "data": bed})
}

func GetAllBeds(c *gin.Context)  {
	var bed []models.BedType

	if err := database.DB.Find(&bed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if bed == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no bed data found!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get Data Success", "data": bed})
}

func GetBed(c *gin.Context) {
	id := c.Param("id")
	bed, err := services.GetBedByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "bed with id " + id + " not found", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get one data guest success", "data": bed})
}

func UpdateBed(c *gin.Context) {
	id := c.Param("id")
	
	bed, err := services.GetBedByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "bed with id " + id + " not found", "error": err.Error()})
		return
	}

	var updatedBed models.Guest
	if err := c.ShouldBindJSON(&updatedBed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&bed).Updates(&updatedBed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to update guest","error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message":"Updated guest successfully", "data": bed})
}

func DeleteBed(c *gin.Context) {
	id := c.Param("id")
	bed, err := services.GetBedByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "bed with id " + id + " not found", "error": err.Error()})
		return
	}

	if err := database.DB.Delete(&bed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed Delete Guest", "error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Guest deleted successfully","data": bed})
}

func CreateTypeRoom(c *gin.Context) {
	var rt models.RoomType
	if err := c.ShouldBindJSON(&rt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, bedType := range rt.BedTypes {
        if err := database.DB.First(&models.BedType{}, "id = ?", bedType.ID).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Bed Type ID"})
            return
        }
    }
	

	if err := database.DB.Create(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Created data success", "data": rt})
}

func GetAllTypeRooms(c *gin.Context) {
	var rt []models.RoomType

	if err := database.DB.Preload("BedTypes").Find(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(rt) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no room type data found!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get Data Success", "data": rt})
}

func GetTypeRoom(c *gin.Context) {
	id := c.Param("id")
	rt, err := services.GetRoomTypeByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "room type with id " + id + " not found", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get one data room type success", "data": rt})
}

func UpdateTypeRoom(c *gin.Context) {
	id := c.Param("id")
	rt, err := services.GetRoomTypeByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "room type with id " + id + " not found", "error": err.Error()})
		return
	}

	var updatedRT models.Guest
	if err := c.ShouldBindJSON(&updatedRT); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&rt).Updates(&updatedRT).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to update room type","error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message":"Updated room type successfully", "data": rt})
}

func DeleteTypeRoom(c *gin.Context) {
	id := c.Param("id")
	rt, err := services.GetRoomTypeByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "room type with id " + id + " not found", "error": err.Error()})
		return
	}

	if err := database.DB.Delete(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed Delete Room Type", "error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Room Type deleted successfully","data": rt})
}

func CreateRoom(c *gin.Context) {
	var r models.Room
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var rt models.RoomType
	if err := database.DB.First(&rt, "id = ?", r.RoomTypeID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Room Type ID"})
		return
	}
	// Setelah mendapatkan BedType dari database, tambahkan ke RoomType
	r.RoomType = rt

	// fmt.Println("id", rt.BedType.ID)

	if err := database.DB.Create(&r).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Created data success", "data": r})
}

func GetAllRooms(c *gin.Context) {
	var r []models.Room

	if err := database.DB.Preload("RoomType.BedTypes").Find(&r).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(r) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no room type data found!!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get Data Success", "data": r})
}

func GetRoom(c *gin.Context) {
	id := c.Param("id")
	r, err := services.GetRoomByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "room with id " + id + " not found", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Get one data room type success", "data": r})
}

func UpdateRoom(c *gin.Context) {
	id := c.Param("id")
	r, err := services.GetRoomByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "room with id " + id + " not found", "error": err.Error()})
		return
	}

	var updatedR models.Guest
	if err := c.ShouldBindJSON(&updatedR); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&r).Updates(&updatedR).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Failed to update room type","error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message":"Updated room type successfully", "data": r})
}

func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	r, err := services.GetRoomByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "room with id " + id + " not found", "error": err.Error()})
		return
	}

	if err := database.DB.Delete(&r).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed Delete Room Type", "error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Room Type deleted successfully","data": r})
}




