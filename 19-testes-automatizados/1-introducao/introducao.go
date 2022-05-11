package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(TipoDeEndereco)
}

//go mod init introducao-testes para conseguir utilizar as dependencias dos pacotes
