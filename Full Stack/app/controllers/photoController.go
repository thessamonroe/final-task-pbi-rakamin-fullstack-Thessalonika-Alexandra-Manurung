package controllers

import (
    "net/http"
    "fullStack/app/models"
    "github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
    var photo models.Photo
    if err := c.ShouldBindJSON(&photo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.CreatePhoto(&photo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal untuk menyimpan foto"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Foto berhasil disimpan"})
}

func GetPhotos(c *gin.Context) {
    photos, err := models.GetPhotos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal untuk mendapatkan daftar foto"})
        return
    }

    c.JSON(http.StatusOK, photos)
}

func UpdatePhoto(c *gin.Context) {
    photoID := c.Param("photoId")

    var updatedPhoto models.Photo
    if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.UpdatePhoto(photoID, &updatedPhoto); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal untuk memperbarui foto"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil diperbarui"})
}

func DeletePhoto(c *gin.Context) {
    photoID := c.Param("photoId")

    if err := models.DeletePhoto(photoID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal untuk menghapus foto"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil dihapus"})
}
