package Sistema_Archivos

import (
	"Backend/Structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

var Sdisk = 0
var Fdisk = "bf"
var Udisk = "m"
var Pdisk = " "
var Directorio_disk = ""

func mkdisk() Structs.Resp {
	val := validar()
	if !val.Val {
		return Structs.Resp{Res: val.Men, Reporte: false}
	}

	var file *os.File
	var errf error
	file, errf = os.OpenFile(Pdisk, os.O_RDWR, 0666)
	defer func() {
		reco := recover()
		if reco != nil {
			fmt.Println(reco)
		}
		Sdisk = 0
		Fdisk = "bf"
		Udisk = "m"
		Pdisk = " "
		Directorio_disk = ""
		if file != nil {
			file.Close()
		}

	}()
	if errf == nil {
		return Structs.Resp{Res: "EL DISCO YA EXISTE", Reporte: false}
	}

	Directorio_disk = getDirectorio(Pdisk)
	err := os.MkdirAll(Directorio_disk, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	}
	/*out, err := exec.Command("mkdir", Directorio_disk, " -p").Output()
	fmt.Println(string(out[:]))
	if err != nil {
		fmt.Printf("%s", err)
	}
	out, err = exec.Command("chmod", "-R ", "777", Directorio_disk).Output()
	fmt.Println(string(out[:]))
	if err != nil {
		fmt.Printf("%s", err)
	}*/

	size := Sdisk
	if Udisk == "m" {
		size = size * 1024
	}

	file, errf = os.OpenFile(Pdisk, os.O_RDWR|os.O_CREATE, 0777)

	var contenedor bytes.Buffer
	var buffer [1024]int8
	for i := 0; i < 1024; i++ {
		buffer[i] = 0
	}

	binary.Write(&contenedor, binary.BigEndian, &buffer)

	for i := 0; i < size; i++ {
		var bufferControl bytes.Buffer
		binary.Write(&bufferControl, binary.BigEndian, contenedor.Bytes())
		_, err = file.Write(bufferControl.Bytes())
		if err != nil {
			fmt.Println(err)
		}
	}

	return Structs.Resp{Res: "SE CREO EL DISCO EXITOSAMENTE", Reporte: false}
}

func validar() Structs.Bandera {
	if Sdisk > 0 {
		if strncmp(Fdisk, "bf") || strncmp(Fdisk, "ff") || strncmp(Fdisk, "wf") {
			if strncmp(Udisk, "k") || strncmp(Udisk, "m") {
				if Pdisk != " " {
					i := find(Pdisk, ".")
					extension := Pdisk[i+1:]
					if strncmp(extension, "dsk") {
						return Structs.Bandera{Val: true, Men: ""}
					} else {
						return Structs.Bandera{Val: false, Men: "EXTENSION INCORRECTA"}
					}
				} else {
					return Structs.Bandera{Val: false, Men: "ASEGURESE DE ESCRIBIR UN RUTA"}
				}
			} else {
				return Structs.Bandera{Val: false, Men: "CONFIGURACION DE UNIDADES DEL TAMAÑO DE MEMORIA INVALIDO"}
			}
		} else {
			return Structs.Bandera{Val: false, Men: "CONFIGURACION DE AJUSTE INVALIDO"}
		}
	} else {
		return Structs.Bandera{Val: false, Men: "EL TAMAÑO DEL DISCO TIENE QUE SER MAYOR A 0"}
	}
}

func getDirectorio(path string) string {
	directorio := ""
	aux := path
	p := strings.Index(aux, "/")
	for p != -1 {
		directorio += aux[:p] + "/"
		aux = aux[p+1:]
		p = strings.Index(aux, "/")
	}

	return directorio
}
