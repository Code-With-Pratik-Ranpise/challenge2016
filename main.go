package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Distributor struct {
	Name            string
	Permissions     []Permissions
	SubDistributors []Distributor
}

type Permissions struct {
	AllowedCountries  []string
	ExcludedProvinces []string
	ExcludedCities    []string
}

func main() {
	csvFile := "cities.csv"
	countryMap, provinceMap, cityMap, err := readCSVFile(csvFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	distributor1 := createDistributor1()

	distributor2 := createDistributor2()

	distributor3 := createDistributor3()

	fmt.Println("Enter Distributor Name (DISTRIBUTOR1, DISTRIBUTOR2, or DISTRIBUTOR3):")
	distributorName := strings.ToUpper(getUserInput())

	var selectedDistributor Distributor

	switch distributorName {
	case "DISTRIBUTOR1":
		selectedDistributor = distributor1
	case "DISTRIBUTOR2":
		selectedDistributor = distributor2
	case "DISTRIBUTOR3":
		selectedDistributor = distributor3
	default:
		fmt.Println("Invalid distributor name.", distributorName)
		return
	}

	fmt.Println("Country Name:- ")
	countryName := strings.ToUpper(getUserInput())
	fmt.Println("Province Name:- ")
	provinceName := strings.ToUpper(getUserInput())
	fmt.Println("City Name:- ")
	cityName := strings.ToUpper(getUserInput())

	permission := PermissionToCheck(selectedDistributor, countryName, provinceName, cityName, countryMap, provinceMap, cityMap)

	if permission {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func createDistributor1() Distributor {
	return Distributor{
		Name: "DISTRIBUTOR1",
		Permissions: []Permissions{
			Permissions{
				AllowedCountries:  []string{"INDIA", "UNITED STATES"},
				ExcludedProvinces: []string{"KARNATAKA"},
				ExcludedCities:    []string{"CHENNAI"},
			},
		},
	}
}

func createDistributor2() Distributor {
	return Distributor{
		Name: "DISTRIBUTOR2",
		Permissions: []Permissions{
			Permissions{
				AllowedCountries:  []string{"INDIA"},
				ExcludedProvinces: []string{"TAMIL NADU"},
			},
		},
	}
}

func createDistributor3() Distributor {
	return Distributor{
		Name: "DISTRIBUTOR3",
		Permissions: []Permissions{
			Permissions{
				AllowedCountries:  []string{"INDIA"},
				ExcludedProvinces: []string{},
				ExcludedCities:    []string{},
			},
		},
	}
}

func readCSVFile(filename string) (map[string]bool, map[string]bool, map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, nil, nil, err
	}

	countryMap := make(map[string]bool)
	provinceMap := make(map[string]bool)
	cityMap := make(map[string]bool)

	for _, line := range lines {
		if len(line) < 3 {
			continue
		}

		country := strings.ToUpper(line[5])
		province := strings.ToUpper(line[4])
		city := strings.ToUpper(line[3])

		countryMap[country] = true
		provinceMap[province] = true
		cityMap[city] = true
	}

	return countryMap, provinceMap, cityMap, nil
}

func PermissionToCheck(distributor Distributor, country, province, city string, countryMap, provinceMap, cityMap map[string]bool) bool {
	if !countryMap[country] {
		return false
	}

	for _, permission := range distributor.Permissions {
		if !stringSliceContains(permission.AllowedCountries, country) {
			continue
		}

		if stringSliceContains(permission.ExcludedProvinces, province) {
			return false
		}

		if stringSliceContains(permission.ExcludedCities, city) {
			return false
		}
	}

	return true
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func stringSliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
