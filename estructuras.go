package estructuras

import "fmt"



type Imagen struct {
	Titulo string
	Formato string
	Canales string
}

func (i *Imagen) Mostrar() {  //metodo de struct
	fmt.Println("Titulo: ", i.Titulo, "Formato: ", i.Formato, "Canales: " , i.Canales)
}


type Audio struct {
	Titulo string
	Formato string
	Duracion float64
}

func (a *Audio) Mostrar() {  //metodo de struct
	fmt.Println("Titulo: ", a.Titulo, "Formato: ", a.Formato, "Duracion: " , a.Duracion)
}


type Video struct {
	Titulo string
	Formato string
	Frames string
}

func (v *Video) Mostrar() {  //metodo de struct
	fmt.Println("Titulo: ", v.Titulo, "Formato: ", v.Formato, "Frames: " , v.Frames)
}