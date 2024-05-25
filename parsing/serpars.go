package parsing

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetTemperature(url string) interface{} {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "ru,en;q=0.9,en-GB;q=0.8,en-US;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "units=metric; _gid=GA1.2.1416876533.1716552432; cityid=524901; gdpr_cookie_consent=eyJpdiI6Im9OeFppV1ROMHJcL0thcjV4em9EcGNnPT0iLCJ2YWx1ZSI6IlZLc3NGTkZFMGZuQ1BMSmpIOEZXKzR5bkk5VkVRcE9vajlCZzBTK2YzenFRa1ZJQjlpR1dMbkUyblluZHFoc2FUeEdsRWxKSmc0TTVcL2RWK0JHUHZvcTlTamVMd1N2QlNOdEFBdjR2UDNveXdwT2pFek56dThab3EwZHNnazBUMSIsIm1hYyI6ImE2OWU4NWZjY2EwMzA3YTI0OWIyZmNmNjhjOTNlOTQyYjU0M2M5YTc3N2NkYjg0MTcyNzllNzE1MWE5ODdlZTUifQ%3D%3D; october_session=eyJpdiI6InJpXC9kOHcyNGhKbFF5SmkxMG5yOXlBPT0iLCJ2YWx1ZSI6InVDbHhmem1vcFkrU0pSbHF4SkRqRU9HbnVHSEhJRjM5dnNUdTkrYXU0S0JJcnBiRnNGaGZlcWJcL05tSldUOXptOHNFU3BQK1M3UDdXNlpSTlBIeGVUWWhXdGFQS3BXK284cGEzZE8zN3lqTnJueTB5UXQrQTVHbHFiYVI5TzBzcCIsIm1hYyI6IjJkNzdjNzE2NWQ2NDU0OGYyODYwY2MxMDU2YTQ0ZDc1OGRjMjczMzY2YjA3ZjY2YWEwODRjMWNkMmQ4OWIwZjQifQ%3D%3D; _ga_31TSX35RJT=GS1.1.1716625076.3.1.1716625112.24.0.0; _ga=GA1.2.777363465.1716552431")
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
	foo := jsonMap["current"]

	FooMap := foo.(map[string]interface{})
	return FooMap["temp"]
}
