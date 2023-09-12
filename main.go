package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "jdan" {
		//123foi
		return "$1$P8o6w3Xn$QZZnoY/ktrSFRIsneLAEZ/"
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Por favor, inserir 'Path' 'Port'")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	httpPort := os.Args[2]
	// fs := http.FileServer(http.Dir(httpDir))

	authentication := auth.NewBasicAuthenticator("meuserver.local", Secret)
	http.HandleFunc("/", authentication.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))
	fmt.Printf("Testando servidor go porta: %s\n", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
