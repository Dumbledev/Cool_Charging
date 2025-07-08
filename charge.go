package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

func chargeRegister(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "charge_register.html", nil)
}

func chargeRegisterHandler(w http.ResponseWriter, r *http.Request) {
	year, month, day := time.Now().Date()
	r.ParseForm()
	var price string = "0"
	name := strings.ToLower(r.FormValue("name"))
	device := r.FormValue("deviceType")
	adult := r.FormValue("adult")
	gender := r.FormValue("gender")
	serialNo := r.FormValue("serialNo")

	cardResp, err := findCardBySerial(dbFindUrl, serialNo)
	if err != nil {
		tmpl.ExecuteTemplate(w, "500.html", "Sever Error")
		return
	}
	if len(cardResp.Body) == 0 {
		tmpl.ExecuteTemplate(w, "charge_register.html", "Card Not Found In Database")
		return
	}
	card := cardResp.Body[0]
	chargeResp, err := findCharge(dbFindUrl, serialNo, "No")
	if err != nil {
		tmpl.ExecuteTemplate(w, "500.html", "Sever Error")
		return
	}
	if len(chargeResp.Body) != 0 {
		tmpl.ExecuteTemplate(w, "charge_register.html", "Card Already In Use")
		return
	}

	var charge = Charge{
		ID:         uuid.NewString(),
		Name:       name,
		DeviceType: device,
		Card:       card,
		Collected:  "No",
		Adult:      adult,
		Gender:     gender,
		Price:      price,
		Year:       year,
		Day:        day,
		Month:      month.String(),
		Date:       time.Now().Local().String(),
		Doctype:    "charge",
	}

	jsonData, err := json.Marshal(charge)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(jsonData))
	request, err := http.NewRequest("POST", dbUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Byte Error", err)
		return
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	res, error := client.Do(request)
	if error != nil {
		fmt.Println("Req Err", error)
	}
	defer res.Body.Close()
	tmpl.ExecuteTemplate(w, "charge_register.html", "Device Resgiter Successful")

}

func charges(w http.ResponseWriter, r *http.Request) {
	var charge []Charge
	chargeResp, err := findCharges(dbFindUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(chargeResp.Body) == 0 {
		tmpl.ExecuteTemplate(w, "charge_no_card.html", "No Devices Found")
		return
	}
	charge = chargeResp.Body
	tmpl.ExecuteTemplate(w, "charges.html", charge)
}

func chargeCollect(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "charge_collect.html", nil)
}

func chargeSearch(w http.ResponseWriter, r *http.Request) {
	var charge Charge
	serialNo := r.FormValue("serialNo")

	chargeResp, err := findCharge(dbFindUrl, serialNo, "No")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(chargeResp.Body) == 0 {
		tmpl.ExecuteTemplate(w, "charge_no_card.html", "No Phone Registered With This Card")
		return
	}
	charge = chargeResp.Body[0]
	tmpl.ExecuteTemplate(w, "charge.html", charge)
}

func chargeSearchLockerHandler(w http.ResponseWriter, r *http.Request) {
	var charge []Charge
	lockerNo := r.FormValue("lockerNo")
	collected := r.FormValue("collected")
	chargeResp, err := findChargeByLocker(dbFindUrl, lockerNo, collected)
	// fmt.Println(chargeResp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(chargeResp.Body) == 0 {
		tmpl.ExecuteTemplate(w, "charge_no_card.html", "No Phone Registered With This Card")
		return
	}
	charge = chargeResp.Body
	tmpl.ExecuteTemplate(w, "charge_collected.html", charge)
}

func chargeSearchNameHandler(w http.ResponseWriter, r *http.Request) {
	var charge []Charge
	name := strings.ToLower(r.FormValue("name"))
	collected := r.FormValue("collected")
	chargeResp, err := findChargeByName(dbFindUrl, name, collected)
	// fmt.Println(chargeResp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(chargeResp.Body) == 0 {
		tmpl.ExecuteTemplate(w, "charge_no_card.html", "No Phone Registered With This Card")
		return
	}
	charge = chargeResp.Body
	tmpl.ExecuteTemplate(w, "charge_collected.html", charge)
}

func chargeCollectHandler(w http.ResponseWriter, r *http.Request) {
	var charge Charge
	id := r.FormValue("chargeId")

	chargeResp, err := findChargeById(dbFindUrl, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(chargeResp.Body) == 0 {
		tmpl.ExecuteTemplate(w, "charge_no_card.html", "No Phone Registered With This Card")
		return
	}
	charge = chargeResp.Body[0]
	charge.Collected = "Yes"

	jsonData, err := json.Marshal(&charge)
	if err != nil {
		fmt.Println(err)
		return
	}
	request, err := http.NewRequest("PUT", dbUrl+"/"+charge.ID, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Set("content-type", "application/json")
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(string(body))
	http.Redirect(w, r, "/charge_collect", http.StatusSeeOther)
}

func chargeSearchName(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "charge_search_name.html", nil)
}

func chargeSearchLocker(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "charge_search_locker.html", nil)
}
