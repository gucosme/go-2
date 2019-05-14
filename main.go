package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type player struct {
	FullName    string
	Nationality string
	Club        string
	Wage        float32
	Age         int
}

func main() {
	//Todas as perguntas são referentes ao arquivo data.csv
	f, err := os.Open("./data.csv")
	check(err)
	defer f.Close()

	r := csv.NewReader(f)
	r.Read()

	players := []player{}

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		wage, _ := strconv.ParseFloat(line[17], 32)
		age, _ := strconv.ParseInt(line[6], 10, 0)

		p := player{
			FullName:    line[1],
			Nationality: line[14],
			Club:        line[3],
			Wage:        float32(wage),
			Age:         int(age),
		}

		players = append(players, p)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go q1(players, wg)
	go q2(players, wg)

	wg.Wait()
}

//Quantas nacionalidades (coluna `nationality`) diferentes existem no arquivo?
func q1(players []player, wg *sync.WaitGroup) (int, error) {
	defer wg.Done()
	nationalities := map[string]int{}

	for _, p := range players {
		nationalities[p.Nationality] = 0
	}

	println("Nationalities", len(nationalities))

	return len(nationalities), nil
}

//Quantos clubes (coluna `club`) diferentes existem no arquivo?
func q2(players []player, wg *sync.WaitGroup) (int, error) {
	defer wg.Done()
	clubs := map[string]int{}

	for _, p := range players {
		clubs[p.Club] = 0
	}

	println("Clubs", len(clubs))

	return len(clubs), nil
}

//Liste o primeiro nome dos 20 primeiros jogadores de acordo com a coluna `full_name`.
func q3(players []player) ([]string, error) {

	return []string{}, fmt.Errorf("Not implemented")
}

//Quem são os top 10 jogadores que ganham mais dinheiro (utilize as colunas `full_name` e `eur_wage`)?
func q4() ([]string, error) {
	return []string{}, fmt.Errorf("Not implemented")
}

//Quem são os 10 jogadores mais velhos (use como critério de desempate o campo `eur_wage`)?
func q5() ([]string, error) {
	return []string{}, fmt.Errorf("Not implemented")
}

//Conte quantos jogadores existem por idade. Para isso, construa um mapa onde as chaves são as idades e os valores a contagem.
func q6() (map[int]int, error) {
	idades := make(map[int]int)
	return idades, fmt.Errorf("Not implemented")
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
		panic(e)
	}
}
