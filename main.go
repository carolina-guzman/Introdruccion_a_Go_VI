package main

import (
	"fmt"

	"./estructuras"
)

func menu() {

}

func main() {
	op := 1
	files := estructuras.ContenidoWeb{}
	for op != 0 {
		fmt.Println("-----Menu------\n1. Agregar imagen\n2.Agregar audio\n3.Agregar video\n4. Imprimir todo\n0. Salir ")
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println("---------Agregar imagen---------")
			im := estructuras.Imagen{}
			fmt.Println("Titulo: ")
			fmt.Scan(&im.Titulo)
			fmt.Println("Formato: ")
			fmt.Scan(&im.Formato)
			fmt.Println("Canal: ")
			fmt.Scan(&im.Canales)
			files.Archivos = append(files.Archivos, &im)
		case 2:
			fmt.Println("---------Agregar audio---------")
		case 3:
			fmt.Println("---------Agregar video---------")
		case 4:
			fmt.Println("-------------SLICE-------------")
			files.Mostrar()
		}

	}
}
