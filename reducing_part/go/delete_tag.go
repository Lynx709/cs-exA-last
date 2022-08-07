package main

import (
    "encoding/csv"
    "os"
	"github.com/gocarina/gocsv"
)

type inputTag struct {
	Tag string
	Lati string
	Long string
	Time string
	Url string
}

type newTag struct{
	Lati string
	Long string
	Time string
	Url string
}

func main() {
    file, err := os.Open("geotag.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    csvReader := csv.NewReader(file)
    inputs := []inputTag{}

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

        inputs = append(inputs, inputTag{
            Tag: tag,
			Lati: lati,
			Long: long,
			Time: time,
			Url: url})
    }
	
	newtag := []newTag{}
	outputfile, _ := os.OpenFile("delete_tag.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	for _, v := range inputs {
		newtag = append(newtag, newTag{
			Lati: v.Lati,
			Long: v.Long,
			Time: v.Time,
			Url: v.Url,
		})
	}
	defer outputfile.Close()
	gocsv.MarshalFile(newtag, outputfile)

}