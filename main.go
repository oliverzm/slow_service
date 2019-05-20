package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func getCustomers(w http.ResponseWriter, r *http.Request) {
	log.Print("get customers")
	listCust := []int{}
	for j := 1; j <= 1000; j++ {
		listCust = append(listCust, j)
	}

	json.NewEncoder(w).Encode(map[string][]int{"data": listCust})
}

func getRules(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	cid := keys.Get("cid")
	if cid == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "param cid is missing"})
		return
	}
	log.Print("get rules for the customer: ", cid)

	listCustRules := []string{}
	for j := 1; j <= 10; j++ {
		rule := strings.Join([]string{"cid", cid, "rule", strconv.Itoa(j)}, "_")
		listCustRules = append(listCustRules, rule)
	}

	json.NewEncoder(w).Encode(map[string][]string{"data": listCustRules})
}

func getData(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	cid := keys.Get("cid")
	rule := keys.Get("rule")
	if cid == "" || rule == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "error with the url params"})
		return
	}

	log.Print("search data for cid: ", cid, " and rule: ", rule)

	time.Sleep(10 * time.Second)
	json.NewEncoder(w).Encode(map[string]string{"data": "ok"})
}

func main() {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	r.Route("/slow", func(r chi.Router) {
		r.Get("/getCustomers", getCustomers)
		r.Get("/getRules", getRules)
		r.Get("/getData", getData)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
