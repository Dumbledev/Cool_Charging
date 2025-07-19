package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/skratchdot/open-golang/open"
)

var tmpl *template.Template

var dbUrl = "http://admin:admin@localhost:5984/cool_app"
var dbFindUrl = dbUrl + "/_find"

func main() {
	Oerr := open.Run("C:/Users/Dumbledev/Desktop/Cool_ Charging")
	if Oerr != nil {
		fmt.Println(Oerr)
	}

	tmpl, _ = template.ParseGlob("templates/*.html")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", home)
	http.HandleFunc("/stats", stats)
	http.HandleFunc("/cards", cards)                                            //
	http.HandleFunc("/card", card)                                              //
	http.HandleFunc("/card_register", cardRegister)                             //
	http.HandleFunc("/card_register_handler", cardRegisterHandler)              //
	http.HandleFunc("/card_search_locker", cardSearchByLockerNo)                //
	http.HandleFunc("/card_search_locker_handler", cardSearchByLockerNoHandler) //
	http.HandleFunc("/card_search_serial", cardSearchBySerialNo)                //
	http.HandleFunc("/card_search_serial_handler", cardSearchBySerialNoHandler) //
	http.HandleFunc("/card_delete/{serialNo}", cardDelete)                      //

	http.HandleFunc("/charge_register", chargeRegister)
	http.HandleFunc("/charge_register_handler", chargeRegisterHandler)
	http.HandleFunc("/charges", charges)
	http.HandleFunc("/charge_collect", chargeCollect)
	http.HandleFunc("/charge_search", chargeSearch)
	http.HandleFunc("/charge_collect_handler", chargeCollectHandler)
	http.HandleFunc("/charge_collected", chargeCollected)
	http.HandleFunc("/charge_uncollected", chargeUncollected)
	http.HandleFunc("/charge_search_name", chargeSearchName)
	http.HandleFunc("/charge_search_name_handler", chargeSearchNameHandler)
	http.HandleFunc("/charge_search_locker", chargeSearchLocker)
	http.HandleFunc("/charge_search_locker_handler", chargeSearchLockerHandler)
	fmt.Println("Application Process: Cool Charging App Running")
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func stats(w http.ResponseWriter, r *http.Request) {
	var cards []Card
	year, month, day := time.Now().Date()
	type PageStruct struct {
		CardsCount       int
		DailyCount       int
		CollectedCount   int
		UnCollectedCount int
	}
	cardResp, err := findCards(dbFindUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	dailyChargeResp, err := findDailyCharges(dbFindUrl, month.String(), year, day)
	if err != nil {
		fmt.Println(err)
		return
	}
	collectedResp, err := findDailyCollectedCharges(dbFindUrl, month.String(), "Yes", year, day)
	if err != nil {
		fmt.Println(err)
		return
	}

	unCollectedResp, err := findDailyCollectedCharges(dbFindUrl, month.String(), "No", year, day)
	if err != nil {
		fmt.Println(err)
		return
	}

	cards = cardResp.Body
	dailyCount := dailyChargeResp.Body
	collectedCount := collectedResp.Body
	unCollectedCount := unCollectedResp.Body
	p := PageStruct{
		CardsCount:       len(cards),
		DailyCount:       len(dailyCount),
		CollectedCount:   len(collectedCount),
		UnCollectedCount: len(unCollectedCount),
	}
	tmpl.ExecuteTemplate(w, "stats.html", p)
}

func chargeCollected(w http.ResponseWriter, r *http.Request) {
	year, month, day := time.Now().Date()
	collectedResp, err := findDailyCollectedCharges(dbFindUrl, month.String(), "Yes", year, day)
	if err != nil {
		fmt.Println(err)
		return
	}
	collected := collectedResp.Body
	tmpl.ExecuteTemplate(w, "charge_collected.html", collected)
}

func chargeUncollected(w http.ResponseWriter, r *http.Request) {
	year, month, day := time.Now().Date()
	unCollectedResp, err := findDailyCollectedCharges(dbFindUrl, month.String(), "No", year, day)
	if err != nil {
		fmt.Println(err)
		return
	}
	unCollected := unCollectedResp.Body
	tmpl.ExecuteTemplate(w, "charge_uncollected.html", unCollected)
}
