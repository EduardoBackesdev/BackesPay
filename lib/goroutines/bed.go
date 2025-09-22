package lib

import (
	"main/internal/repositories"
	"main/internal/ws"
	"time"
)

type Response struct {
	Message string
}

func Routine_bed(data repositories.BedRequest) {

	for {

		now := time.Now()

		start := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())
		end := time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, now.Location())
		weekday := now.Weekday()

		if weekday >= time.Monday && weekday <= time.Friday && now.After(start) && now.Before(end) {
			err := repositories.Bed(data)
			if err != nil {
				ws.Mutex.Lock()
				if conn, ok := ws.Map_clients[data.Id]; ok {
					conn.WriteJSON(Response{
						Message: "Bed sent fail!",
					})
				}
				ws.Mutex.Unlock()
				break
			}
			ws.Mutex.Lock()
			if conn, ok := ws.Map_clients[data.Id]; ok {
				conn.WriteJSON(Response{
					Message: "Bed sent successfully!",
				})
			}
			ws.Mutex.Unlock()
			break
		}

		time.Sleep(1 * time.Minute)

	}

}
