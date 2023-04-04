package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

type TEvent struct {
	Time     time.Time
	Event    TEVENT
	Charname string
}

type TEVENT string

const (
	dirStat = "stats"

	EVENT_LOGIN       TEVENT = "LOGIN"
	EVENT_LOGOUT      TEVENT = "LOGOUT"
	EVENT_CONTINUE    TEVENT = "CONTINUE"
	EVENT_INITIALIZED TEVENT = "INITIALIZED"
)

// generateStatisticsFolder creates the folder where the statistics will be stored
func generateStatisticsFolder() error {
	if !PathExists(dirStat) {
		if err := CreateFolder(dirStat); err != nil {

			return err
		}
	}

	return nil
}

// getEvents returns a list of events from a file
func getEvents(r io.Reader, date string) []TEvent {
	var (
		result = make([]TEvent, 0)
	)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		linea := string(scanner.Bytes())
		if len(linea) > 0 {
			temp := strings.Split(linea, "\t")
			tempDate, _ := time.ParseInLocation("2006-01-02 15:04:05", date+" "+temp[0], time.Local)
			event := TEvent{
				Time:     tempDate,
				Charname: temp[1],
				Event:    TEVENT(temp[2]),
			}
			result = append(result, event)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Time.Unix() < result[j].Time.Unix()
	})

	return result
}

// generateDateItems returns a list of dates
func generateDateItems(days int) []string {
	items := make([]string, 0)
	for i := 0; i < days; i++ {
		items = append(items, fmt.Sprintf("%s", time.Now().AddDate(0, 0, -i).Format("2006-01-02")))
	}
	return items
}

// GetTotalCharsByDay returns the total of characters logged and the record of characters logged
func GetTotalCharsByDay(date string) (logged int, record int) {
	logged = 0
	record = 0

	err := generateStatisticsFolder()
	if err != nil {
		return

	}

	if !PathExists(dirStat + "/" + date + ".txt") {
		return
	}

	f, err := OpenFile(dirStat + "/" + date + ".txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer f.Close()

	var (
		chars       []string
		currentList = make([]string, 0)
		maxRecord   = 0
	)

	events := getEvents(f, date)

	for _, event := range events {
		if event.Event == EVENT_LOGIN || event.Event == EVENT_CONTINUE {
			currentList = ArrayStringAddOnce(currentList, event.Charname)
			chars = ArrayStringAddOnce(chars, event.Charname)
		} else if event.Event == EVENT_LOGOUT {
			currentList = ArrayStringRemove(currentList, event.Charname)
		} else if event.Event == EVENT_INITIALIZED {
			currentList = make([]string, 0)
		}
		if len(currentList) > maxRecord {
			maxRecord = len(currentList)
		}
	}

	return len(chars), maxRecord
}
