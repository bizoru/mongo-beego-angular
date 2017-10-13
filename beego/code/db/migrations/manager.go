package migrations

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"api/models"
	"path/filepath"
	"api/db"
	"reflect"
)

func MigrateJSON(file string, obj interface{}) string {

	f,_:= filepath.Abs("./db/migrations/"+file+".json")
	raw, _ := ioutil.ReadFile(f)

	if reflect.TypeOf(obj) == reflect.TypeOf(models.Jornada{}) {
		fmt.Println("Es una jornada")
		var jornadas []models.Jornada
		err := json.Unmarshal(raw, &jornadas)
		if err != nil {
			panic(err)
		}
		for k := range jornadas {
			session, _ := db.GetSession()
			jornada := jornadas[k]
			models.InsertJornada(session,jornada)
			fmt.Println("Insert Jornada with name: " + jornada.Nombre)
		}
		return "jornadas migrated"
	}

	if reflect.TypeOf(obj) == reflect.TypeOf(models.Candidato{}) {
		fmt.Println("Es un un candidato")
		var candidatos []models.Candidato
		err := json.Unmarshal(raw, &candidatos)
		if err != nil {
			panic(err)
		}
		for k := range candidatos {
			session, _ := db.GetSession()
			candidato := candidatos[k]
			models.InsertCandidato(session,candidato)
			fmt.Println("Insert Candidato with name: " + candidato.Nombre)
		}
		return "candidatos migrated"
	}

	return "none"

}

func GetMenu() []byte {
	f, _ := filepath.Abs("./db/static/menu.json")
	raw, _ := ioutil.ReadFile(f)
	return raw
}