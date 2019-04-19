package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	pop := Population{size: 5000, target: []string{"H", "e", "l", "l", "o", " ", " W", "o", "r", "l", "d"},
		mutation_rate: 0.01, target_score: 1, finished: false}

	for pop.finished == false {

		pop = generate_population(pop)

		pop = generate_mating_pool(pop)

		pop = create_children(pop)

		pop = find_most_fit(pop)

		fmt.Println(pop.curr_best)

	}

}
