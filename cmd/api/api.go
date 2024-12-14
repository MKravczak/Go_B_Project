package main // Pakiet główny

import ( // Importowanie pakietów
	"github.com/go-chi/chi/v5"            // Importowanie chi routera
	"github.com/go-chi/chi/v5/middleware" // Importowanie middleware z chi
	"log"                                 // Importowanie pakietu logowania
	"net/http"                            // Importowanie pakietu HTTP
	"time"                                // Importowanie pakietu czasu
)

type application struct { // Definicja struktury application
	config config // Pole config typu config
}
type config struct { // Definicja struktury config
	addr string // Pole addr typu string
}

func (app *application) mount() http.Handler { // Metoda mount zwracająca http.Handler
	r := chi.NewRouter()        // Tworzenie nowego routera chi
	r.Use(middleware.RequestID) // Używanie middleware do generowania unikalnych identyfikatorów żądań
	r.Use(middleware.RealIP)    // Używanie middleware do uzyskiwania rzeczywistego adresu IP klienta
	r.Use(middleware.Logger)    // Używanie middleware do logowania żądań HTTP
	r.Use(middleware.Recoverer) // Używanie middleware do odzyskiwania po panice

	// Ustawia limit czasu na kontekście żądania (ctx), który sygnalizuje przez ctx.Done()
	// że żądanie przekroczyło limit czasu i dalsze przetwarzanie powinno zostać zatrzymane.
	r.Use(middleware.Timeout(60 * time.Second)) // Używanie middleware do ustawiania limitu czasu na żądania

	r.Route("/v1", func(r chi.Router) { // Definiowanie trasy dla wersji API v1
		r.Get("/health", app.healthCheckHandler) // Obsługa żądań GET na endpoint /health
	})
	return r // Zwracanie routera
}

func (app *application) run(mux http.Handler) error { // Metoda run uruchamiająca serwer
	srv := &http.Server{ // Tworzenie nowego serwera HTTP
		Addr:         app.config.addr,  // Ustawianie adresu serwera
		Handler:      mux,              // Ustawianie handlera serwera
		WriteTimeout: time.Second * 30, // Ustawianie limitu czasu na zapis
		ReadTimeout:  time.Second * 10, // Ustawianie limitu czasu na odczyt
		IdleTimeout:  time.Minute,      // Ustawianie limitu czasu na bezczynność
	}
	start := time.Now()

	log.Printf("Starting server at %s", app.config.addr)
	elapsed := time.Since(start)
	log.Printf("Server started after %s seconds", elapsed)

	return srv.ListenAndServe() // Uruchamianie serwera
}
