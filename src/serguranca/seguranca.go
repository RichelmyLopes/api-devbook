package serguranca

import "golang.org/x/crypto/bcrypt"

// recebe string e coloca hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// Comprara sneha com hash e valida se sao iguais
func VerificarSenha(senha string, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senha), []byte(senhaHash))
}
