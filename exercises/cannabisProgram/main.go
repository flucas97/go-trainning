package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	fmt.Printf("Starting application...\n\n\n")
	time.Sleep(time.Second)

	createExamples()
}

type (
	flowers struct {
		name   string
		thc    int
		origin map[string]int
	}
	cannabis struct {
		name    string
		flowers []flowers
	}
	flowerCollection interface {
		showAllFlowers()
	}
)

var (
	cSativa           = cannabis{name: "Sativa"}
	cIndica           = cannabis{name: "Indica"}
	flowersCollection []flowerCollection
	args              = os.Args[1:]
	divisionLine      = func() { fmt.Println(strings.Repeat("-", 50)) }
)

func main() {
	switch len(args) {
	case 0:
		break
	case 5:
		createFlowerFromUser()
	default:
		fmt.Println("[type] [flowers] [thc] [country] [qtde]")
		return
	}

	flowersCollection = append(flowersCollection, &cSativa, &cIndica)

	fmt.Printf("\t\t C A N N A B I S:\n")
	divisionLine()
	for _, cannabis := range flowersCollection {
		cannabis.showAllFlowers()
	}
}

func (c *cannabis) updateFlower(n string, thc int, origin map[string]int) error {
	var flag bool

	if thc <= 0 {
		return errors.New("THC must be bigger or equal than 0")
	}

	newFlower := flowers{
		name:   n,
		thc:    thc,
		origin: make(map[string]int),
	}
	newFlower.origin = origin

	if flag = c.flowerExists(n); !flag {
		c.flowers = append(c.flowers, newFlower)
	} else {
		return errors.New("Flower alreay exists")
	}

	return nil
}

func (c *cannabis) flowerExists(flowerName string) bool {
	for _, f := range c.flowers {
		if f.name != flowerName {
			continue
		} else {
			return true
		}
	}
	return false
}

func (c *cannabis) showAllFlowers() {
	fmt.Printf("\nThe %s flowers are:\n", c.name)
	for _, f := range c.flowers {
		fmt.Printf("Name: %v \t ------- \tTHC: %v\n", f.name, f.thc)
		for cannabis, qtde := range f.origin {
			fmt.Printf("Country: %v \t ------- \tQtde: %v\n\n", cannabis, qtde)
		}
	}
	divisionLine()
}

func createOrigin(c string, q int) (map[string]int, error) {
	if q < 0 {
		return nil, errors.New("Quantity must be bigger or equal than 0")
	}

	return map[string]int{
		c: q,
	}, nil
}

func createExamples() {
	eX, _ := createOrigin("Brazil", 15)
	eY, _ := createOrigin("France", 200)

	_ = (&cSativa).updateFlower("Gorilla Haze", 27, eX)

	_ = cIndica.updateFlower("Notherland", 22, eY)

}
func createFlowerFromUser() {
	cannabisType, flowers, country, q, typeExists := args[0], args[1], args[3], args[4], false

	thc, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalln(err)
		return
	}

	avaliableTypes := []string{
		"Sativa",
		"Indica",
	}
	_ = typeExists

	for _, existentType := range avaliableTypes {
		if cannabisType == existentType {
			typeExists = true
			break
		}

		avaliableTypes = append(avaliableTypes, cannabisType)
	}

	qtde, err := strconv.Atoi(q)
	if err != nil {
		log.Fatalln(err)
		return
	}

	origin, err := createOrigin(country, qtde)
	if err != nil {
		log.Fatalln(err)
		return
	}

	switch cannabisType {
	case "Sativa":
		if err = cSativa.updateFlower(flowers, thc, origin); err != nil {
			log.Fatalln(err)
			return
		}

	case "Indica":
		err = cIndica.updateFlower(flowers, thc, origin)
		if err != nil {
			log.Fatalln(err)
			return
		}
	default:
		// TODO
	}
}
