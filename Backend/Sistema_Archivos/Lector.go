package Sistema_Archivos

import (
	"Backend/Structs"
	"strconv"
	"strings"
)

var UsuarioL Structs.Usuario
var Mlist MountList

func Lector(comando string) Structs.Resp {
	entradaO := comando
	entradaL := strings.ToLower(comando)

	if len(entradaO) > 0 {
		if strncmp(entradaL, "#") {
			return Structs.Resp{Res: ""}
		} else if strncmp(entradaL, "mkdisk") {
			i := 6
			for entradaL[i] == ' ' && len(entradaL) > 0 {
				i++
			}
			entradaL = entradaL[i:]
			entradaO = entradaO[i:]

			for len(entradaO) > 0 {
				if strncmp(entradaL, "-size") {
					i = find(entradaL, "=") + 1
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

				} else if strncmp(entradaL, "-fit") {
					i = find(entradaL, "=") + 1
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

				} else if strncmp(entradaL, "-unit") {
					i = find(entradaL, "=") + 1
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
					i = find(entradaL, "=") + 1
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
					return Structs.Resp{Res: "ERROR EN EL COMANDO DE ENTRADA: " + entradaO}
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
				if strncmp(entradaL, "-path") {
					i = find(entradaL, "=") + 1
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
					return Structs.Resp{Res: "ERROR EN EL COMANDO DE ENTRADA: " + entradaO}
				}
			}
			return rmdisk()
		} else if strncmp(entradaL, "fdisk") {
			i := 5
			for entradaL[i] == ' ' && len(entradaL) > 0 {
				i++
			}
			entradaL = entradaL[i:]
			entradaO = entradaO[i:]
			for len(entradaO) > 0 {
				if strncmp(entradaL, "-size") {
					i = find(entradaL, "=") + 1
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = find(entradaL, " ")
					s, _ := strconv.Atoi(entradaL[:i])
					Spart = s
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-unit") {
					i = find(entradaL, "=") + 1
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = find(entradaL, " ")
					u := entradaL[:i]
					Upart = u[0]
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-path") {
					i = find(entradaL, "=") + 1
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
						Ppart = p
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						p := entradaO[:i]
						Ppart = p
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}

				} else if strncmp(entradaL, "-type") {
					i = find(entradaL, "=") + 1
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = find(entradaL, " ")
					t := entradaL[:i]
					Tpart = t[0]
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-fit") {
					i = find(entradaL, "=") + 1
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					i = find(entradaL, " ")
					f := entradaL[:i]
					Fpart = f[0]
					for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

				} else if strncmp(entradaL, "-name") {
					i = find(entradaL, "=") + 1
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					if entradaL[0] == '"' {
						entradaL = entradaL[1:]
						entradaO = entradaO[1:]

						i = find(entradaL, "\"")
						n := entradaO[:i]
						Namepart = n
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						n := entradaO[:i]
						Namepart = n
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}

				} else if strncmp(entradaL, "#") {
					break
				} else {
					return Structs.Resp{Res: "ERROR EN EL COMANDO DE ENTRADA: " + entradaO}
				}
			}
			return fdisk()
		} else if strncmp(entradaL, "mount") {
			i := 5
			for entradaL[i] == ' ' && len(entradaL) > 0 {
				i++
			}
			entradaL = entradaL[i:]
			entradaO = entradaO[i:]
			for len(entradaO) > 0 {
				if strncmp(entradaL, "-path") {
					i = find(entradaL, "=") + 1
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
						Pmontar = p
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						p := entradaO[:i]
						Pmontar = p
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}

				} else if strncmp(entradaL, "-name") {
					i = find(entradaL, "=") + 1
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					if entradaL[0] == '"' {
						entradaL = entradaL[1:]
						entradaO = entradaO[1:]

						i = find(entradaL, "\"")
						n := entradaO[:i]
						Namemontar = n
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						n := entradaO[:i]
						Namemontar = n
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}

				} else if strncmp(entradaL, "#") {
					break
				} else {
					return Structs.Resp{Res: "ERROR EN EL COMANDO DE ENTRADA: " + entradaO}
				}
			}
			return mount()
		} else if strncmp(entradaL, "rep") {
			i := 3
			for entradaL[i] == ' ' && len(entradaL) > 0 {
				i++
			}
			entradaL = entradaL[i:]
			entradaO = entradaO[i:]
			for len(entradaO) > 0 {
				if strncmp(entradaL, "-path") {
					i = find(entradaL, "=") + 1
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
						Prep = p
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						p := entradaO[:i]
						Prep = p
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}

				} else if strncmp(entradaL, "-name") {
					i = find(entradaL, "=") + 1
					for entradaL[i] == ' ' && len(entradaL) > 0 {
						i++
					}
					entradaL = entradaL[i:]
					entradaO = entradaO[i:]

					if entradaL[0] == '"' {
						entradaL = entradaL[1:]
						entradaO = entradaO[1:]

						i = find(entradaL, "\"")
						n := entradaO[:i]
						Namerep = n
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						n := entradaO[:i]
						Namerep = n
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}
				} else if strncmp(entradaL, "-ruta") {
					i = find(entradaL, "=") + 1
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
						Rutarep = p
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						p := entradaO[:i]
						Rutarep = p
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}
				} else if strncmp(entradaL, "-id") {
					i = find(entradaL, "=") + 1
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
						Idrep = p
						i += 1
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]

					} else {
						i = find(entradaL, " ")
						p := entradaO[:i]
						Idrep = p
						for i < len(entradaL) && entradaL[i] == ' ' && len(entradaL) > 0 {
							i++
						}
						entradaL = entradaL[i:]
						entradaO = entradaO[i:]
					}
				} else if strncmp(entradaL, "#") {
					break
				} else {
					return Structs.Resp{Res: "ERROR EN EL COMANDO DE ENTRADA: " + entradaO}
				}
			}
			return GenerateRep()
		} else {
			return Structs.Resp{Res: "COMANDO NO RECONOCIDO"}
		}
	}

	return Structs.Resp{Res: ""}
}

func strncmp(entrada string, comparacion string) bool {
	if len(comparacion) > len(entrada) {
		return false
	}

	for i := 0; i < len(comparacion); i++ {
		if i >= len(entrada) {
			return false
		}
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
