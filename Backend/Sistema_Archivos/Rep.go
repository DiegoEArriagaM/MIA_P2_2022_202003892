package Sistema_Archivos

import "Backend/Structs"

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
	return Structs.Resp{Res: "disk"}
}

func tree() Structs.Resp {
	return Structs.Resp{Res: "tree"}
}

func fileR() Structs.Resp {
	return Structs.Resp{Res: "file"}
}

func sbR() Structs.Resp {
	return Structs.Resp{Res: "sb"}
}
