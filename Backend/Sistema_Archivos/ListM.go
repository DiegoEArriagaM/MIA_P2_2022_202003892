package Sistema_Archivos

import "Backend/Structs"

type Nodo_M struct {
	Path  string
	Name  string
	Id    string
	Num   int
	Pos   int
	Type  byte
	Letra byte
	Start int
	Sig   *Nodo_M
}

type MountList struct {
	Primero *Nodo_M
	Ultimo  *Nodo_M
}

func (L *MountList) add(path string, name string, ty byte, start int, pos int) Structs.Resp {
	if !L.existMount(path, name) {

	}
	return Structs.Resp{Res: "LA PARTICION " + name + " YA ESTA MONTADA"}
}

func (L *MountList) existMount(path string, name string) bool {
	aux := L.Primero
	for aux != nil {
		if aux.Path == path && aux.Name == name {
			return true
		}
		aux = aux.Sig
	}
	return false
}

func (L *MountList) getNum(path string) int {
	mayor := 0
	aux := L.Primero
	for aux != nil {
		if aux.Path == path && aux.Num > mayor {
			mayor = aux.Num
		}
		aux = aux.Sig
	}
	return mayor + 1
}

func (L *MountList) getLetra(path string) byte {
	aux := L.Primero
	var letraMayor byte = 64
	for aux != nil {
		letraAct := aux.Letra
		if aux.Path == path {
			return aux.Letra
		}
		if letraAct > letraMayor {
			letraMayor = letraAct
		}
		aux = aux.Sig
	}
	return letraMayor + 1
}
