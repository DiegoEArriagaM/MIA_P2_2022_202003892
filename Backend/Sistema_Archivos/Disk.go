package Sistema_Archivos

import (
	"Backend/Structs"
)

var Sdisk = 0
var Fdisk = "bf"
var Udisk = "m"
var Pdisk = " "
var Directorio_disk = ""

func mkdisk() Structs.Resp {

	return Structs.Resp{Res: "", Reporte: false}
}
