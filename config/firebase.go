package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// FirebaseAuth adalah instance Firebase Auth yang dipakai untuk verify token
var FirebaseAuth *auth.Client

func InitFirebase() {
	// Ambil path file dari variabel di dalam .env
	credPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if credPath == "" {
		credPath = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	}

	// Jika masih kosong, beri peringatan yang jelas
	if credPath == "" {
		log.Fatal("ERROR: Path kredensial Firebase tidak ditemukan. Pastikan FIREBASE_CREDENTIALS_PATH ada di file .env")
	}

	// Inisialisasi Firebase App dengan service account credentials
	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Gagal init Firebase: %v", err)
	}

	// Dapatkan Firebase Auth client
	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Gagal mendapatkan Firebase Auth client: %v", err)
	}

	log.Println("Firebase Admin SDK berhasil diinisialisasi")
}