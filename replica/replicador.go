package replica

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var bdReplica []Trucos

type Trucos struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

func Sincronizador() {
	for {
		resp, err := http.Get("http://localhost:4000/truco")
		if err != nil {
			fmt.Println("Error al conectarse al servidor principal:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error al leer la solicitud: ", err)
			resp.Body.Close()
			time.Sleep(2 * time.Second)
			continue
		}
		resp.Body.Close()

		var newTruco []Trucos
		if err := json.Unmarshal(body, &newTruco); err != nil {
			fmt.Println("Error al parsear JSON: ", err)
			time.Sleep(2 * time.Second)
			continue
		}

		bdReplica = newTruco
		fmt.Println("Servidor de replicación actualizado", bdReplica)

		time.Sleep(5 * time.Second)
	}
}
