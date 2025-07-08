package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func cards(w http.ResponseWriter, r *http.Request) {
	var cards []Card

	cardResp, err := findCards(dbFindUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	cards = cardResp.Body
	tmpl.ExecuteTemplate(w, "cards.html", cards)
}

func card(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "card.html", nil)
}

func cardRegister(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "card_register.html", nil)
}

func cardRegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lockerNo := r.FormValue("lockerNo")
	serialNo := r.FormValue("serialNo")

	cardResp, err := findCardBySerial(dbFindUrl, serialNo)
	if err != nil {
		tmpl.ExecuteTemplate(w, "500.html", "Sever Error")
		return
	}
	if len(cardResp.Body) != 0 {
		tmpl.ExecuteTemplate(w, "card_register.html", "Card Already Registered")
		return
	}

	cardLockerResp, err := findCardByLocker(dbFindUrl, lockerNo)
	if err != nil {
		tmpl.ExecuteTemplate(w, "500.html", "Sever Error")
		return
	}
	if len(cardLockerResp.Body) != 0 {
		tmpl.ExecuteTemplate(w, "card_register.html", "Locker Already Registered To A Card. Delete Card Before Registering To Locker")
		return
	}

	var card = Card{
		ID:       uuid.NewString(),
		LockerNo: lockerNo,
		SerialNo: serialNo,
		Doctype:  "card",
	}

	jsonData, err := json.Marshal(card)
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
	tmpl.ExecuteTemplate(w, "card_register.html", "Card Register Success")
}

func cardSearchByLockerNo(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "card_search_locker.html", nil)
}

func cardSearchByLockerNoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lockerNo := r.FormValue("lockerNo")
	cardResp, err := findCardByLocker(dbFindUrl, lockerNo)
	if err != nil {
		fmt.Println(err)
		return
	}
	card := cardResp.Body[0]

	tmpl.ExecuteTemplate(w, "card.html", card)
}

func cardSearchBySerialNo(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "card_verify.html", nil)
}

func cardSearchBySerialNoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	serialNo := r.FormValue("serialNo")
	cardResp, err := findCardBySerial(dbFindUrl, serialNo)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(cardResp.Body) == 0 {
		tmpl.ExecuteTemplate(w, "card_not_found.html", "Card Not Found")
		return
	}
	card := cardResp.Body[0]
	tmpl.ExecuteTemplate(w, "card.html", card)
}

func cardDelete(w http.ResponseWriter, r *http.Request) {
	var card Card
	serialNo := r.PathValue("serialNo")
	cardResp, err := findCardBySerial(dbFindUrl, serialNo)
	if err != nil {
		fmt.Println(err)
	}
	if len(cardResp.Body) == 0 {
		fmt.Println("No Record Found")
		return
	}
	card = cardResp.Body[0]
	jsonData, err := json.Marshal(card)
	if err != nil {
		fmt.Println("Marshal Err", err)
		return
	}
	request, err := http.NewRequest("DELETE", dbUrl+"/"+card.ID, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Set("content-type", "application/json")
	request.Header.Set("If-Match", card.Rev)
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

	http.Redirect(w, r, "/cards", http.StatusSeeOther)
}
