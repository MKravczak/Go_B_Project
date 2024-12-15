package main // Pakiet główny

import (
	//"Go_Course/internal/env"
	"github.com/joho/godotenv"
	"log"
	"os"
) // Importowanie pakietu logowania

func main() { // Funkcja główna
	err := godotenv.Load()
	if err != nil {
		log.Println("Brak pliku .env lub nie można go załadować, kontynuacja bez .env")
	}
	cfg := config{ // Tworzenie nowej konfiguracji
		addr: os.Getenv("ADDR"), // Ustawianie adresu serwera
	}

	app := &application{ // Tworzenie nowej aplikacji
		config: cfg, // Przypisywanie konfiguracji do aplikacji
	}
	mux := app.mount()      // Montowanie routera
	log.Fatal(app.run(mux)) // Uruchamianie serwera i logowanie błędów
}
