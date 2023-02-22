func main() {
	c := Calendar{}
	c.Name = "8A"

	ev := Event{}
	var err error

	ev.Start, err = time.Parse("02.01.2006 15:04:05", "01.09.2022 8:30:00")
	if err != nil {
		log.Fatal(err)
	}

	ev.End, err = time.Parse("02.01.2006 15:04:05", "01.07.2023 9:15:00")
	if err != nil {
		log.Fatal(err)
	}

	ev.Summary = "Алгебра"

	ev.Location = "М41"

	ev.Alarm = 10

	c.Events = append(c.Events, ev)
	c.Events = append(c.Events, ev)
	c.Events = append(c.Events, ev)

	Make_calendar(c, "./cals/8A/cal.ics")
}
