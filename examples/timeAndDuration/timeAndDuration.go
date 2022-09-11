package timeAndDuration

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	readerPromptWithDefault "jerno.playground.com/readerPromptWithDefault"
	utils "jerno.playground.com/utils"
)

var reader = readerPromptWithDefault.ReaderWrapper{Reader: bufio.NewReader(os.Stdin)}

func Run() {
	defer utils.StopWatchLogger("Example")()

	formatLayout := "02 Jan 06 15:04 MST" // See https://www.geeksforgeeks.org/time-formatting-in-golang/

	fmt.Println("Provide zones to convert between")
	fmt.Printf("💡 Find full list of zones here: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones\n")

	zone1Name, zone2Name, zonePromptError := promtZones()
	if zonePromptError != nil {
		fmt.Printf("Error: %v, Cannot read input\n", zonePromptError)
		return
	}

	fmt.Println("Convert " + zone1Name + " time to " + zone2Name + " time")

	timeStringZone1, readError := reader.PromptWithDefault("Time to convert in HH:MM format", "16:00")
	if readError != nil {
		fmt.Printf("Error: %v, Cannot read input\n", readError)
		return
	}
	now := time.Now().Format(time.RFC3339)
	timeStringZone1 = now[:11] + timeStringZone1 + now[16:]

	zone1, zone2, zoneLoadError := loadZones(zone1Name, zone2Name)
	if zoneLoadError != nil {
		fmt.Printf("Error: %v, Cannot load zones\n", zoneLoadError)
		return
	}

	tZone1, parseError := time.ParseInLocation(time.RFC3339, timeStringZone1, zone1)
	if parseError != nil {
		fmt.Printf("Error: %v, Cannot parse input\n", parseError)
		return
	}

	fmt.Printf("🕑 Time in %v: %v\n", zone1Name, tZone1.Format(formatLayout))
	fmt.Printf("🕑 Time in %v: %v\n", zone2Name, tZone1.In(zone2).Format(formatLayout))
}

// Load zones corresponding to the user input. See https://en.wikipedia.org/wiki/List_of_tz_database_time_zones for more information
func loadZones(zoneName1, zoneName2 string) (*time.Location, *time.Location, error) {
	zone1, errZ1 := time.LoadLocation(zoneName1)
	if errZ1 != nil {
		return nil, nil, errZ1
	}

	zone2, errZ2 := time.LoadLocation(zoneName2)
	if errZ2 != nil {
		return nil, nil, errZ2
	}

	return zone1, zone2, nil
}

func promtZones() (string, string, error) {
	zone1Name, read1Error := reader.PromptWithDefault("Zone to convert from", "Europe/Paris")
	if read1Error != nil {
		return "", "", read1Error
	}

	zone2Name, read2Error := reader.PromptWithDefault("Zone to convert from", "America/Chicago")
	if read2Error != nil {
		return "", "", read2Error
	}

	return zone1Name, zone2Name, nil
}

func getInputTimeComponents(components []string) (int, int, error) {
	hour, errH := strconv.Atoi(components[0])
	if errH != nil {
		return 0, 0, errH
	}

	minute, errM := strconv.Atoi(components[1])
	if errM != nil {
		return 0, 0, errM
	}

	return hour, minute, nil
}
