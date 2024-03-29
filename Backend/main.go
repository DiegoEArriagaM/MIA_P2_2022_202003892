package main

import (
	"Backend/Sistema_Archivos"
	"Backend/Structs"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Inicio")
	router := mux.NewRouter()
	enableCORS(router)

	router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		res := Structs.Inicio{
			Res: "Simulador de Disco Duro Web Corriendo",
			U:   Sistema_Archivos.UsuarioL,
		}
		json.NewEncoder(writer).Encode(res)
	}).Methods("GET")

	router.HandleFunc("/ListRep", func(writer http.ResponseWriter, req *http.Request) {
		r := make([]string, 2)
		reportes, er := os.ReadDir("Reportes/")
		if er != nil {
			fmt.Println(er)
		}
		for _, reporte := range reportes {
			i := find(reporte.Name(), ".")
			if i >= len(reporte.Name()) {
				r = append(r, reporte.Name())
			}

		}
		json.NewEncoder(writer).Encode(r)
	}).Methods("GET")

	router.HandleFunc("/Entrada", func(writer http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body) // response body is []byte
		var entrada Structs.Entrada
		if err := json.Unmarshal(body, &entrada); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Error al recibir el comando")
			fmt.Println(err)
		}
		reco := recover()
		if reco != nil {
			json.NewEncoder(writer).Encode(Structs.Inicio{Res: "Error en la entrada"})
		}

		Sistema_Archivos.UsuarioL = Structs.Usuario{
			IdU:     entrada.IdU,
			IdG:     entrada.IdG,
			IdMount: entrada.IdMount,
			NombreU: entrada.NombreU,
			Login:   entrada.Login,
		}
		r := Sistema_Archivos.Lector(entrada.Command)
		r.U = Sistema_Archivos.UsuarioL

		json.NewEncoder(writer).Encode(r)
	}).Methods("GET", "POST")

	router.HandleFunc("/Exec", func(writer http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body) // response body is []byte
		var entrada Structs.Exec
		if err := json.Unmarshal(body, &entrada); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Error al recibir el comando")
			fmt.Println(err)
		}
		reco := recover()
		if reco != nil {
			json.NewEncoder(writer).Encode(Structs.Inicio{Res: "Error en la entrada"})
		}

		Sistema_Archivos.UsuarioL = Structs.Usuario{
			IdU:     entrada.IdU,
			IdG:     entrada.IdG,
			IdMount: entrada.IdMount,
			NombreU: entrada.NombreU,
			Login:   entrada.Login,
		}
		respuesta := ""
		for _, s := range entrada.Commands {
			fmt.Println(s)
			if s == "pause" {
				fmt.Println("Presione Enter para continuar")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				fmt.Println()
				entrada.I++
				continue
			}
			eje := Sistema_Archivos.Lector(s)
			if eje.Res != "" {
				respuesta += strconv.Itoa(entrada.I) + ")" + eje.Res + "\n"
			}
			entrada.I++
		}
		r := Structs.Resp{Res: respuesta}
		r.U = Sistema_Archivos.UsuarioL

		json.NewEncoder(writer).Encode(r)
	}).Methods("GET", "POST")

	router.PathPrefix("/Reportes/").Handler(http.StripPrefix("/Reportes/", http.FileServer(http.Dir("./Reportes/"))))

	log.Fatal(http.ListenAndServe(":8000", router))
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,Access-Control-Allow-Origin")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

func find(cadena string, substring string) int {
	i := strings.Index(cadena, substring)
	if i == -1 {
		i = len(cadena)
	}
	return i
}
