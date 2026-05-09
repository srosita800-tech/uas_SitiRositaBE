# Catalog Srosita 

## 📝 Catatan Mengenai Riwayat Commit

**Maaf sebelumnya**, jika Anda melihat total *commit* di repositori ini sangat sedikit. 

Sebenarnya, saya telah mengerjakan proyek ini secara bertahap dan melakukan *commit* rutin untuk setiap perubahan file. Namun, karena kekhilafan saya yang **lupa menyembunyikan file rahasia (`.env` dan Service Account Key Firebase)** pada *commit-commit* awal, kredensial keamanan tersebut sempat terekspos ke dalam riwayat Git.

Demi menjaga keamanan aplikasi dan mengikuti *best practice*, saya mengambil langkah berikut:
1. **Menghapus folder `.git` secara paksa** untuk membuang seluruh riwayat lama yang mengandung data sensitif.
2. **Melakukan inisialisasi ulang Git** dari titik nol yang sudah bersih.
3. **Menghubungkan kembali** ke repositori ini dengan kondisi file rahasia sudah masuk ke `.gitignore`.

Meskipun jumlah *commit*-nya terlihat sedikit, seluruh fitur di sini adalah hasil dari proses pengembangan yang panjang. Pelajaran berharga: **Selalu cek `.gitignore` sebelum melakukan `git add .`!**

