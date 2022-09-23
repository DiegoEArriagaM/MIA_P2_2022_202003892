package Sistema_Archivos

import (
	"Backend/Structs"
)

var Spart = 0
var Upart = 'k'
var Ppart = " "
var Tpart = 'p'
var Fpart = 'w'
var Namepart = " "
var Banderapart = "n"

func fdisk() Structs.Resp {
	return Structs.Resp{Res: "Fdisk"}
}
