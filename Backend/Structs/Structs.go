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

type SuperBloque struct {
	s_filesystem_type   int32
	s_inodes_count      int32
	s_blocks_count      int32
	s_free_blocks_count int32
	s_free_inodes_count int32
	s_mtime             int64
	s_umtime            int64
	s_mnt_count         int32
	s_magic             int32
	s_inode_s           int32
	s_block_s           int32
	s_firts_ino         int32
	s_first_blo         int32
	s_bm_inode_start    int32
	s_bm_block_start    int32
	s_inode_start       int32
	s_block_start       int32
}
type TablaInodo struct {
	i_uid   int32
	i_gid   int32
	i_s     int32
	i_atime int64
	i_ctime int64
	i_mtime int64
	i_block [15]int32
	i_type  byte
	i_perm  int32
}

type Content struct {
	b_name  [12]byte
	b_inodo int32
}

type BloqueCarpeta struct {
	b_content [4]Content
}
type BloqueArchivo struct {
	b_content [64]byte
}
