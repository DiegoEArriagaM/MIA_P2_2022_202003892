package Sistema_Archivos

import (
	"Backend/Structs"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

var IdUsuario = " "
var NameUsuario = " "
var PassUsuario = " "
var GruopUsuario = " "
var CambioCont = false
var sb Structs.SuperBloque
var file *os.File
var errf error

func logout() Structs.Resp {
	if UsuarioL.IdU == 0 {
		return Structs.Resp{Res: "NO SE HA INICIADO UNA SESION CON ANTERIORIDAD"}
	}
	UsuarioL.NombreU = " "
	UsuarioL.IdMount = " "
	UsuarioL.IdU = 0
	UsuarioL.IdG = 0
	UsuarioL.Login = false
	return Structs.Resp{Res: "SE CERRO LA SESION"}
}

func login() Structs.Resp {
	defer func() {
		IdUsuario = " "
		NameUsuario = " "
		PassUsuario = " "
		GruopUsuario = " "
		sb = Structs.SuperBloque{}
		file = nil
	}()

	if UsuarioL.IdU != 0 {
		return Structs.Resp{Res: "YA HAY UNA SESION ACTIVA"}
	}

	nodo := Mlist.buscar(IdUsuario)
	if nodo != nil {

		file, errf = os.OpenFile(nodo.Path, os.O_RDWR, 0777)
		if errf == nil {
			banderaU := false
			banderaG := false
			if nodo.Type == 'p' {
				mbr := Structs.MBR{}
				file.Seek(0, 0)
				errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(mbr))), binary.BigEndian, &mbr)
				if mbr.Mbr_partition[nodo.Pos].Part_status != '2' {
					return Structs.Resp{Res: "NO SE HA FORMATEADO LA PARTICION"}
				}
				file.Seek(int64(nodo.Start), 0)
				errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(sb))), binary.BigEndian, &sb)
			} else if nodo.Type == 'l' {
				ebr := Structs.EBR{}
				file.Seek(int64(nodo.Start), 0)
				errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(ebr))), binary.BigEndian, &ebr)
				if ebr.Part_status != '2' {
					return Structs.Resp{Res: "NO SE HA FORMATEADO LA PARTICION"}
				}
				file.Seek(int64(nodo.Start+int(unsafe.Sizeof(Structs.EBR{}))), 0)
				errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(sb))), binary.BigEndian, &sb)
			}
			content := getConten(int(sb.S_inode_start) + int(unsafe.Sizeof(Structs.TablaInodo{})))
			usuarios := splitUsr(content)
			grupos := splitGrp(content)
			var datosU []string
			var datosG []string
			for i := 0; i < len(usuarios); i++ {
				datosU = strings.Split(usuarios[i], ",")
				if datosU[3] == NameUsuario {
					banderaU = true
					for j := 0; j < len(usuarios); j++ {
						datosG = strings.Split(grupos[j], ",")
						if datosG[2] == datosU[2] {
							banderaG = true
							goto t0
						}
					}
				}
			}
		t0:
			fmt.Println(usuarios, grupos)
			if banderaU {
				if banderaG {
					UsuarioL.NombreU = datosU[3]
					UsuarioL.IdMount = IdUsuario
					IdU, _ := strconv.Atoi(datosU[0])
					UsuarioL.IdU = int32(IdU)
					IdG, _ := strconv.Atoi(datosG[0])
					UsuarioL.IdG = int32(IdG)
					UsuarioL.Login = true
					file.Close()
					return Structs.Resp{Res: "SE INICIO SESION COMO " + UsuarioL.NombreU}
				}
				file.Close()
				return Structs.Resp{Res: "GRUPO NO ENCONTRADO"}
			}
			file.Close()
			return Structs.Resp{Res: "USUARIO NO ENCONTRADO"}
		}
		return Structs.Resp{Res: "DISCO INEXISTENTE"}
	}
	return Structs.Resp{Res: "NO SE HA ENCONTRADO ALGUNA MONTURA CON EL ID: " + IdUsuario}
}

func getConten(inodoStart int) string {
	var inodo Structs.TablaInodo
	var archivo Structs.BloqueArchivo
	file.Seek(int64(inodoStart), 0)
	errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(inodo))), binary.BigEndian, &inodo)
	content := ""
	for i := 0; i < 16; i++ {
		if inodo.I_block[i] != -1 {
			file.Seek(int64(inodo.I_block[i]), 0)
			errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(archivo))), binary.BigEndian, &archivo)
			content += archivoContent2(archivo.B_content)
		}
	}
	return content
}

func splitUsr(cadena string) []string {
	var split []string
	content := strings.Split(cadena, "\n")
	for i := 0; i < len(content); i++ {
		if content[i] != "" {
			datos := strings.Split(content[i], ",")
			if datos[1] == "U" && datos[0] != "0" {
				split = append(split, content[i])
			}
		}

	}
	return split
}

func splitGrp(cadena string) []string {
	var split []string
	content := strings.Split(cadena, "\n")
	for i := 0; i < len(content); i++ {
		if content[i] != "" {
			datos := strings.Split(content[i], ",")
			if datos[1] == "G" && datos[0] != "0" {
				split = append(split, content[i])
			}
		}
	}
	return split
}
