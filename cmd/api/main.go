package main // Pakiet główny

import "log" // Importowanie pakietu logowania

func main() { // Funkcja główna
	cfg := config{ // Tworzenie nowej konfiguracji
		addr: ":8080", // Ustawianie adresu serwera
	}

	app := &application{ // Tworzenie nowej aplikacji
		config: cfg, // Przypisywanie konfiguracji do aplikacji
	}
	mux := app.mount()      // Montowanie routera
	log.Fatal(app.run(mux)) // Uruchamianie serwera i logowanie błędów
}
