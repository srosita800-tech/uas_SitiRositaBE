package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/srosita800-tech/uas_1125170150BE/config"
	"github.com/srosita800-tech/uas_1125170150BE/models"
)

type CartHandler struct{}

func NewCartHandler() *CartHandler {
	return &CartHandler{}
}

// Helper untuk mendapatkan UserID dari context
func getUserID(c *gin.Context) uint {
	val, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	switch v := val.(type) {
	case float64:
		return uint(v)
	case string:
		id, _ := strconv.ParseUint(v, 10, 32)
		return uint(id)
	case int:
		return uint(v)
	case uint:
		return v
	default:
		return 0
	}
}

// GetCart - Mendapatkan isi keranjang user yang sedang login
func (h *CartHandler) GetCart(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak valid"})
		return
	}

	var cartItems []models.Cart
	if err := config.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mengambil keranjang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cartItems,
	})
}

// AddToCart - Menambah produk ke keranjang
func (h *CartHandler) AddToCart(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak valid"})
		return
	}

	var input struct {
		ProductID uint `json:"product_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Input tidak valid"})
		return
	}

	var cart models.Cart
	err := config.DB.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&cart).Error

	if err == nil {
		// Jika sudah ada, tambah quantity
		cart.Quantity += 1
		config.DB.Save(&cart)
	} else {
		// Jika belum ada, buat baru
		cart = models.Cart{
			UserID:    uint(userID),
			ProductID: input.ProductID,
			Quantity:  1,
		}
		config.DB.Create(&cart)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Produk berhasil ditambah ke keranjang"})
}

// ReduceQuantity - Mengurangi quantity produk di keranjang
func (h *CartHandler) ReduceQuantity(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak valid"})
		return
	}

	var input struct {
		ProductID uint `json:"product_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Input tidak valid"})
		return
	}

	var cart models.Cart
	if err := config.DB.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Produk tidak ada di keranjang"})
		return
	}

	if cart.Quantity > 1 {
		cart.Quantity -= 1
		config.DB.Save(&cart)
	} else {
		config.DB.Delete(&cart)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Quantity berhasil dikurangi"})
}

// ClearCart - Mengosongkan keranjang (biasanya setelah bayar)
func (h *CartHandler) ClearCart(c *gin.Context) {
	userID := getUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak valid"})
		return
	}

	if err := config.DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mengosongkan keranjang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Keranjang berhasil dikosongkan"})
}
