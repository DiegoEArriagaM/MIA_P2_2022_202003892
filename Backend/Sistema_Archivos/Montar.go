package Sistema_Archivos

import (
	"Backend/Structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
	"unsafe"
)

var Pmontar = " "
var Namemontar = " "
var IdMontar = " "
var Tmontar = "full"

func mount() Structs.Resp {
	defer func() {
		Pmontar = " "
		Namemontar = " "
		IdMontar = " "
		Tmontar = "full"
	}()

	if Pmontar != " " {
		if Namemontar != " " {
			pos := -1
			file, errf := os.OpenFile(Pmontar, os.O_RDWR, 0777)
			if errf == nil {
				file.Seek(0, 0)
				mbr := Structs.MBR{}
				errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(mbr))), binary.BigEndian, &mbr)
				for i := 0; i < 4; i++ {
					name1 := string(mbr.Mbr_partition[i].Part_name[:])
					if strncmp(name1, Namemontar) {
						pos = i
						break
					} else if mbr.Mbr_partition[i].Part_type == 'e' {
						ebr := Structs.EBR{}
						sb := Structs.SuperBloque{}
						file.Seek(int64(mbr.Mbr_partition[i].Part_start), 0)
						errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(ebr))), binary.BigEndian, &ebr)
						if errf != nil {
							fmt.Println(errf)
						}
						if ebr.Part_next != -1 || ebr.Part_s != -1 {
							name1 = string(ebr.Part_name[:])
							if strncmp(name1, Namemontar) {
								if ebr.Part_status == '0' || ebr.Part_status == '1' {
									ebr.Part_status = '1'
								}

								file.Seek(int64(ebr.Part_start), 0)
								var bufferEBRN bytes.Buffer
								errf = binary.Write(&bufferEBRN, binary.BigEndian, ebr)
								EscribirFile(file, bufferEBRN.Bytes())

								if ebr.Part_status == '2' {
									file.Seek(int64(ebr.Part_start)+int64(unsafe.Sizeof(Structs.EBR{})), 0)
									errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(sb))), binary.BigEndian, &sb)
									sb.S_mtime = time.Now().Unix()
									sb.S_mnt_count += 1
									file.Seek(int64(ebr.Part_start)+int64(unsafe.Sizeof(Structs.EBR{})), 0)
									var bufferSB bytes.Buffer
									errf = binary.Write(&bufferSB, binary.BigEndian, sb)
									EscribirFile(file, bufferSB.Bytes())
								}
								file.Close()
								return Mlist.add(Pmontar, Namemontar, 'l', int(ebr.Part_start), -1)

							} else if ebr.Part_next != -1 {
								file.Seek(int64(ebr.Part_next), 0)
								errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(ebr))), binary.BigEndian, &ebr)
								for true {
									name1 = string(ebr.Part_name[:])
									if strncmp(name1, Namemontar) {
										if ebr.Part_status == '0' || ebr.Part_status == '1' {
											ebr.Part_status = '1'
										}

										file.Seek(int64(ebr.Part_start), 0)
										var bufferEBRN bytes.Buffer
										errf = binary.Write(&bufferEBRN, binary.BigEndian, ebr)
										EscribirFile(file, bufferEBRN.Bytes())

										if ebr.Part_status == '2' {
											file.Seek(int64(ebr.Part_start)+int64(unsafe.Sizeof(Structs.EBR{})), 0)
											errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(sb))), binary.BigEndian, &sb)
											sb.S_mtime = time.Now().Unix()
											sb.S_mnt_count += 1
											file.Seek(int64(ebr.Part_start)+int64(unsafe.Sizeof(Structs.EBR{})), 0)
											var bufferSB bytes.Buffer
											errf = binary.Write(&bufferSB, binary.BigEndian, sb)
											EscribirFile(file, bufferSB.Bytes())
										}
										file.Close()
										return Mlist.add(Pmontar, Namemontar, 'l', int(ebr.Part_start), -1)

									}

									if ebr.Part_next == -1 {
										break
									}
									file.Seek(int64(ebr.Part_next), 0)
									errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(ebr))), binary.BigEndian, &ebr)
								}
							}
						}
					}
				}

				if pos != -1 {
					if mbr.Mbr_partition[pos].Part_type == 'e' {
						file.Close()
						return Structs.Resp{Res: "NO SE PUEDE MONTAR UNA PARTICION EXTENDIDA "}
					}
					sb := Structs.SuperBloque{}
					if mbr.Mbr_partition[pos].Part_status == '0' || mbr.Mbr_partition[pos].Part_status == '1' {
						mbr.Mbr_partition[pos].Part_status = '1'
					}

					file.Seek(0, 0)
					var bufferMBR bytes.Buffer
					errf = binary.Write(&bufferMBR, binary.BigEndian, mbr)
					EscribirFile(file, bufferMBR.Bytes())

					if mbr.Mbr_partition[pos].Part_status == '2' {
						file.Seek(int64(mbr.Mbr_partition[pos].Part_start), 0)
						errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(sb))), binary.BigEndian, &sb)
						sb.S_mtime = time.Now().Unix()
						sb.S_mnt_count += 1
						file.Seek(int64(mbr.Mbr_partition[pos].Part_start), 0)
						var bufferSB bytes.Buffer
						errf = binary.Write(&bufferSB, binary.BigEndian, sb)
						EscribirFile(file, bufferSB.Bytes())
					}
					file.Close()
					return Mlist.add(Pmontar, Namemontar, 'l', int(mbr.Mbr_partition[pos].Part_start), -1)

				}
				file.Close()
				return Structs.Resp{Res: "NO EXISTE ESA PARTICION"}
			}
			return Structs.Resp{Res: "DISCO INEXISTENTE"}
		}
		return Structs.Resp{Res: "ASEGURESE DE ESCRIBIR EL NOMBRE DE LA PARTICION"}
	}
	return Structs.Resp{Res: "ASEGURESE DE ESCRIBIR UN RUTA"}
}
