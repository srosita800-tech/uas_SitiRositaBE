package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/srosita800-tech/uas_1125170150BE/handlers"
	"github.com/srosita800-tech/uas_1125170150BE/middleware"
)

func SetupRouter() *gin.Engine {
	// gin.Default() sudah include Logger & Recovery middleware
	r := gin.Default()

	// ─── CORS Middleware (izinkan request dari Flutter app) ───
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ─── Init handlers ────────────────────────────────────────
	authHandler := handlers.NewAuthHandler()
	productHandler := handlers.NewProductHandler()
	cartHandler := handlers.NewCartHandler()

	// ─── API v1 group ─────────────────────────────────────────
	v1 := r.Group("/v1")
	{
		// Health check — tidak perlu auth
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "service": "Backend Catalog Srosita"})
		})

		// ── Auth routes (public) ──────────────────────────────
		auth := v1.Group("/auth")
		{
			// Utama: Verifikasi Firebase token → return Backend JWT
			auth.POST("/verify-token", authHandler.VerifyToken)
			
			// Tambahan untuk kompatibilitas frontend
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// ── Protected routes (butuh Backend JWT) ──────────────
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Products — semua user terautentikasi bisa GET
			products := protected.Group("/products")
			{
				products.GET("", productHandler.GetAll)      // GET /v1/products
				products.GET("/:id", productHandler.GetByID) // GET /v1/products/:id

				// Create, Update, Delete — hanya admin
				adminProducts := products.Group("")
				adminProducts.Use(middleware.AdminOnly())
				{
					adminProducts.POST("", productHandler.Create)       // POST /v1/products
					adminProducts.PUT("/:id", productHandler.Update)    // PUT /v1/products/:id
					adminProducts.DELETE("/:id", productHandler.Delete) // DELETE /v1/products/:id
				}
			}

			// ── Cart routes ───────────────────────────────────────
			cart := protected.Group("/cart")
			{
				cart.GET("", cartHandler.GetCart)             // GET /v1/cart
				cart.POST("/add", cartHandler.AddToCart)      // POST /v1/cart/add
				cart.POST("/reduce", cartHandler.ReduceQuantity) // POST /v1/cart/reduce
				cart.DELETE("", cartHandler.ClearCart)        // DELETE /v1/cart
			}
		}
	}

	return r
}