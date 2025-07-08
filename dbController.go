package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func findCards(url string) (CardResponse, error) {
	var response CardResponse
	jsonData := map[string]map[string]any{"selector": {"doctype": "card"}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findCardBySerial(url string, serialNo string) (CardResponse, error) {
	var response CardResponse
	// jsonData := map[string]map[string]string{"selector": {"userId": selector}}
	jsonData := map[string]map[string]any{"selector": {"serialNo": serialNo, "doctype": "card"}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findCardByLocker(url string, lockerNo string) (CardResponse, error) {
	var response CardResponse
	// jsonData := map[string]map[string]string{"selector": {"userId": selector}}
	jsonData := map[string]map[string]any{"selector": {"lockerNo": lockerNo, "doctype": "card"}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findCharges(url string) (ChargeResponse, error) {
	var response ChargeResponse

	type Selector struct {
		Doctype string `json:"doctype"`
	}

	type Query struct {
		Selector Selector `json:"selector"`
		Limit    int      `json:"limit"`
	}
	query := Query{
		Selector: Selector{
			Doctype: "charge",
		},
		Limit: 200,
	}

	// jsonData := map[string]map[string]any{"selector": {"doctype": "charge"}}
	data, error := json.Marshal(query)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findChargeById(url, id string) (ChargeResponse, error) {
	var response ChargeResponse
	jsonData := map[string]map[string]any{"selector": {"doctype": "charge", "_id": id}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findChargeByName(url, name, collected string) (ChargeResponse, error) {
	var response ChargeResponse
	jsonData := map[string]map[string]any{"selector": {"doctype": "charge", "name": name, "collected": collected}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findCharge(url, cardSerial, collected string) (ChargeResponse, error) {
	var response ChargeResponse
	jsonData := map[string]map[string]any{"selector": {"doctype": "charge", "card.serialNo": cardSerial, "collected": collected}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findChargeByLocker(url, lockerNo, collected string) (ChargeResponse, error) {
	var response ChargeResponse
	jsonData := map[string]map[string]any{"selector": {"doctype": "charge", "card.lockerNo": lockerNo, "collected": collected}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findDailyCharges(url, month string, year, day int) (ChargeResponse, error) {
	var response ChargeResponse
	jsonData := map[string]map[string]any{"selector": {"doctype": "charge", "day": day, "month": month, "year": year}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}

func findDailyCollectedCharges(url, month, collected string, year, day int) (ChargeResponse, error) {
	var response ChargeResponse
	jsonData := map[string]map[string]any{"selector": {"doctype": "charge", "day": day, "month": month, "year": year, "collected": collected}}
	data, error := json.Marshal(jsonData)
	if error != nil {
		log.Fatalln("Marshal", error)
	}
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if error != nil {
		fmt.Println("Byte Error", error)
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	error = json.Unmarshal(body, &response)
	if error != nil {
		log.Fatalln("UnMarshal Err: ", error)
	}
	return response, error
}
