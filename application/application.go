package application

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Calendar struct {
	Name   string
	Events []Event
}

// Структура события
type Event struct {
	Start    time.Time
	End      time.Time
	Summary  string
	Location string
	Alarm    int
}

// Функция, генерирующая текст события из структуры
func Generate_event(event_struct Event) string {

	return fmt.Sprint(`BEGIN:VEVENT
DTSTAMP:`, time.Now().Format("20060201T150405"), `
UID:`, strings.ToUpper(uuid.New().String()), `
DTSTART;TZID=Europe/Moscow:`, event_struct.Start.Format("20060201T150405"), `
DTEND;TZID=Europe/Moscow:`, event_struct.End.Format("20060201T150405"), `
SUMMARY:`, event_struct.Summary, `
LOCATION:`, event_struct.Location, `
BEGIN:VALARM
ACTION:DISPLAY
DESCRIPTION:`, event_struct.Summary, ` - `, event_struct.Location, `
TRIGGER:-PT`, math.Abs(float64(event_struct.Alarm)), `M
END:VALARM
END:VEVENT`)

}

// Функция, записывающая в файл с календарём новые события
func Add_events(events []Event, directory string) {

	// считываем данные из файла
	data, err := os.ReadFile(directory)
	if err != nil {
		log.Fatal(err)
	}

	// разбиваем данные на строки
	data_strs := strings.Split(string(data), "\n")

	// открываем файл на запись
	file, err := os.OpenFile(directory, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	// добавляем в массив строк строки нового события
	for i := 0; i < len(events); i++ {
		data_strs = append(data_strs[0:len(data_strs)-1], Generate_event(events[i]), data_strs[len(data_strs)-1])
	}

	// выводим в файл
	_, err = file.WriteString(strings.Join(data_strs, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
}

// Функция, генерирующая пустой календарь
func Make_calendar(name string, directory string) {

	var file *os.File

	if _, err := os.Stat(directory); err != nil {

		fmt.Println(directory)

		file, err = os.Create(directory)
		if err != nil {
			log.Fatal(err)
		}

		/* Парсим файл */
		cal, err := parse_ical(directory)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(cal)

		/* Сортируем массив? */

		/* Добавляем события */

	} else {

		calendar_str := `BEGIN:VCALENDAR
PRODID:SCHOOL80-SCHEDULER
NAME` + name + `
VERSION:2.0
CALSCALE:GREGORIAN
BEGIN:VTIMEZONE
TZID:Europe/Moscow
TZURL:http://tzurl.org/zoneinfo-outlook/Europe/Moscow
X-LIC-LOCATION:Europe/Moscow
BEGIN:STANDARD
TZNAME:MSK
TZOFFSETFROM:+0300
TZOFFSETTO:+0300
DTSTART:19700101T000000
END:STANDARD
END:VTIMEZONE
END:VCALENDAR`

		file, err = os.Create(directory)
		if err != nil {
			log.Fatal(err)
		}

		_, err := file.WriteString(calendar_str)
		if err != nil {
			log.Fatal(err)
		}

	}

	file.Close()
}

func parse_ical(directory string) (cal Calendar, err error) {

	// считываем данные из файла
	data, err := os.ReadFile(directory)
	if err != nil {
		log.Fatal(err)
	}

	// разбиваем данные на строки
	data_strs := strings.Split(string(data), "\n")

	for _, data_str := range data_strs {
		fmt.Println(data_str)
	}

	return Calendar{}, nil
}
