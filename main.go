package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"training-service/session"

	"github.com/gorilla/mux"
)

func main() {
	mongoDbUrl := flag.String("mongodb", "localhost:27017", "Connection String to a MongoDB")
	apiPort := flag.String("apiPort", "7000", "REST API Port")

	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	sessionRepo := session.MongoDBSessionRepository{dbUrl: mongoDbUrl}

	fmt.Println("REST API listening on " + *apiPort)
	log.Fatal(http.ListenAndServe(":"+*apiPort, router))
}

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hello")
}

func oldmain() {
	memRepo := MemorizingRepo{}
	memRepo.trainings = make(map[string]Training)
	nonMemRepo := NonMemorizingRepo{}

	savedNonMem := save(nonMemRepo, Training{title: "T1", location: "Zuhause"})
	fmt.Println(savedNonMem)

	savedMem := save(memRepo, Training{title: "T1", location: "Zuhause"})
	fmt.Println(savedMem)
	savedTraining, found := load(memRepo, "T1")
	fmt.Println(found)
	fmt.Println(savedTraining)

}

func load(repo Repository, title string) (*Training, bool) {
	tr, found := repo.load(title)
	return tr, found
}

func save(repo Repository, training Training) bool {
	saved := repo.save(training)
	return saved
}

type Training struct {
	title, location string
}

type MemorizingRepo struct {
	trainings map[string]Training
}

type NonMemorizingRepo struct {
}

type Repository interface {
	load(title string) (*Training, bool)
	save(t Training) bool
}

func (repo NonMemorizingRepo) load(title string) (*Training, bool) {
	return nil, false
}

func (repo NonMemorizingRepo) save(t Training) bool {
	return false
}

func (repo MemorizingRepo) load(title string) (*Training, bool) {
	training, found := repo.trainings[title]
	return &training, found
}

func (repo MemorizingRepo) save(training Training) bool {
	repo.trainings[training.title] = training
	return true
}
