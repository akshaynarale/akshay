package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type OrgCsv struct {
	id                    string
	orgNodeName           string
	name                  string
	orgExtCode            string
	description           string
	orgTypeName           string
	orgCodeID             string
	organization_category string
	prodName              string
	countryID             string
	stateID               string
	cityID                string
	parentID              string
	costCenterID          string
	status                string
	IsDeactive            string
	isRootNode            string
	orgNodeLeaderID       string
	parentBUID            string
}

func main() {

	records, err := readData("./csv.csv")

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {

		user := OrgCsv{
			id:                    record[0],
			orgNodeName:           record[1],
			name:                  record[2],
			orgExtCode:            record[3],
			description:           record[4],
			orgTypeName:           record[5],
			orgCodeID:             record[6],
			organization_category: record[7],
			prodName:              record[8],
			countryID:             record[9],
			stateID:               record[10],
			cityID:                record[11],
			parentID:              record[12],
			costCenterID:          record[13],
			status:                record[14],
			IsDeactive:            record[15],
			isRootNode:            record[16],
			orgNodeLeaderID:       record[17],
			parentBUID:            record[18],
		}

		fmt.Printf("%s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s %s  is a %s\n", user.id, user.orgNodeName,
			user.name, user.orgExtCode, user.description, user.orgTypeName, user.orgCodeID,
			user.organization_category, user.prodName, user.countryID, user.stateID, user.cityID,
			user.parentID, user.costCenterID, user.status, user.IsDeactive, user.isRootNode, user.orgNodeLeaderID, user.parentBUID)
	}
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
