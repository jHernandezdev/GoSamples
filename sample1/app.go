package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func main() {

	http.HandleFunc("/", func (rw http.ResponseWriter, r *http.Request) {	
		names := r.URL.Query()["name"]

		var name string		
		if len(names) == 1 {
			name = names[0]			
		}
		m := map[string] string{"name": name}
		enc := json.NewEncoder(rw)
		enc.Encode(m)				
	})

	http.HandleFunc("/hola", func(rw http.ResponseWriter, r *http.Request) {		
		rw.Write([]byte("Hola"))
	})

	err := http.ListenAndServe(":3000", nil)
	
	if (err != nil){
		log.Fatal(err)
	}
}
