package main

import (
    "encoding/csv"
    "os"
	"github.com/gocarina/gocsv"
	"fmt"
	//"unsafe"
)

type InputTag struct {
	Time string
	Lati string
	Long string
	Url string
	Tag string
}
type newTag struct {
	Tag string
	Lati string
	Long string
	Time string
	Url string
}

func main() {
    file, err := os.Open("true.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    csvReader := csv.NewReader(file)
    inputs := []InputTag{}

    for {
        line, err := csvReader.Read()
        if err != nil {
            break
        }

        time := line[0]
		lati := line[1]
		long := line[2]
		url := line[3]
		tag := line[4]
		
        if err != nil {
            continue
        }

        inputs = append(inputs, InputTag{
            Time: time,
			Lati: lati,
			Long: long,
			Url: url,
			Tag: tag})
    }
	
	newtag := []newTag{}
	outputfile, _ := os.OpenFile("geotag.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	tmp := inputs[0].Tag
	nowTag := inputs[0].Tag
	cnt := 0

	for index, v := range inputs {
		if index % 100000 == 0{
			fmt.Println(index)
		}
		nowTag = v.Tag
		if nowTag != tmp {
			tmp = nowTag
			cnt = 0
			newtag = append(newtag, newTag{
				Tag: v.Tag,
				Lati: v.Lati,
				Long: v.Long,
				Time: v.Time,
				Url: v.Url,
			})
		} else {
			if cnt < 100{
				newtag = append(newtag, newTag{
					Tag: v.Tag,
					Lati: v.Lati,
					Long: v.Long,
					Time: v.Time,
					Url: v.Url,
				})
				cnt += 1
			}
		}
	}
	defer outputfile.Close()
	gocsv.MarshalFile(newtag, outputfile)

}