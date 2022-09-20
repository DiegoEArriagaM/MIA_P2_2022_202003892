package Structs

// Structs de administrador
type Resp struct {
	Res     string `json:"res"`
	Reporte bool   `json:"reporte"`
	idRep   string `json:"id_rep"`
}

type Bandera struct {
	Val bool
	Men string
}

type Inicio struct {
	Res string `json:"res"`
}

type Comando struct {
	Command string `json:"comando"`
}
