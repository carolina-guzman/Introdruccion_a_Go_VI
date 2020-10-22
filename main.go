package main

import (
	"fmt"

	"./estructuras"
)

func menu() {

}

func main() {
	op := 9
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
			au := estructuras.Audio{}
			fmt.Println("Titulo: ")
			fmt.Scan(&au.Titulo)
			fmt.Println("Formato: ")
			fmt.Scan(&au.Formato)
			fmt.Println("Duración: ")
			fmt.Scan(&au.Duracion)
			files.Archivos = append(files.Archivos, &au)
		case 3:
			fmt.Println("---------Agregar video---------")
			vi := estructuras.Video{}
			fmt.Println("Titulo: ")
			fmt.Scan(&vi.Titulo)
			fmt.Println("Formato: ")
			fmt.Scan(&vi.Formato)
			fmt.Println("Frames: ")
			fmt.Scan(&vi.Frames)
			files.Archivos = append(files.Archivos, &vi)
		case 4:
			fmt.Println("-------------SLICE-------------")
			files.Mostrar()
		case 0:
			fmt.Println("Adios!")
		default:
			fmt.Println("error")
		}

	}
}
