package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/RufoBernedo/twittor/bd"
	"github.com/RufoBernedo/twittor/models"
)

/*Email valor de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken procesa token para extraer sus valores*/
func ProcesoToken(token string) (*models.Claim, bool, string, error) {
	miClave := []byte("LosAngelesLakers")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido") //no incluir tildes ni simbolos en mensajes al crear error, da error de compilación
	}
	token = strings.TrimSpace(splitToken[1])

	//validación del token
	tkn, err := jwt.ParseWithClaims(token, claims, func(tok *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
