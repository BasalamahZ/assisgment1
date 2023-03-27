package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Id        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var persons []Person = []Person{
	{
		Id:     1,
		Nama:   "Zidane",
		Alamat: "Jalan Soekarno Hatta",
		Pekerjaan: "programmer",
		Alasan: "karena saya ingin memperdalam ilmu tentang bahasa golang",
	},
	{
		Id:     2,
		Nama:   "Ali",
		Alamat: "Jalan Bunga Coklat",
		Pekerjaan: "programmer",
		Alasan: "Karena saya ingin mempelajari bahasa baru",
	},
	{
		Id:     3,
		Nama:   "Basalamah",
		Alamat: "Jalan Raya Sulfat",
		Pekerjaan: "programmer",
		Alasan: "Karena saya suka ngoding dan suka mengeksplore seputar ngoding",
	},
}

func main() {
	var input = os.Args
	result, err := FindBiodata(input[1])

	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("%+v \n", result)

}

func FindBiodata(personName string) (Person, error) {
	for _, value := range persons {
		if value.Nama == personName {
			return value, nil
		}
	}

	return Person{}, errors.New("Biodata not found")
}
