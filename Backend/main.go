package main

import (
	"Backend/Sistema_Archivos"
	"Backend/Structs"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Inicio")
	router := mux.NewRouter()
	enableCORS(router)
	router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		json.NewEncoder(writer).Encode(Structs.Inicio{Res: "Simulador de Disco Duro Web Corriendo"})
	}).Methods("GET")

	router.HandleFunc("/Entrada", func(writer http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body) // response body is []byte
		var command Structs.Comando
		if err := json.Unmarshal(body, &command); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Error al recibir el comando")
		}
		reco := recover()
		if reco != nil {
			json.NewEncoder(writer).Encode(Structs.Inicio{Res: "Error en la entrada"})
		}
		r := Sistema_Archivos.Lector(command.Command)
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
