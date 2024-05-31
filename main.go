package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fungsi aktivasi step function
func fungsiAktivasi(x float64) int {
	if x > 0 {
		return 1
	}
	return 0
}

// Struktur untuk JST (NN)
type NN struct {
	aktivasi float64
	error    float64
	weight   [3]float64
	miu      float64
	no       int
}

// Konstruktor untuk JST (NN)
func NewNN() *NN {
	nn := &NN{
		miu: 0.1,
		no:  1,
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		nn.weight[i] = rand.Float64()
	}
	return nn
}

// Metode untuk menghitung output
func (nn *NN) hitungOutput(x1, x2 int) int {
	nn.aktivasi = nn.weight[0] + float64(x1)*nn.weight[1] + float64(x2)*nn.weight[2]
	return fungsiAktivasi(nn.aktivasi)
}

// Metode untuk melakukan pembelajaran
func (nn *NN) latih(x1, x2, target int) {
	output := nn.hitungOutput(x1, x2)
	nn.error = float64(target - output)

	fmt.Printf("%-4d%-3d%-3d%-3d%-8.3f%-8.3f%-8.3f%-10.3f%-7d%-10d%-5.3f\n",
		nn.no, 1, x1, x2, nn.weight[0], nn.weight[1], nn.weight[2], nn.aktivasi, output, target, nn.error)
	nn.no++

	// Mengupdate bobot
	if nn.error != 0 {
		nn.weight[0] += nn.miu * nn.error
		nn.weight[1] += nn.miu * nn.error * float64(x1)
		nn.weight[2] += nn.miu * nn.error * float64(x2)
	}
}

// Metode untuk menguji model pada data uji
func (nn *NN) uji(x1, x2 int) int {
	output := nn.hitungOutput(x1, x2)
	return output
}

func main() {
	for {
		var pilihan int
		fmt.Println("\nPilih opsi:")
		fmt.Println("1. Studi Kasus OR")
		fmt.Println("2. Studi Kasus AND")
		fmt.Println("3. Studi Kasus AND NOT")
		fmt.Println("0. Keluar")
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan program.")
			break
		}

		totalSukses := 0
		maxEpochs := 1000

		var input [][2]int
		var targetOutput [4]int

		// Menentukan input dan targetOutput berdasarkan pilihan
		switch pilihan {
		case 1: // Studi Kasus OR
			input = [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
			targetOutput = [4]int{0, 1, 1, 1}
		case 2: // Studi Kasus AND
			input = [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
			targetOutput = [4]int{0, 0, 0, 1}
		case 3: // Studi Kasus AND NOT
			input = [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
			targetOutput = [4]int{1, 0, 1, 0}
		default:
			fmt.Println("Pilihan tidak valid.")
			continue
		}

		AND := NewNN()

		// Melakukan pembelajaran
		fmt.Printf("%-4s%-3s%-3s%-3s%-8s%-8s%-8s%-10s%-7s%-10s%-5s\n",
			"no", "1", "x1", "x2", "w1", "w2", "w3", "target", "error")
		for totalSukses < 4 && AND.no/4 < maxEpochs {
			fmt.Printf("\nEpoch: %d\n", AND.no/4+1)
			for i := 0; i < 4; i++ {
				x1 := input[i][0]
				x2 := input[i][1]
				target := targetOutput[i]

				AND.latih(x1, x2, target)

				if AND.error == 0 && input[i][0] == 0 && input[i][1] == 0 {
					totalSukses = 1
				} else if AND.error == 0 {
					totalSukses++
				} else {
					totalSukses = 0
				}
			}
		}

		fmt.Printf("\nw1 = %.3f\nw2 = %.3f\nw3 = %.3f\nepoch = %d\n", AND.weight[0], AND.weight[1], AND.weight[2], AND.no/4)

		// Pengujian model pada data uji
		fmt.Println("\nHasil Pengujian pada Data Uji:")
		fmt.Printf("%-3s%-3s%-8s%-8s%-7s%-7s\n", "x1", "x2", "target", "output", "Benar?", "Akurasi")
		totalBenar := 0
		for i := 0; i < 4; i++ {
			x1 := input[i][0]
			x2 := input[i][1]
			target := targetOutput[i]

			output := AND.uji(x1, x2)
			benar := output == target
			if benar {
				totalBenar++
			}
			fmt.Printf("%-3d%-3d%-8d%-8d%-7t%-7.2f%%\n", x1, x2, target, output, benar, float64(totalBenar)/float64(i+1)*100)
		}
	}
}
