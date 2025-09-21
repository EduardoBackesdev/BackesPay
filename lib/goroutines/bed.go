package lib

import (
	"main/internal/repositories"
	"time"
)

func Routine_bed(data repositories.BedRequest) {

	for {

		now := time.Now()

		start := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())
		end := time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, now.Location())
		weekday := now.Weekday()

		if weekday >= time.Monday && weekday <= time.Friday && now.After(start) && now.Before(end) {
			err := repositories.Bed(data)
			if err != nil {
				// Manda notificacao que houve algum erro ao enviar a ted
				break
			}
			// Aqui a logica que vai mandar a notificacao pro cliente
			break
		}

		time.Sleep(1 * time.Minute)

	}

}
