package principal

import "errors"

type Trucos struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var bd []Trucos
var idCounter = 1

func GetAllTrucos() []Trucos {
	return bd
}

func CreateTruco(name string) Trucos {
	truco := Trucos{Id: idCounter, Name: name}
	bd = append(bd, truco)
	idCounter++
	return truco
}

func UpdateTruco(id int, name string) error {
	for i, n := range bd {
		if n.Id == id {
			bd[i].Name = name
			return nil
		}
	}
	return errors.New("Truco no encontrado")
}

func DeleteTruco(id int) error {
	for i, n := range bd {
		if n.Id == id {
			bd = append(bd[:i], bd[i+1:]...)
			return nil
		}
	}
	return errors.New("Truco no encontrado")
}