package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Struct Pakaian
type Pakaian struct {
	Nama       string
	Tipe       string
	Kategori   []string
	Warna      string
	Formalitas int
	Kehangatan int
}

// Struct untuk menyimpan kombinasi yang pernah dipakai
type Kombinasi struct {
	Atasan          Pakaian
	Bawahan         Pakaian
	AlasKaki        Pakaian
	FormalAvg       float64
	WarmAvg         float64
	TerakhirDipakai time.Time
}

// Data Dummy
var daftarPakaian = []Pakaian{
	{"Kemeja", "atasan", []string{"formal"}, "Putih", 8, 3},
	{"Kaos", "atasan", []string{"casual"}, "Hitam", 3, 2},
	{"Blouse", "atasan", []string{"formal", "casual"}, "Merah", 6, 4},
	{"Hoodie", "atasan", []string{"sporty", "casual"}, "Abu-abu", 4, 7},
	{"Tank Top", "atasan", []string{"casual"}, "Biru", 2, 1},
	{"Kemeja Batik", "atasan", []string{"formal"}, "Batik", 9, 3},
	{"Sweater Rajut", "atasan", []string{"casual"}, "Coklat", 5, 6},

	{"Celana Jeans", "bawahan", []string{"casual"}, "Biru", 4, 4},
	{"Celana Bahan", "bawahan", []string{"formal"}, "Hitam", 8, 5},
	{"Rok Mini", "bawahan", []string{"casual"}, "Merah", 3, 2},
	{"Rok Span", "bawahan", []string{"formal"}, "Abu-abu", 7, 4},
	{"Celana Pendek", "bawahan", []string{"sporty"}, "Hijau", 2, 1},
	{"Legging", "bawahan", []string{"sporty"}, "Hitam", 3, 6},
	{"Rok Lipit", "bawahan", []string{"casual", "formal"}, "Cream", 5, 3},

	{"Sneakers", "alas kaki", []string{"sporty", "casual"}, "Putih", 3, 3},
	{"Pantofel", "alas kaki", []string{"formal"}, "Hitam", 9, 2},
	{"Sandal Jepit", "alas kaki", []string{"casual"}, "Biru", 1, 1},
	{"Boots", "alas kaki", []string{"formal", "casual"}, "Coklat", 7, 7},
}

// Deklarasi Variabel untuk menyimpan outfit yang pernah dipakai
var riwayatKombinasi []Kombinasi

// Cuma untuk tampilan biar terminal kosong
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Tampilan agar harus mengklik Enter sebelum berpindah page
func waitForEnter() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Print("\nTekan Enter untuk melanjutkan...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// Template header tampilan
func tampilkanHeader(judul string) {
	panjang := (50 - len(judul)) / 2
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(strings.Repeat(" ", panjang), judul, strings.Repeat(" ", panjang))
	fmt.Println(strings.Repeat("=", 50))
}

// Menambah Pakaian Baru
func tambahPakaian() {
	clearScreen()
	tampilkanHeader("Tambah Pakaian Baru")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama pakaian : ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Tipe (atasan/bawahan/alas kaki): ")
	tipe, _ := reader.ReadString('\n')
	tipe = strings.TrimSpace(tipe)

	fmt.Println("Kategori : ")
	fmt.Println("Casual")
	fmt.Println("Formal")
	fmt.Println("Sporty")
	fmt.Print("(pisahkan dengan koma, contoh: casual,sporty) : ")
	kategoriStr, _ := reader.ReadString('\n')
	kategoriStr = strings.TrimSpace(kategoriStr)
	kategori := strings.Split(kategoriStr, ",")
	for i := range kategori {
		kategori[i] = strings.TrimSpace(kategori[i])
	}

	fmt.Print("Warna: ")
	warna, _ := reader.ReadString('\n')
	warna = strings.TrimSpace(warna)

	fmt.Print("Tingkat Formalitas (1-10): ")
	formalitasStr, _ := reader.ReadString('\n')
	formalitasStr = strings.TrimSpace(formalitasStr)
	formalitas, _ := strconv.Atoi(formalitasStr)

	fmt.Print("Tingkat Kehangatan (1-10): ")
	kehangatanStr, _ := reader.ReadString('\n')
	kehangatanStr = strings.TrimSpace(kehangatanStr)
	kehangatan, _ := strconv.Atoi(kehangatanStr)

	// Menambahkan ke Struck pakaian
	p := Pakaian{Nama: nama, Tipe: tipe, Kategori: kategori, Warna: warna, Formalitas: formalitas, Kehangatan: kehangatan}
	daftarPakaian = append(daftarPakaian, p)
	fmt.Println("\nPakaian berhasil ditambahkan!")
	waitForEnter()
}

// Menampilkan seluruh pakaian
func tampilkanDaftar() {
	clearScreen()
	tampilkanHeader("Daftar Pakaian")
	if len(daftarPakaian) == 0 {
		fmt.Println("\n Belum ada pakaian yang ditambahkan.")
		return
	}
	for i, p := range daftarPakaian {
		var tipe string

		if p.Tipe == "atasan" {
			tipe = "atasan"
		} else if p.Tipe == "bawahan" {
			tipe = "bawahan"
		} else if p.Tipe == "alas kaki" {
			tipe = "alas kaki"
		} else {
			tipe = p.Tipe
		}
		fmt.Printf("%d. %s (%s)\n", i+1, p.Nama, tipe)
		fmt.Printf("Kategori  : %s\n", strings.Join(p.Kategori, ", "))
		fmt.Printf("Warna     : %s\n", p.Warna)
		fmt.Printf("Formalitas: %d\n", p.Formalitas)
		fmt.Printf("Kehangatan: %d\n", p.Kehangatan)
		fmt.Println(strings.Repeat("-", 50))
	}
}

// Menghapus Pakaian dari struct
func hapusPakaian() {
	tampilkanHeader("Hapus Pakaian")
	if len(daftarPakaian) == 0 {
		fmt.Println(" Daftar pakaian kosong.")
		return
	}
	tampilkanDaftar()
	fmt.Print("Masukkan nomor pakaian yang ingin dihapus: ")
	var index int
	fmt.Scanln(&index)
	if index > 0 && index <= len(daftarPakaian) {
		// Menghapus Pakaian.
		daftarPakaian = append(daftarPakaian[:index-1], daftarPakaian[index:]...)
		fmt.Println(" Pakaian berhasil dihapus.")
	} else {
		fmt.Println(" Nomor tidak valid.")
	}
	waitForEnter()
}

// Mengubah Data Pakaian
func editPakaian() {
	tampilkanHeader("Edit Pakaian")
	if len(daftarPakaian) == 0 {
		fmt.Println(" Daftar pakaian kosong.")
		return
	}
	tampilkanDaftar()
	fmt.Print(" Masukkan nomor pakaian yang ingin diedit: ")
	var index int
	fmt.Scanln(&index)
	if index <= 0 || index > len(daftarPakaian) {
		fmt.Println(" Nomor tidak valid.")
		return
	}
	p := &daftarPakaian[index-1]
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Nama lama: %s\n   Nama baru (kosongkan jika tidak diubah): ", p.Nama)
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama != "" {
		p.Nama = nama
	}

	fmt.Printf("Tipe lama: %s\n   Tipe baru (kosongkan jika tidak diubah): ", p.Tipe)
	tipe, _ := reader.ReadString('\n')
	tipe = strings.TrimSpace(tipe)
	if tipe != "" {
		p.Tipe = tipe
	}

	fmt.Printf("Kategori lama: %v\n   Kategori baru (pisahkan dengan koma, kosongkan jika tidak diubah): ", p.Kategori)
	kategoriStr, _ := reader.ReadString('\n')
	kategoriStr = strings.TrimSpace(kategoriStr)
	if kategoriStr != "" {
		katList := strings.Split(kategoriStr, ",")
		for i := range katList {
			katList[i] = strings.TrimSpace(katList[i])
		}
		p.Kategori = katList
	}

	fmt.Printf("Warna lama: %s\n   Warna baru (kosongkan jika tidak diubah): ", p.Warna)
	warna, _ := reader.ReadString('\n')
	warna = strings.TrimSpace(warna)
	if warna != "" {
		p.Warna = warna
	}

	fmt.Printf("Formalitas lama: %d\n   Formalitas baru (kosongkan jika tidak diubah): ", p.Formalitas)
	formalitasStr, _ := reader.ReadString('\n')
	formalitasStr = strings.TrimSpace(formalitasStr)
	if formalitasStr != "" {
		formalitas, err := strconv.Atoi(formalitasStr)
		if err == nil {
			p.Formalitas = formalitas
		}
	}

	fmt.Printf("Formalitas lama: %d\n   Kehangatan baru (kosongkan jika tidak diubah): ", p.Kehangatan)
	kehangatanStr, _ := reader.ReadString('\n')
	kehangatanStr = strings.TrimSpace(kehangatanStr)
	if kehangatanStr != "" {
		kehangatan, err := strconv.Atoi(kehangatanStr)
		if err == nil {
			p.Kehangatan = kehangatan
		}
	}

	fmt.Println(" Pakaian berhasil diedit.")
	waitForEnter()
}

// Func untuk memberikan nilai formalitas dan nilai kehangatan
// berdasarkan kombinasi outfit yang diberikan
func kombinasiOutfit() {
	clearScreen()
	tampilkanHeader("Kombinasi Outfit")

	var atasan, bawahan, alasKaki *Pakaian

	fmt.Println("Pilih Atasan:")
	for i := range daftarPakaian {
		if daftarPakaian[i].Tipe == "atasan" {
			fmt.Printf("%d. %s -  %s\n", i+1, daftarPakaian[i].Nama, daftarPakaian[i].Warna)
		}
	}
	fmt.Print("Pilih nomor atasan: ")
	var iAtasan int
	fmt.Scanln(&iAtasan)
	atasan = &daftarPakaian[iAtasan-1]

	fmt.Println("\nPilih Bawahan:")
	for i := range daftarPakaian {
		if daftarPakaian[i].Tipe == "bawahan" {
			fmt.Printf("%d. %s -  %s\n", i+1, daftarPakaian[i].Nama, daftarPakaian[i].Warna)
		}
	}
	fmt.Print("Pilih nomor bawahan: ")
	var iBawahan int
	fmt.Scanln(&iBawahan)
	bawahan = &daftarPakaian[iBawahan-1]

	fmt.Println("\nPilih Alas Kaki:")
	for i := range daftarPakaian {
		if daftarPakaian[i].Tipe == "alas kaki" {
			fmt.Printf("%d. %s -  %s\n", i+1, daftarPakaian[i].Nama, daftarPakaian[i].Warna)
		}
	}
	fmt.Print("Pilih nomor alas kaki: ")
	var iAlasKaki int
	fmt.Scanln(&iAlasKaki)
	alasKaki = &daftarPakaian[iAlasKaki-1]

	// Menghitung rata rata tingkat formalitas
	rataFormalitas := float64(atasan.Formalitas+bawahan.Formalitas+alasKaki.Formalitas) / 3
	kategoriGabung := append(append([]string{}, atasan.Kategori...), bawahan.Kategori...)
	kategoriGabung = append(kategoriGabung, alasKaki.Kategori...)

	// Mencari kategori yang dominan
	kategoriCount := map[string]int{}
	for _, k := range kategoriGabung {
		kategoriCount[k]++
	}
	dominan := ""
	max := 0
	for k, v := range kategoriCount {
		if v > max {
			max = v
			dominan = k
		}
	}

	fmt.Printf("\n Tingkat Formalitas Rata-rata: %.2f\n", rataFormalitas)
	fmt.Printf("  Kategori Dominan: %s\n", dominan)
	waitForEnter()
}

// Memberikan rekomendasi berdasarkan ketentuan (outdoor, indoor, cerah, hujan)
func rekomendasiOutfit() {
	clearScreen()
	tampilkanHeader("Rekomendasi Outfit")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Tempat [outdoor/indoor]: ")
	tempat, _ := reader.ReadString('\n')
	tempat = strings.TrimSpace(strings.ToLower(tempat))

	cuaca := ""
	if tempat == "outdoor" {
		fmt.Print("Cuaca [cerah/hujan]: ")
		cuaca, _ = reader.ReadString('\n')
		cuaca = strings.TrimSpace(strings.ToLower(cuaca))
	}

	fmt.Print("Kategori [sporty/casual/formal]: ")
	kategori, _ := reader.ReadString('\n')
	kategori = strings.TrimSpace(strings.ToLower(kategori))

	var atasanList, bawahanList, alasKakiList []Pakaian
	for _, p := range daftarPakaian {
		if !contains(p.Kategori, kategori) {
			continue
		}
		sesuaiCuaca := true
		if cuaca == "cerah" && p.Kehangatan > 5 {
			sesuaiCuaca = false
		}
		if cuaca == "hujan" && p.Kehangatan < 4 {
			sesuaiCuaca = false
		}

		// memasukkan  seluruh atasan, bawahan, alas kaki yang sesuai kriteria.
		if sesuaiCuaca {
			switch p.Tipe {
			case "atasan":
				atasanList = append(atasanList, p)
			case "bawahan":
				bawahanList = append(bawahanList, p)
			case "alas kaki":
				alasKakiList = append(alasKakiList, p)
			}
		}
	}

	// Membuat kombinasi berdasarkan pakaian yang sesuai kriteria.
	var kombinasiList []Kombinasi
	for _, a := range atasanList {
		for _, b := range bawahanList {
			for _, c := range alasKakiList {
				fAvg := float64(a.Formalitas+b.Formalitas+c.Formalitas) / 3
				wAvg := float64(a.Kehangatan+b.Kehangatan+c.Kehangatan) / 3
				kombinasiList = append(kombinasiList, Kombinasi{
					Atasan: a, Bawahan: b, AlasKaki: c,
					FormalAvg: fAvg, WarmAvg: wAvg,
				})
			}
		}
	}

	// Sorting kombinasi berdasarkan preferensi
	if cuaca == "cerah" {
		sort.SliceStable(kombinasiList, func(i, j int) bool {
			if kombinasiList[i].FormalAvg == kombinasiList[j].FormalAvg {
				return kombinasiList[i].WarmAvg < kombinasiList[j].WarmAvg
			}
			return kombinasiList[i].FormalAvg > kombinasiList[j].FormalAvg
		})
	} else if cuaca == "hujan" {
		sort.SliceStable(kombinasiList, func(i, j int) bool {
			if kombinasiList[i].FormalAvg == kombinasiList[j].FormalAvg {
				return kombinasiList[i].WarmAvg > kombinasiList[j].WarmAvg
			}
			return kombinasiList[i].FormalAvg > kombinasiList[j].FormalAvg
		})
	} else {
		sort.SliceStable(kombinasiList, func(i, j int) bool {
			return kombinasiList[i].FormalAvg > kombinasiList[j].FormalAvg
		})
	}

	// Menampilkan maksimal 3 kombinasi
	fmt.Println("\n🎽 Rekomendasi Outfit:")
	jumlah := 0
	for i, k := range kombinasiList {
		fmt.Printf("\nOutfit %d:\n", i+1)
		fmt.Printf(" Atasan   : %s (%s)\n", k.Atasan.Nama, k.Atasan.Warna)
		fmt.Printf(" Bawahan  : %s (%s)\n", k.Bawahan.Nama, k.Bawahan.Warna)
		fmt.Printf(" Alas Kaki: %s (%s)\n", k.AlasKaki.Nama, k.AlasKaki.Warna)
		fmt.Printf(" Formalitas: %.2f |  Kehangatan: %.2f\n", k.FormalAvg, k.WarmAvg)
		jumlah++
		if jumlah == 3 {
			break
		}
	}

	if jumlah == 0 {
		fmt.Println("Tidak ada kombinasi outfit yang cocok ditemukan.")
	} else {
		fmt.Print("\nPilih nomor outfit yang ingin kamu pakai : ")
		pilihStr, _ := reader.ReadString('\n')
		pilihStr = strings.TrimSpace(pilihStr)
		pilih, err := strconv.Atoi(pilihStr)
		if err == nil && pilih >= 1 && pilih <= 3 {
			dipilih := kombinasiList[pilih-1]
			dipilih.TerakhirDipakai = time.Now()

			// Cek apakah kombinasi sudah ada di riwayat
			found := false
			for i, k := range riwayatKombinasi {
				if samaKombinasi(k, dipilih) {
					riwayatKombinasi[i].TerakhirDipakai = time.Now()
					found = true
					break
				}
			}
			if !found {
				riwayatKombinasi = append(riwayatKombinasi, dipilih)
			}

			fmt.Println("Outfit telah disimpan sebagai terakhir kali dipakai.")
		}
	}

	waitForEnter()
}

// Pengecekan apakah terdapat kombinasi yang sama pada sebuah rekomendasi outfit
func samaKombinasi(a, b Kombinasi) bool {
	return a.Atasan.Nama == b.Atasan.Nama &&
		a.Bawahan.Nama == b.Bawahan.Nama &&
		a.AlasKaki.Nama == b.AlasKaki.Nama
}

// Pengecekan pada rekomendasi outfit apakah user menginputkan kategori yang sesuai
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Mencari pakaian berdasarkan warna (Sequential Search)
func cariBerdasarkanWarna() {
	clearScreen()
	tampilkanHeader("Cari Pakaian Berdasarkan Warna")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan warna yang dicari: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	ditemukan := false

	fmt.Println("")

	// Sequential Search
	for _, p := range daftarPakaian {
		if strings.ToLower(p.Warna) == input {
			fmt.Printf("- %s (%s), Warna: %s, Kategori: %s\n", p.Nama, p.Tipe, p.Warna, strings.Join(p.Kategori, ", "))
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ditemukan pakaian dengan warna tersebut.")
	}
	waitForEnter()
}

// Mencari Pakaian berdasarkan kategori (Binary Search)
func cariBerdasarkanKategori() {
	clearScreen()
	tampilkanHeader("Cari Pakaian Berdasarkan Kategori")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kategori yang dicari (casual/formal/sporty): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	fmt.Println("")
	// Buat salinan daftar dan urutkan berdasarkan kategori pertama
	salinan := make([]Pakaian, len(daftarPakaian))
	copy(salinan, daftarPakaian)

	sort.Slice(salinan, func(i, j int) bool {
		if len(salinan[i].Kategori) == 0 || len(salinan[j].Kategori) == 0 {
			return false
		}
		return salinan[i].Kategori[0] < salinan[j].Kategori[0]
	})

	// Binary search
	low, high := 0, len(salinan)-1
	ditemukan := false
	for low <= high {
		mid := (low + high) / 2
		if len(salinan[mid].Kategori) == 0 {
			break
		}
		k := salinan[mid].Kategori[0]
		if k == input {
			// Kategori cocok, cari semua yang cocok (karena bisa lebih dari satu)
			ditemukan = true
			i := mid
			for i >= 0 && len(salinan[i].Kategori) > 0 && salinan[i].Kategori[0] == input {
				i--
			}
			i++
			for i < len(salinan) && len(salinan[i].Kategori) > 0 && salinan[i].Kategori[0] == input {
				fmt.Printf("- %s (%s), Kategori: %s, Warna: %s\n", salinan[i].Nama, salinan[i].Tipe, strings.Join(salinan[i].Kategori, ", "), salinan[i].Warna)
				i++
			}
			break
		} else if k < input {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ditemukan pakaian dengan kategori tersebut.")
	}
	waitForEnter()
}

// Mengurutkan Pakaian Berdasarkan Formalitas (Selection Sort)
func urutkanPakaianBerdasarkanFormalitas(pakaian []Pakaian) {
	// Selection Sort
	n := len(pakaian)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if pakaian[j].Formalitas > pakaian[maxIdx].Formalitas {
				maxIdx = j
			}
		}
		pakaian[i], pakaian[maxIdx] = pakaian[maxIdx], pakaian[i]
	}
}

// Mengurutkan pakaian Berdasarkan Terakhir Dipakai (Insertion Sort)
func urutkanBerdasarkanTerakhirDipakai(k []Kombinasi) {
	// Insertion Sort
	for i := 1; i < len(k); i++ {
		key := k[i]
		j := i - 1
		for j >= 0 && k[j].TerakhirDipakai.Before(key.TerakhirDipakai) {
			k[j+1] = k[j]
			j--
		}
		k[j+1] = key
	}
}

// Menu Utama
func main() {
	for {
		clearScreen()
		tampilkanHeader("OOTD Planner")
		fmt.Println("1. Tambah Pakaian")
		fmt.Println("2. Lihat Daftar Pakaian")
		fmt.Println("3. Edit Pakaian")
		fmt.Println("4. Hapus Pakaian")
		fmt.Println("5. Cek Kombinasi Outfit")
		fmt.Println("6. Rekomendasi Outfit")
		fmt.Println("7. Cari Pakaian")
		fmt.Println("8. Urutkan Pakaian")
		fmt.Println("0. Keluar")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Print("Pilih menu (0-8): ")
		var pilihan int
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			tambahPakaian()
		case 2:
			tampilkanDaftar()
			waitForEnter()
		case 3:
			editPakaian()
		case 4:
			hapusPakaian()
		case 5:
			kombinasiOutfit()
		case 6:
			rekomendasiOutfit()
		case 7:
			clearScreen()
			tampilkanHeader("Cari Pakaian")
			fmt.Println("1. Berdasarkan warna (Sequential Search)")
			fmt.Println("2. Berdasarkan Kategori (Binary Search)")
			fmt.Println(strings.Repeat("=", 50))
			fmt.Print("Pilih menu (1-2): ")

			var pilihancari int
			fmt.Scanln(&pilihancari)
			clearScreen()

			switch pilihancari {
			case 1:
				cariBerdasarkanWarna()
			case 2:
				cariBerdasarkanKategori()
			default:
				fmt.Println("Pilihan tidak valid, coba lagi.")
				waitForEnter()
			}
		case 8:
			clearScreen()
			tampilkanHeader("Urutkan Pakaian")
			fmt.Println("1. Berdasarkan formalitas (Selection Sort)")
			fmt.Println("2. Berdasarkan terakhir dipakai (Insertion Sort)")
			fmt.Println(strings.Repeat("=", 50))
			fmt.Print("Pilih menu (1-2): ")

			var pilihanUrutan int
			fmt.Scanln(&pilihanUrutan)
			clearScreen()

			switch pilihanUrutan {
			case 1:
				urutkanPakaianBerdasarkanFormalitas(daftarPakaian)

				tampilkanHeader("Daftar Pakaian berdasarkan Formalitas")
				for _, p := range daftarPakaian {
					fmt.Printf("- %s (%s), Formalitas: %d\n", p.Nama, p.Tipe, p.Formalitas)
				}

				waitForEnter()
			case 2:
				tampilkanHeader("Daftar Pakaian berdasarkan Terakhir Dipakai")
				if len(riwayatKombinasi) != 0 {
					urutkanBerdasarkanTerakhirDipakai(riwayatKombinasi)

					for i, k := range riwayatKombinasi {
						fmt.Printf("\nRiwayat %d [%s]\n", i+1, k.TerakhirDipakai.Format("02 Jan 2006"))
						fmt.Printf("Atasan   : %s (%s)\n", k.Atasan.Nama, k.Atasan.Warna)
						fmt.Printf("Bawahan  : %s (%s)\n", k.Bawahan.Nama, k.Bawahan.Warna)
						fmt.Printf("Alas Kaki: %s (%s)\n", k.AlasKaki.Nama, k.AlasKaki.Warna)
						fmt.Printf("Formalitas: %.2f |  Kehangatan: %.2f\n", k.FormalAvg, k.WarmAvg)
					}

					waitForEnter()
				} else {
					fmt.Println("Belum ada riwayat outfit.")
					waitForEnter()
				}
			default:
				fmt.Println("Pilihan tidak valid, coba lagi.")
				waitForEnter()
			}
		case 0:
			fmt.Println("Terima kasih telah menggunakan OOTD Planner!")
			return
		default:
			fmt.Println("Pilihan tidak valid, coba lagi.")
			waitForEnter()
		}
	}
}
