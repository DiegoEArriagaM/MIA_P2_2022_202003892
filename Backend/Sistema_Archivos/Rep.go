package Sistema_Archivos

import (
	"Backend/Structs"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unsafe"
)

var Prep = " "
var Namerep = " "
var Idrep = " "
var Rutarep = " "
var Dirrep = " "
var Extrep = " "

func GenerateRep() Structs.Resp {
	if Prep != " " {
		if Idrep != " " {
			if Namerep == "disk" {
				return disk()
			} else if Namerep == "tree" {
				return tree()
			} else if Namerep == "file" {
				return fileR()
			} else if Namerep == "sb" {
				return sbR()
			}
			return Structs.Resp{Res: "NOMBRE DE REPORTE INVALIDO"}
		}
		return Structs.Resp{Res: "FALTA EL ID DE LA PARTICION"}
	}
	return Structs.Resp{Res: "FALTA LA UBICACION DONDE SE GUARDARA EL REPORTE"}
}

func disk() Structs.Resp {
	nodo := Mlist.buscar(Idrep)
	if nodo != nil {
		Dirrep = GetDirectorio(Prep)
		Extrep = GetExtension(Prep)
		nombreD := nombre(Prep)
		err := os.MkdirAll(Dirrep, 0777)
		if err != nil {
			fmt.Printf("%s", err)
		}
		file, errf := os.OpenFile(nodo.Path, os.O_RDWR, 0777)
		if errf == nil {
			var mbr Structs.MBR
			file.Seek(0, 0)
			errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(mbr))), binary.BigEndian, &mbr)
			tamanioT := int(mbr.Mbr_tamanio)
			dotS := ""
			dot, errD := os.OpenFile("Reportes/disk.dot", os.O_CREATE, 0777)
			dot.Close()
			if errD != nil {
				fmt.Println(errD)
			}

			dotS += "digraph G {\n"
			dotS += "node[shape=none]\n"
			dotS += "start[label=<<table><tr>"
			dotS += "<td rowspan=\"2\">MBR</td>"

			i := 0
			inicio := int(unsafe.Sizeof(Structs.MBR{}))
			for i < 4 {
				if mbr.Mbr_partition[i].Part_start != -1 {
					if mbr.Mbr_partition[i].Part_type == 'p' {
						var p1 = float32(mbr.Mbr_partition[i].Part_s) / float32(tamanioT)

					}
				} else {

				}
				i++
			}
			dotS += "</tr></table>>];\n"
			dotS += "}"
			errD = os.WriteFile("Reportes/disk.dot", []byte(dotS), 0777)
			if errD != nil {
				fmt.Println(errD)
			}

			file.Close()
			_, errD = exec.Command("dot", "-T"+Extrep, "Reportes/disk.dot", "-o", "Reportes/"+nombreD).Output()
			if errD != nil {
				fmt.Printf("%s", errD)
			}
			_, errD = exec.Command("dot", "-T"+Extrep, "Reportes/disk.dot", "-o", Dirrep+nombreD).Output()
			if errD != nil {
				fmt.Printf("%s", errD)
			}
			return Structs.Resp{Res: "SE GENERO EL REPORTE DISK"}
		}
		return Structs.Resp{Res: "DISCO INEXISTENTE"}
	}
	return Structs.Resp{Res: "NO SE HA ENCONTRADO ALGUNA MONTURA CON EL ID: " + Idrep}

}

func tree() Structs.Resp {
	nodo := Mlist.buscar(Idrep)
	if nodo != nil {

	}
	return Structs.Resp{Res: "NO SE HA ENCONTRADO ALGUNA MONTURA CON EL ID: " + Idrep}
}

func fileR() Structs.Resp {
	nodo := Mlist.buscar(Idrep)
	if nodo != nil {

	}
	return Structs.Resp{Res: "NO SE HA ENCONTRADO ALGUNA MONTURA CON EL ID: " + Idrep}
}

func sbR() Structs.Resp {
	nodo := Mlist.buscar(Idrep)
	if nodo != nil {

	}
	return Structs.Resp{Res: "NO SE HA ENCONTRADO ALGUNA MONTURA CON EL ID: " + Idrep}
}

func nombre(path string) string {
	directorio := ""
	aux := path
	p := strings.Index(aux, "/")
	for p != -1 {
		directorio += aux[:p] + "/"
		aux = aux[p+1:]
		p = strings.Index(aux, "/")
	}
	i := find(aux, ".")

	return aux[:i]
}
