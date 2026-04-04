package main

import (
	"log"

	"github.com/indra1nkuss/week4_catalog_inkus/config"
	"github.com/indra1nkuss/week4_catalog_inkus/models"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load env agar bisa konek DB
	godotenv.Load()

	// 2. Konek ke Database
	config.InitDatabase()

	// 3. Daftar data yang mau dimasukkan
	products := []models.Product{
		{Name: "Nasi Goreng Spesial", Price: 25000, Category: "Makanan", Stock: 50,
			Description: "Nasi goreng dengan telur dan ayam",
			ImageURL:    "https://picsum.photos/400"},
		{Name: "Sate Ayam 10 Tusuk", Price: 20000, Category: "Makanan",
			Stock:       100,
			Description: "Sate ayam dengan bumbu kacang",
			ImageURL:    "https://picsum.photos/401"},
		{Name: "Es Teh Manis", Price: 8000, Category: "Minuman",
			Stock:       200,
			Description: "Es teh manis segar",
			ImageURL:    "https://picsum.photos/402"},
		{Name: "Kopi Susu", Price: 15000, Category: "Minuman",
			Stock:       150,
			Description: "Kopi susu kekinian",
			ImageURL:    "https://picsum.photos/403"},
		{Name: "Ayam Bakar", Price: 30000, Category: "Makanan", Stock: 30,
			Description: "Ayam bakar dengan sambal",
			ImageURL:    "https://picsum.photos/404"},
	}

	// 4. Masukkan ke Database
	for _, p := range products {
		config.DB.Create(&p)
	}

	log.Printf("Seed berhasil: %d produk ditambahkan ke database!", len(products))
}