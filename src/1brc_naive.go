package oneBRC

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
)

type stationCityMetric struct {
	minTemperature   float64
	sumOfTemperature float64
	maxTemperature   float64
	count            int
}

type metricEntry struct {
	stationCityName   string
	stationCityMetric *stationCityMetric
}

func roundTowardPositive(x float64) float64 {
	return math.Ceil(x*10) / 10
}

func populateStationCitiesMap(filePath string, stationCities map[string]*stationCityMetric) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ";")
		stationCity := words[0]
		stationCityTemperature, err := strconv.ParseFloat(words[1], 64)
		if err != nil {
			log.Fatalln("Error parsing string for station temperature into float64", err)
		}
		_, ok := stationCities[stationCity]
		if ok {
			stationCities[stationCity].maxTemperature = max(stationCities[stationCity].maxTemperature, stationCityTemperature)
			stationCities[stationCity].minTemperature = min(stationCities[stationCity].minTemperature, stationCityTemperature)
			stationCities[stationCity].sumOfTemperature += stationCityTemperature
			stationCities[stationCity].count++
		} else {
			stationCities[stationCity] = &stationCityMetric{
				maxTemperature:   stationCityTemperature,
				minTemperature:   stationCityTemperature,
				sumOfTemperature: stationCityTemperature,
				count:            1,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading file:", err)
	}
}

func generateOutput(stationCities map[string]*stationCityMetric) string {
	stationCityMetrics := []*metricEntry{}

	var outputString string
	outputString += "{"
	for key, value := range stationCities {
		stationCityMetrics = append(stationCityMetrics, &metricEntry{
			stationCityName:   key,
			stationCityMetric: value,
		})
	}
	sort.Slice(stationCityMetrics, func(i int, j int) bool {
		return stationCityMetrics[i].stationCityName < stationCityMetrics[j].stationCityName
	})
	for _, resMetric := range stationCityMetrics {
		resMetric := fmt.Sprintf("%v=%.1f/%.1f/%.1f, ", resMetric.stationCityName, resMetric.stationCityMetric.minTemperature, roundTowardPositive(resMetric.stationCityMetric.sumOfTemperature/float64(resMetric.stationCityMetric.count)), resMetric.stationCityMetric.maxTemperature)
		outputString += resMetric
	}
	outputString = outputString[:len(outputString)-2]
	outputString += "}\n"

	return outputString
}

func BrcNaive(filePath string) string {
	file, err := os.Create("naive_cpu.prof")
	if err != nil {
		log.Fatalln("Error in opening file: ", err)
	}
	defer file.Close()
	if err := pprof.StartCPUProfile(file); err != nil {
		log.Fatalln("Failed to start cpu profile: ", err)
	}
	defer pprof.StopCPUProfile()
	stationCities := map[string]*stationCityMetric{}
	populateStationCitiesMap(filePath, stationCities)

	answer := generateOutput(stationCities)
	return answer
}
