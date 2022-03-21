package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

type Things struct {
	InventoryId int        `json:"inventory_id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Tags        []string   `json:"tags"`
	PurchasedAt int        `json:"purchased_at"`
	Placement   Placements `json:"placement"`
}

type Placements struct {
	RoomId int    `json:"room_id"`
	Name   string `json:"name"`
}

func main() {
	fmt.Println("Hello World !")
	fmt.Println("-------------")
	list := readJson("./data.json")
	ListItems(list)
	fmt.Println("-------------")
	ListItemsMeetingRoom(list)
	fmt.Println("-------------")
	ListAllElectronicDevices(list)
	fmt.Println("-------------")
	ListOfFurniture(list)
	fmt.Println("-------------")
	ListItemsPurchasedAt(list, "16 January 2020")
	fmt.Println("-------------")
	ListItemsBrownColor(list)
	fmt.Println("-------------")
}

func ListItemsBrownColor(list []Things) {
	no := 0
	fmt.Println("List Items with Brown Color :")
	for _, l := range list {
		for _, j := range l.Tags {
			if j == "brown" {
				no = no + 1
				fmt.Printf("%d. %s\n", no, l.Name)
				break
			}
		}
	}
}

func ListItemsPurchasedAt(list []Things, date string) {
	no := 0
	fmt.Println("List Items Purchased At", date, ":")
	for _, l := range list {
		datePurchased := strconv.Itoa(l.PurchasedAt)
		k, err := strconv.ParseInt(datePurchased, 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(k, 0)
		timeString := tm.Format("02 January 2006")
		if timeString == date {
			no = no + 1
			fmt.Printf("%d. %s %s\n", no, l.Name, timeString)
		}
	}
}

func ListOfFurniture(list []Things) {
	x := "furniture"
	no := 0
	fmt.Println("List of Electronic Devices :")
	for _, l := range list {
		if l.Type == x {
			no = no + 1
			fmt.Printf("%d. %s\n", no, l.Name)
		}
	}
}

func ListAllElectronicDevices(list []Things) {
	x := "electronic"
	no := 0
	fmt.Println("List of Electronic Devices :")
	for _, l := range list {
		if l.Type == x {
			no = no + 1
			fmt.Printf("%d. %s\n", no, l.Name)
		}
	}
}

func ListItemsMeetingRoom(list []Things) {
	x := "Meeting Room"
	no := 0
	fmt.Println("List of items in Meeting Room :")
	for _, l := range list {
		if l.Placement.Name == x {
			no = no + 1
			fmt.Printf("%d. %s\n", no, l.Name)
		}
	}
}

func ListItems(list []Things) {
	fmt.Println("List of items :")
	for i, l := range list {
		fmt.Printf("%d. %s\n", i+1, l.Name)
	}
}

func readJson(lokasi string) []Things {
	content, err := ioutil.ReadFile(lokasi)
	if err != nil {
		log.Fatal("Ada Error :", err)
	}
	var payloads []Things
	err = json.Unmarshal(content, &payloads)
	if err != nil {
		log.Fatal("Ada Error :", err)
	}
	return payloads
}
