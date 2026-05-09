package main

import (
	"log"

	"github.com/srosita800-tech/uas_1125170150BE/config"
	"github.com/srosita800-tech/uas_1125170150BE/models"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load env agar bisa konek DB
	godotenv.Load()

	// 2. Konek ke Database
	config.InitDatabase()

	// 3. Daftar data yang mau dimasukkan
	products := []models.Product{
		{Name: "Tas Ransel Laptop Modern", Price: 350000, Category: "Tas Pria", Stock: 25,
			Description: "Tas ransel dengan kompartemen laptop 15 inch dan desain minimalis waterproof.",
			ImageURL:    "https://images.unsplash.com/photo-1553062407-98eeb64c6a62?q=80&w=1000&auto=format&fit=crop"},
		{Name: "Tas Selempang Kulit Casual", Price: 225000, Category: "Tas Pria",
			Stock:       40,
			Description: "Tas selempang bahan kulit sintetis premium untuk kegiatan sehari-hari.",
			ImageURL:    "https://images.unsplash.com/photo-1590874103328-eac38a683ce7?q=80&w=1000&auto=format&fit=crop"},
		{Name: "Tote Bag Kanvas Vintage", Price: 85000, Category: "Tas Wanita",
			Stock:       100,
			Description: "Tote bag bahan kanvas tebal dengan motif estetik vintage.",
			ImageURL:    "https://images.unsplash.com/photo-1544816155-12df9643f363?q=80&w=1000&auto=format&fit=crop"},
		{Name: "Tas Tangan Elegan", Price: 450000, Category: "Tas Wanita",
			Stock:       15,
			Description: "Handbag premium untuk acara formal dengan aksen emas yang mewah.",
			ImageURL:    "https://images.unsplash.com/photo-1584917865442-de89df76afd3?q=80&w=1000&auto=format&fit=crop"},
		{Name: "Tas Gunung 45L Professional", Price: 750000, Category: "Tas Outdoor", Stock: 10,
			Description: "Tas gunung kapasitas 45 liter dengan sistem sirkulasi udara di punggung.",
			ImageURL:    "https://images.unsplash.com/photo-1622560480605-d83c853bc5c3?q=80&w=1000&auto=format&fit=crop"},
	}

	// 4. Masukkan ke Database
	for _, p := range products {
		config.DB.Create(&p)
	}

	log.Printf("Seed berhasil: %d produk ditambahkan ke database!", len(products))
}