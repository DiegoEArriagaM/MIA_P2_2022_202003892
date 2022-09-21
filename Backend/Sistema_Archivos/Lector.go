package Sistema_Archivos

import (
	"Backend/Structs"
	"strconv"
	"strings"
)

func Lector(comando string) Structs.Resp {
	res := ""
	entradaO := comando
	entradaL := strings.ToLower(comando)

	if len(entradaO) > 0 {
		if strncmp(entradaL, "#") {
			return Structs.Resp{Res: res, Reporte: false}
		} else if strncmp(entradaL, "mkdisk") {
			i := 6
			for entradaL[i] == ' ' && len(entradaL) > 0 {
				i++
			}
			entradaL = entradaL[i:]
			entradaO = entradaO[i:]

			for len(entradaO) > 0 {
				if strncmp(entradaL, "-s") {
					i = find(entradaL, "->") + 2
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = find(entradaL, " ")
					s, _ := strconv.Atoi(entradaL[:i])
					Sdisk = s
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-f") {
					i = find(entradaL, "->") + 2
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = find(entradaL, " ")
					f := entradaL[:i]
					Fdisk = f
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-u") {
					i = find(entradaL, "->") + 2
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = find(entradaL, " ")
					u := entradaL[:i]
					Udisk = u
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-path") {
					i = find(entradaL, "->") + 2
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					if entradaL[0] == '"' {
						entradaL = entradaL[1:]
						entradaO = entradaO[1:]

						i = find(entradaL, "\"")
						p := entradaO[:i]
						Pdisk = p
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						p := entradaO[:i]
						Pdisk = p
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}

				} else if strncmp(entradaL, "#") {
					break
				} else {
					return Structs.Resp{Res: "ERROR EN EL COMANDO DE ENTRADA: " + entradaO, Reporte: false}
				}
			}

			return mkdisk()
		} else if strncmp(entradaL, "rmdisk") {
			i := 6
			for entradaL[i] == ' ' && len(entradaL) > 0 {
				i++
			}
			entradaL = entradaL[i:]
			entradaO = entradaO[i:]
			for len(entradaO) > 0 {

			}
			return rmdisk()
		} else {
			return Structs.Resp{Res: "Comando no reconocido", Reporte: false}
		}
	}

	return Structs.Resp{Res: res, Reporte: false}
}

func strncmp(entrada string, comparacion string) bool {
	if len(comparacion) > len(entrada) {
		return false
	}

	for i := 0; i < len(comparacion); i++ {
		if entrada[i] != comparacion[i] {
			return false
		}
	}

	return true
}

func find(cadena string, substring string) int {
	i := strings.Index(cadena, substring)
	if i == -1 {
		i = len(cadena)
	}
	return i
}
