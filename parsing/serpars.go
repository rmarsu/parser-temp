package parsing

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetTemperature(url string) int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "ru,en;q=0.9,en-GB;q=0.8,en-US;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "units=metric; _gid=GA1.2.1416876533.1716552432; cityid=524901; october_session=eyJpdiI6IlJpSWQxNWtkU3E5TEFURkJ3RVkxamc9PSIsInZhbHVlIjoiOGZleVJwanBYV2RTb1wveEdKd1JDUjYyd0JEQTNLQ25rOFl2cmFVZ1p2ZERYY2hQaDN0amp3Z01oM3QzY2JSbThRXC9IQm5hKzIxYzFrT0hhVDZkWnh3dUg4OGxUcUJlaTlEa29vMjhRMHcydmZBRlFzUG9VZUhGd2Ywb05xR1VtbCIsIm1hYyI6Ijg2ZTEzMDFjYzAzNTczZTg0MzYwN2YxM2JiZWJiNGQ0OTcwZjhiMzYxNjU1Zjc2OGRmOTU2MzFlYjI3YmYwNTQifQ%3D%3D; _ga_31TSX35RJT=GS1.1.1716552430.1.1.1716552529.60.0.0; _ga=GA1.1.777363465.1716552431")
	req.Header.Set("Referer", "https://openweathermap.org/city/524901")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Mobile Safari/537.36 Edg/126.0.0.0")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Microsoft Edge";v="126"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", `"Android"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jsonMap map[string]interface{}
	json.Unmarshal(bodyText, &jsonMap)
	// inner items are of type interface{}
	foo := jsonMap["current"]

	FooMap := foo.(map[string]interface{})
	ans , _ := fmt.Println(FooMap["temp"])
	return ans
}

