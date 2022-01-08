package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
)

type Response struct {
		Rate Rate `json:"rates"`
	}

type Rate struct {
		XAU float64 `json:"XAU"`

}

func main() {

	var response Response

	url := "https://www.metals-api.com/api/latest?access_key=XXXXXXXXXXXXX&base=INR&symbols=XAU"
	resp, err := http.Get(url)


	if err != nil {
		panic(err)
	}

	// In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns
	defer resp.Body.Close()

	// Let's Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("1 Ounce Rate = %v\n", response.Rate.XAU)

	// 1 ounces = 28.3495 grams
	// 1 gm = ?

	// 1 ounces = 133965.24182391775
	// 1 gram = 4725.488697293347

	oneGram := (response.Rate.XAU / 28.3495 )


	fmt.Println("One Gram Price is = ", oneGram)

	file, fileErr := os.Create("file")
	if fileErr != nil {
	    fmt.Println(fileErr)
	    return
	}
	fmt.Fprintf(file, "gold_value,currency=INR value=%v\n", oneGram)

}

// curl -i -XPOST 'http://localhost:8086/write?db=our_expenses' --data-binary 'cpu_load_short,host=server02 value=0.67'
