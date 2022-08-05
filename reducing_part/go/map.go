package main

import (
    "encoding/csv"
    "os"
	"github.com/gocarina/gocsv"
)

type newTag struct {
	Tag string
	Lati string
	Long string
	Time string
	Url string
}

type mapTag struct{
	Index int
	Tag string
}

func main() {
    file, err := os.Open("geotag.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    csvReader := csv.NewReader(file)
    inputs := []newTag{}

    for {
        line, err := csvReader.Read()
        if err != nil {
            break
        }

        tag := line[0]
		lati := line[1]
		long := line[2]
		time := line[3]
		url := line[4]
		
        if err != nil {
            continue
        }

        inputs = append(inputs, newTag{
            Tag: tag,
			Lati: lati,
			Long: long,
			Time: time,
			Url: url})
    }
	
	maptag := []mapTag{}
	outputfile, _ := os.OpenFile("map.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	tmp := inputs[0].Tag
	nowTag := inputs[0].Tag
	cnt := 0
	firstIndex := 0

	for index, v := range inputs {
		nowTag = v.Tag
		if nowTag != tmp {
			if cnt >= 99{
				maptag = append(maptag, mapTag{
					Index: firstIndex,
					Tag: tmp,
				})
			}
			firstIndex = index
			cnt = 0
			tmp = nowTag
		} else {
			cnt += 1
		}
	}
	defer outputfile.Close()
	gocsv.MarshalFile(maptag, outputfile)

}