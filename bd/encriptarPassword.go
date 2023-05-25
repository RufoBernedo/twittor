package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/*Recibe un password y lo retorna encriptado*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo) // a GenerateFromPassword se le pasa el pass casteado a un slice de bytes
	return string(bytes), err
}
