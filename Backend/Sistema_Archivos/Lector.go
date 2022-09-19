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
					i = strings.Index(entradaL, "->") + 2
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = strings.Index(entradaL, " ")
					if i == -1 {
						i = len(entradaL)
					}
					s, _ := strconv.Atoi(entradaL[:i])
					Sdisk = s
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-f") {
					i = strings.Index(entradaL, "->") + 2
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = strings.Index(entradaL, " ")
					if i == -1 {
						i = len(entradaL)
					}
					f := entradaL[:i]
					Fdisk = f
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-u") {
					i = strings.Index(entradaL, "->") + 2
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = strings.Index(entradaL, " ")
					if i == -1 {
						i = len(entradaL)
					}
					u := entradaL[:i]
					Udisk = u
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-path") {
					i = strings.Index(entradaL, "->") + 2
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					if entradaL[0] == '"' {
						entradaL = entradaL[1:]
						entradaO = entradaO[1:]
						i = strings.Index(entradaL, "\"")
						if i == -1 {
							i = len(entradaL)
						}
						p := entradaO[:i]
						Pdisk = p
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = strings.Index(entradaL, " ")
						if i == -1 {
							i = len(entradaL)
						}
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
