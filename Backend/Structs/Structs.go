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

// Structs del Sistema de Archivos
type Partition struct {
	Part_status byte
	Part_type   byte
	Part_fit    byte
	Part_start  int32
	Part_s      int32
	Part_name   [16]byte
}

type MBR struct {
	Mbr_tamanio        int32
	Mbr_fecha_creacion int64
	Mbr_disk_signature int32
	Disk_fit           byte
	Mbr_partition      [4]Partition
}

type EBR struct {
	Part_status byte
	Part_fit    byte
	Part_start  int32
	Part_s      int32
	Part_next   int32
	Part_name   [16]byte
}