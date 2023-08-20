package controllers

import "net/http"

// Cria um usuario no bd
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))
}

// Buscar todos usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}

// buscar usuario especifico pelo id
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario"))
}

// update em usuario por id
func Atualizar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

// delete usuario pelo id
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuario"))
}
