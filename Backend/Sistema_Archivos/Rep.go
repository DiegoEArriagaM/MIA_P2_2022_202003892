package Sistema_Archivos

import (
	"Backend/Structs"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

var Prep = " "
var Namerep = " "
var Idrep = " "
var Rutarep = " "
var Dirrep = " "
var Extrep = " "

func GenerateRep() Structs.Resp {
	defer func() {
		Prep = " "
		Namerep = " "
		Idrep = " "
		Rutarep = " "
		Dirrep = " "
		Extrep = " "
	}()
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
			dot, errD := os.OpenFile("Reportes/"+nombreD+".dot", os.O_CREATE, 0777)
			dot.Close()
			if errD != nil {
				fmt.Println(errD)
			}

			dotS += "digraph G {\n"
			dotS += "node[shape=none]\n"
			dotS += "start[label=<<table><tr>\n"
			dotS += "<td rowspan=\"2\">MBR</td>\n"

			i := 0
			inicio := int(unsafe.Sizeof(Structs.MBR{}))
			for i < 4 {
				if mbr.Mbr_partition[i].Part_start != -1 {
					if mbr.Mbr_partition[i].Part_type == 'p' {
						porcentaje := (float64(mbr.Mbr_partition[i].Part_s) / float64(tamanioT)) * 100
						trunc := float64(int(porcentaje*100)) / 100
						name1 := getPartName(mbr.Mbr_partition[i].Part_name)
						dotS += "<td rowspan=\"2\">" + name1 + " <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
						if i != 3 {
							if (mbr.Mbr_partition[i].Part_start + mbr.Mbr_partition[i].Part_s) < mbr.Mbr_partition[i+1].Part_start {
								porcentaje = (float64(mbr.Mbr_partition[i+1].Part_start-(mbr.Mbr_partition[i].Part_start+mbr.Mbr_partition[i].Part_s)) / float64(tamanioT)) * 100
								trunc = float64(int(porcentaje*100)) / 100
								dotS += "<td rowspan=\"2\">LIBRE <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
							}
						} else if int(mbr.Mbr_partition[i].Part_start+mbr.Mbr_partition[i].Part_s) < tamanioT {
							porcentaje = (float64(tamanioT-int(mbr.Mbr_partition[i].Part_start+mbr.Mbr_partition[i].Part_s)) / float64(tamanioT)) * 100
							trunc = float64(int(porcentaje*100)) / 100
							dotS += "<td rowspan=\"2\">LIBRE <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
						}
					} else if mbr.Mbr_partition[i].Part_type == 'e' {
						porcentaje := (float64(mbr.Mbr_partition[i].Part_s) / float64(tamanioT)) * 100
						dotS += "<td rowspan=\"2\">EXTENDIDA</td>\n"
						ebr := Structs.EBR{}
						file.Seek(int64(mbr.Mbr_partition[i].Part_start), 0)
						errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(ebr))), binary.BigEndian, &ebr)
						if !(ebr.Part_s == -1 && ebr.Part_next == -1) {
							if ebr.Part_s > -1 {
								name1 := getPartName(ebr.Part_name)
								dotS += "<td rowspan=\"2\">EBR <br/>" + name1 + "</td>\n"
								porcentaje = (float64(ebr.Part_s) / float64(tamanioT)) * 100.0
								trunc := float64(int(porcentaje*100)) / 100
								dotS += "<td rowspan=\"2\">Logica <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
							} else {
								dotS += "<td rowspan=\"2\">EBR</td>\n"
								porcentaje = ((float64(ebr.Part_next - ebr.Part_start)) / float64(tamanioT)) * 100.0
								trunc := float64(int(porcentaje*100)) / 100
								dotS += "<td rowspan=\"2\">Libre <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
							}
							if ebr.Part_next != -1 {
								file.Seek(int64(ebr.Part_next), 0)
								errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(ebr))), binary.BigEndian, &ebr)
								for true {
									name1 := getPartName(ebr.Part_name)
									dotS += "<td rowspan=\"2\">EBR <br/>" + name1 + "</td>\n"
									porcentaje = (float64(ebr.Part_s) / float64(tamanioT)) * 100.0
									trunc := float64(int(porcentaje*100)) / 100
									dotS += "<td rowspan=\"2\">Logica <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"

									if ebr.Part_next == -1 {
										if (ebr.Part_start + ebr.Part_s) < mbr.Mbr_partition[i].Part_s {
											porcentaje = (float64(mbr.Mbr_partition[i].Part_s-(ebr.Part_start+ebr.Part_s)) / float64(tamanioT)) * 100
											trunc = float64(int(porcentaje*100)) / 100
											dotS += "<td rowspan=\"2\">Libre <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
										}
										break
									}
									if (ebr.Part_start + ebr.Part_s) < ebr.Part_next {
										porcentaje = (float64(ebr.Part_next-(ebr.Part_start+ebr.Part_s)) / float64(tamanioT)) * 100
										trunc = float64(int(porcentaje*100)) / 100
										dotS += "<td rowspan=\"2\">Libre <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
									}
									file.Seek(int64(ebr.Part_next), 0)
									errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(ebr))), binary.BigEndian, &ebr)
								}
							}
						}
						dotS += "<td rowspan=\"2\">EXTENDIDA</td>\n"
					}
					inicio = int(mbr.Mbr_partition[i].Part_start + mbr.Mbr_partition[i].Part_s)
				} else {
					i++
					for i < 4 {
						if mbr.Mbr_partition[i].Part_start != -1 {
							porcentaje := (float64(int(mbr.Mbr_partition[i].Part_start)-inicio) / float64(tamanioT)) * 100
							trunc := float64(int(porcentaje*100)) / 100
							dotS += "<td rowspan=\"2\">Libre <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
							break
						}
						i++
					}
					if i == 4 {
						porcentaje := float64(tamanioT-inicio) / float64(tamanioT) * 100
						trunc := float64(int(porcentaje*100)) / 100
						dotS += "<td rowspan=\"2\">Libre <br/>" + fmt.Sprintf("%v", trunc) + "</td>\n"
						goto t0
					}
					i--
				}
				i++
			}
		t0:
			dotS += "</tr></table>>];\n"
			dotS += "}"
			errD = os.WriteFile("Reportes/"+nombreD+".dot", []byte(dotS), 0777)
			if errD != nil {
				fmt.Println(errD)
			}

			file.Close()
			ext := Extrep
			_, errD = exec.Command("dot", "-T"+Extrep, "Reportes/"+nombreD+".dot", "-o", "Reportes/"+nombreD).Output()
			if errD != nil {
				fmt.Printf("%s", errD)
			}
			_, errD = exec.Command("dot", "-T"+ext, "Reportes/"+nombreD+".dot", "-o", Dirrep+nombreD).Output()
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

func sbR() Structs.Resp {
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
			sb := Structs.SuperBloque{}
			if nodo.Type == 'p' {
				mbr := Structs.MBR{}
				file.Seek(0, 0)
				errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(mbr))), binary.BigEndian, &mbr)
				if mbr.Mbr_partition[nodo.Pos].Part_status != '2' {
					file.Close()
					return Structs.Resp{Res: "NO SE HA FORMATEADO LA MONTURA DE LA PARTICION " + nodo.Name}
				}
				file.Seek(int64(nodo.Start), 0)
			} else if nodo.Type == 'l' {
				ebr := Structs.EBR{}
				file.Seek(int64(nodo.Start), 0)
				errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(ebr))), binary.BigEndian, &ebr)
				if ebr.Part_status != '2' {
					file.Close()
					return Structs.Resp{Res: "NO SE HA FORMATEADO LA MONTURA DE LA PARTICION " + nodo.Name}
				}
				file.Seek(int64(nodo.Start+int(unsafe.Sizeof(Structs.EBR{}))), 0)
			}
			errf = binary.Read(LeerFile(file, int(unsafe.Sizeof(sb))), binary.BigEndian, &sb)
			file.Close()

			dot, errD := os.OpenFile("Reportes/"+nombreD+".dot", os.O_CREATE, 0777)
			dot.Close()
			if errD != nil {
				fmt.Println(errD)
			}

			dotS := ""
			dotS += "digraph G {\n"
			dotS += "node[shape=none]\n"
			dotS += "start[label=<<table>\n"
			dotS += "<tr><td colspan=\"2\" bgcolor=\"#147e0d\"><font color=\"white\">REPORTE DE SUPERBLOQUE</font></td></tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">sb_nombre_hd</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + nombreD + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_filesystem_type</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + strconv.Itoa(int(sb.S_filesystem_type)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">s_inodes_count</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + strconv.Itoa(int(sb.S_inodes_count)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_blocks_count</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + strconv.Itoa(int(sb.S_blocks_count)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">s_free_blocks_count</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + strconv.Itoa(int(sb.S_free_blocks_count)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_free_inodes_count</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + strconv.Itoa(int(sb.S_free_inodes_count)) + "</td>\n"
			dotS += "</tr>\n"

			tm := time.Unix(sb.S_mtime, 0)
			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">s_mtime</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + tm.Format("2006-01-02 15:04:05") + "</td>\n"
			dotS += "</tr>\n"

			tm = time.Unix(sb.S_umtime, 0)
			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_umtime</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + tm.Format("2006-01-02 15:04:05") + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">s_mnt_count</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + strconv.Itoa(int(sb.S_mnt_count)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_magic</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + strconv.Itoa(int(sb.S_magic)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">s_inode_s</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + strconv.Itoa(int(sb.S_inode_s)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_block_s</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + strconv.Itoa(int(sb.S_block_s)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">s_firts_ino</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + strconv.Itoa(int(sb.S_firts_ino)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_first_blo</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + strconv.Itoa(int(sb.S_first_blo)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">s_bm_inode_start</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + strconv.Itoa(int(sb.S_bm_inode_start)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_bm_block_start</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + strconv.Itoa(int(sb.S_bm_block_start)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">s_inode_start</td>\n"
			dotS += "<td bgcolor=\"#b4f0b1\">" + strconv.Itoa(int(sb.S_inode_start)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "<tr>\n"
			dotS += "<td bgcolor=\"#27ba40\">s_block_start</td>\n"
			dotS += "<td bgcolor=\"#27ba40\">" + strconv.Itoa(int(sb.S_block_start)) + "</td>\n"
			dotS += "</tr>\n"

			dotS += "</table>>];\n"
			dotS += "}"

			errD = os.WriteFile("Reportes/"+nombreD+".dot", []byte(dotS), 0777)
			if errD != nil {
				fmt.Println(errD)
			}

			ext := Extrep
			_, errD = exec.Command("dot", "-T"+Extrep, "Reportes/"+nombreD+".dot", "-o", "Reportes/"+nombreD).Output()
			if errD != nil {
				fmt.Printf("%s", errD)
			}
			_, errD = exec.Command("dot", "-T"+ext, "Reportes/"+nombreD+".dot", "-o", Dirrep+nombreD).Output()
			if errD != nil {
				fmt.Printf("%s", errD)
			}

			return Structs.Resp{Res: "SE GENERO EL REPORTE SB"}
		}
		return Structs.Resp{Res: "DISCO INEXISTENTE"}
	}
	return Structs.Resp{Res: "NO SE HA ENCONTRADO ALGUNA MONTURA CON EL ID: " + Idrep}
}

func fileR() Structs.Resp {
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

func getPartName(partName [16]byte) string {
	name := ""
	for i := 0; i < 16; i++ {
		if partName[i] == '\000' {
			break
		}
		name += string(partName[i])
	}
	return name
}
