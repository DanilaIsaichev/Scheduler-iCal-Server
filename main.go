package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	db_data "github.com/danilaisaichev/schedule_db_data"
	ical "github.com/danilaisaichev/schedule_ical"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/gen_ical", gen_ical)

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "3636"
	}

	fmt.Println("iCal server is listening on port: " + port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func gen_ical(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Not allowed"))

	} else {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// Считываем тело запроса в буффер
		buff, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Объявляем переменную для расписания на неделю
		week := db_data.Week{}

		// Парсинг
		err = json.Unmarshal(buff, &week)
		if err != nil {
			log.Fatal(err)
		}

		for _, day := range week.Data {

			for _, class := range day.Schedule {

				cal := ical.Calendar{}

				cal.Name = class.Class

				for _, lesson := range class.Lessons {

					event := ical.Event{}

					event.Set_datetime(day.Date, lesson.Number)

					event.Summary = lesson.Name
					event.Location = lesson.Room

					event.Alarm = 10

					cal.Events = append(cal.Events, event)
				}

				ical.Make_calendar(cal, "./cals/"+class.Class+"/cal.ics")
			}
		}

	}

}
