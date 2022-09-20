package Structs

import "time"

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
	part_status byte
	part_type   byte
	part_fit    byte
	part_start  int
	part_s      int
	part_name   [16]byte
}

type MBR struct {
	mbr_tamanio        int
	time_t             time.Time
	mbr_disk_signature int
	disk_fit           byte
	mbr_partition      [4]Partition
}

type EBR struct {
	part_status byte
	part_fit    byte
	part_start  int
	part_s      int
	part_next   int
	part_name   [16]byte
}
