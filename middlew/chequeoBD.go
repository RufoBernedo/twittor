package middlew

import (
	"net/http"

	"github.com/RufoBernedo/twittor/bd"
)

/*ChequeoBD es el middleware que me permite saber el estado de la BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
