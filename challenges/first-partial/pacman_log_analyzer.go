package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//OSData ...
type OSData struct {
	InstalledPackages map[string]bool
	UpgradedPackages  map[string]bool
	RemovedPackages   map[string]bool
}

//PackageDetails ...
type PackageDetails struct {
	InstalledAt  string
	UpdatedAt    string
	TimesUpdated int
	DeletedAt    string
	PackageName  string
}

func main() {
	//fmt.Println("Pacman Log Analyzer")

	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//START: CREATE GLOBAL STORAGE VARIABLES
	Report := OSData{InstalledPackages: make(map[string]bool), UpgradedPackages: make(map[string]bool), RemovedPackages: make(map[string]bool)}
	Packages := make(map[string]PackageDetails)
	//END

	//START: SCAN FILE
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			fmt.Println("bruh1")
		}
		//START: PROCESS LINE
		lineFromScanner := scanner.Text()
		parsedDataFromLine := strings.Split(lineFromScanner, " ")
		if len(parsedDataFromLine) < 5 {
			continue
		}
		//END

		//START: CREATE LOCAL VARIABLES WITH DATA
		queryAction := parsedDataFromLine[3]
		queryDate := parsedDataFromLine[0] + " " + parsedDataFromLine[1]
		queryDate = strings.TrimSuffix(queryDate, "]")
		queryDate = strings.Replace(queryDate, "[", "", 1)
		queryPackage := parsedDataFromLine[4]
		//END

		//START: PROCESS COLLECTED DATA
		switch queryAction {
		case "installed":
			if !Report.InstalledPackages[queryPackage] {
				Report.InstalledPackages[queryPackage] = true
			}

			if Packages[queryPackage].DeletedAt != "-" {
				Packages[queryPackage] = PackageDetails{
					PackageName: queryPackage,
					InstalledAt: queryDate,
					UpdatedAt:   queryDate,
					DeletedAt:   "-",
				}

			} else {
				Packages[queryPackage] = PackageDetails{
					PackageName:  queryPackage,
					InstalledAt:  queryDate,
					UpdatedAt:    queryDate,
					TimesUpdated: 0,
					DeletedAt:    "-",
				}
			}
			break
		case "upgraded":
			if !Report.UpgradedPackages[queryPackage] {
				Report.UpgradedPackages[queryPackage] = true
			}

			if Packages[queryPackage].InstalledAt != "" {
				packageDetails := Packages[queryPackage]
				packageDetails.UpdatedAt = queryDate
				packageDetails.TimesUpdated++
				Packages[queryPackage] = packageDetails
			}
			break
		case "removed":
			if !Report.RemovedPackages[queryPackage] {
				Report.RemovedPackages[queryPackage] = true
			}

			if Packages[queryPackage].InstalledAt != "" {
				packageDetails := Packages[queryPackage]
				packageDetails.DeletedAt = queryDate
				Packages[queryPackage] = packageDetails
			}
			break
		}
		//END
	}
	//END: SCAN FILE

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//START: LOG TO FILE
	f, err := os.OpenFile("packages_report.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log.SetFlags(0)
	log.SetOutput(f)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	printReport(Report)
	printList(Packages)
	//END
}

func printReport(Report OSData) {
	log.Println("Pacman Packages Report")
	log.Println("----------------------")
	log.Println(" - Installed packages: " + strconv.Itoa(len(Report.InstalledPackages)))
	log.Println(" - Upgraded packages: " + strconv.Itoa(len(Report.UpgradedPackages)))
	log.Println(" - Removed packages: " + strconv.Itoa(len(Report.RemovedPackages)))
	log.Println(" - Current installed: " + strconv.Itoa(len(Report.InstalledPackages)-len(Report.RemovedPackages)))
	log.Println()
}

func printList(data map[string]PackageDetails) {
	log.Println("List of packages")
	log.Println("------------------")

	for _, packageDetails := range data {
		log.Println("- Package Name : " + packageDetails.PackageName)
		log.Println("\t- Install queryDate : " + packageDetails.InstalledAt)
		log.Println("\t- Last update queryDate : " + packageDetails.InstalledAt)
		log.Println("\t- How many updates : " + strconv.Itoa(packageDetails.TimesUpdated))
		log.Println("\t- Removal queryDate : " + packageDetails.DeletedAt)
	}
}
