package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://api-adservices.apple.com/api/v1/"

	payload := strings.NewReader("HxIcmHHiDnrcnpcyQ6i9FmZNnw4IdeN5K9C7EMygUms0uq9sWrPjTgK8juZUkKJD7e75u3vJf/uADwNM2PjZ2F6OCPREZ57Q4wAAAVADAAAA5QAAAIBILcUe6yKld4vdMKDRX0qXjzcNqMEdOQX4EshII95yoHHv2bNB6fQu7GVcHmn6+mNWlITj8PMt3yqxfMU0yDKo5BNcYMfzJoeKDJ84+XXuOaE95OLy5l5eF0zthTiNWZrLC75qHz1HIQX+1LU0+0PeVCQ0MLItoluIECZ5T5w4/gAAABj0l9LP4SPQTIeEMLITwwn9L4jtrOrxMAkAAACfASCsc/+vPkfK2eU9vXH/i4p9BpQ2AAAAhgAEd2E5hVtqT/Gdql0urexcUWmqN205VxVr5ZsBUtngJpo3+SYITQvJzGHfCTfv+eeGfXF9ZMI1aIseLeoVyNX9izAKjsEFE6Zp8T4e/FSMIYud+C93l6iwZP+9O6/DA5gBKO3I6M9Vto87/LzxFli+vL0mvKl6Y5q670ljU9uomeLwmL+wAAAAAAAAAAABBEocAAA=")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("User-Agent", "PostmanRuntime-ApipostRuntime/1.1.0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
