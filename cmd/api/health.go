package main // Pakiet główny

import ( // Importowanie pakietów
	"net/http" // Importowanie pakietu HTTP
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) { // Metoda healthCheckHandler
	w.Write([]byte("Health's OK")) // Zapis odpowiedzi
}
