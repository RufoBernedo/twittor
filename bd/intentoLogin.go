package bd

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/RufoBernedo/twittor/models"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usuario.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usuario, false
	}
	return usuario, true
}
