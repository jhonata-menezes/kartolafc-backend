package kartolafc

import (
	"time"
	"log"
	"bytes"
	"encoding/gob"
)

func TemJogoComPartidaRolando() bool {
	horaAtual := time.Now()
	for _, p := range CachePartidas[0].Partidas {
		horaPartida, err := time.Parse("2006-01-02 15:04:05", p.PartidaData)
		if err != nil {
			log.Println("nao identificou a hora do jogo")
		} else {
			horaPartida = horaPartida.Add(3 * time.Hour)
			if horaAtual.Unix() >= horaPartida.Unix() && horaAtual.Unix() <= horaPartida.Add(2 * time.Hour).Unix() {
				return true
			}
		}

	}
	return false
}

func CopyStructs(from, to interface{}) {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(from)
	dec.Decode(to)
}
