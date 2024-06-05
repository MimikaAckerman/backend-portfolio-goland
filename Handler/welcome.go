package Handler

import (
	"fmt"
    "net/http"
)
func Welcome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Bienvenido a la API de mi portfolio!")
}