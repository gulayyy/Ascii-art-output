package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Output dosyasını flag olarak tanımla
	var outputFileName string // Bu satırda, outputFileName adında bir değişken tanımlanır. Bu değişken, komut satırından alınan dosya adını saklayacak.
	flag.StringVar(&outputFileName, "output", "", "Output file name")
	flag.Parse() // Bu fonksiyon, komut satırı argümanlarını işler ve tanımlanan bayrakların değerlerini belirlenmiş değişkenlere atar.

	// Dosya adı boş ise kullanıcıya hata mesajı göster ve çık
	if outputFileName == "" {
		fmt.Println("Hata: --output=<fileName.txt> bayrağı ile bir dosya adı belirtmelisiniz.")
		return
	}

	// Kelimeyi komut satırından al
	word := os.Args[2]  // İkinci argümanı word değişkenine al
	words := os.Args[3] // Üçüncü argümanı words değişkenine al

	// \n kaçış dizisini gerçek satır sonu karakterine dönüştür
	word = strings.ReplaceAll(word, "\\n", "\n")

	// Dosya içeriğini oku
	var fileContent string // Dosyanın içeriğini tutacak değişken
	if words == "standard" {
		file, err := os.ReadFile("standard.txt") // "standard.txt" dosyasını oku
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file) // Dosya içeriğini string formatına çevir ve fileContent'e ata
	} else if words == "shadow" {
		file, err := os.ReadFile("shadow.txt") // "shadow.txt" dosyasını oku
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file) // Dosya içeriğini string formatına çevir ve fileContent'e ata
	} else if words == "thinkertoy" {
		file, err := os.ReadFile("thinkertoy.txt") // "thinkertoy.txt" dosyasını oku
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file) // Dosya içeriğini string formatına çevir ve fileContent'e ata
	} else {
		fmt.Println("Geçersiz kelime grubu:", words)
		return
	}

	// Dosya içeriğini satırlara ayır
	lines := strings.Split(fileContent, "\n") // Dosya içeriğini satırlara ayır ve lines değişkenine ata

	// Dosyayı oluştur veya aç, dosya adında bir dosya nesnesi dön
	dosya, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer dosya.Close() // Fonksiyon sonunda dosyayı kapat

	// Her kelime için ASCII sanatını yazdır
	for i, line := range strings.Split(word, "\n") { // word dizesini satırlara ayır ve her satırı line değişkenine al
		if line == "" { // Eğer satır boşsa
			if i != 0 { // Eğer bu boş satır dosyanın en üstündeki boşluklardan değilse
				_, err := dosya.WriteString("\n") // Dosyaya yeni bir satır ekle
				if err != nil {
					log.Fatal(err)
				}
			}
			continue // Bir sonraki satıra geç
		}

		// Her bir karakteri işleyerek ASCII sanatını oluştur
		for h := 1; h < 9; h++ { // Her karakter için 9 satır oluştur
			for _, char := range line { // Her karakteri al
				printAsciiArtForCharacter(char, h, lines, dosya) // Karakterin ASCII sanatını oluştur ve dosyaya yaz
			}
			_, err := dosya.WriteString("\n") // Bir satır tamamlandığında yeni bir satır ekle
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	_, err = dosya.WriteString("\n") // Dosyaya bir satır daha ekle
	if err != nil {
		log.Fatal(err)
	}
}

func printAsciiArtForCharacter(char rune, lineIndex int, lines []string, dosya *os.File) {
	// ASCII karakterinin indeksini hesapla
	index := (int(char) - 32) * 9 // Karakterin ASCII kodunu al, 32 çıkar ve 9 ile çarp

	// ASCII sanatını yazdır
	if index >= 0 && index+8 <= len(lines) { // Eğer ASCII karakterinin indeksi geçerliyse
		_, err := dosya.WriteString(lines[index+lineIndex]) // Dosyaya ASCII sanatını yaz
		if err != nil {
			log.Fatal(err)
		}
	}
}
