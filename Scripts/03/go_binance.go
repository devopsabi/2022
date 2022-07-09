package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	Log *log.Logger
)
var wg sync.WaitGroup

func init() {
	// set location of log file
	var logpath = "cryto.log"
	log.Println(logpath)

	flag.Parse()
	var file, err1 = os.OpenFile(logpath, os.O_APPEND|os.O_WRONLY, 0644)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}

func CallUrl(url string, cryptoname string) {

	defer wg.Done()

	fmt.Println("Calling url ", url)

	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	onedayregE := regexp.MustCompile(`price":"\$ \d*\D\d*`)
	OneDayPrice := onedayregE.FindAllString(string(html), -1)
	fmt.Println(OneDayPrice)

	if len(OneDayPrice) > 0 {
		currPrice := strings.Replace(OneDayPrice[0], "price\":\"$", "", -1)
		currPrice = strings.Replace(currPrice, ",", "", -1)
		currPrice = strings.ReplaceAll(currPrice, " ", "")
		file, fileErr := os.Create(cryptoname + ".txt")
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}
		fmt.Fprintf(file, "bitcoin,currency=%v value=%v\n", cryptoname, currPrice)
		Log.Printf("bitcoin,currency=%v value=%v\n", cryptoname, currPrice)

		fmt.Println("Current Price in Dollar", currPrice)

		priceChange1hPositiveregE := regexp.MustCompile(`class="css-1q7gaws">\+\d.\d*?%`)
		priceChange1hPositive := priceChange1hPositiveregE.FindAllString(string(html), -1)


		if len(priceChange1hPositive) > 0 {
			fmt.Println("POSITIVE Value + + + + + + + + + + + + + + + + + +")
			oneHourChangePositive := strings.Replace(priceChange1hPositive[0], "class=\"css-1q7gaws\">", "", -1)
			fmt.Println(oneHourChangePositive)
			oneHourChangePositive = strings.Replace(oneHourChangePositive, "+", "", -1)
			oneHourChangePositive = strings.Replace(oneHourChangePositive, "%", "", -1)
			file, fileErr := os.Create(cryptoname + "_1hPositive.txt")
			if fileErr != nil {
				fmt.Println(fileErr)
				return
			}
			fmt.Fprintf(file, "bitcoin_1hPositive,currency=%v value=%v\n", cryptoname, oneHourChangePositive)
			Log.Printf("bitcoin_1hPositive,currency=%v value=%v\n", cryptoname, oneHourChangePositive)

		} else {
			fmt.Println("NEGATIVE Value - - - - - - - - - - - - - - - - - -")
			priceChange1hNegativeregE := regexp.MustCompile(`class="css-okmmzw">\-\d.\d*?%`)
			priceChange1hNegative := priceChange1hNegativeregE.FindAllString(string(html), -1)

			oneHourChangeNegative := strings.Replace(priceChange1hNegative[0], "class=\"css-okmmzw\">", "", -1)
			fmt.Println("1 Hour fluctuation percentage", oneHourChangeNegative)
			oneHourChangeNegative = strings.Replace(oneHourChangeNegative, "+", "", -1)
			oneHourChangeNegative = strings.Replace(oneHourChangeNegative, "%", "", -1)
			file, fileErr := os.Create(cryptoname + "_1hNegative.txt")
			if fileErr != nil {
				fmt.Println(fileErr)
				return
			}
			fmt.Fprintf(file, "bitcoin_1hNegative,currency=%v value=%v\n", cryptoname, oneHourChangeNegative)
			Log.Printf("bitcoin_1hNegative,currency=%v value=%v\n", cryptoname, oneHourChangeNegative)

			oneDayChangeNegative := strings.Replace(priceChange1hNegative[1], "class=\"css-okmmzw\">-", "", -1)
			fmt.Println("1 Day fluctuation percentage", oneDayChangeNegative)
			Log.Printf("1 Day fluctuation percentage", oneDayChangeNegative)


		}

	}
}

func main() {



	wg.Add(3)

	go CallUrl("https://www.binance.com/en/price/bitcoin", "bitcoin")
	go CallUrl("https://www.binance.com/en/price/ethereum", "ethereum")
	go CallUrl("https://www.binance.com/en/price/shiba-inu", "shiba-inu")

	wg.Wait()
}
