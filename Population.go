package main

import (
	"math/rand"
	"strings"
)

type Population struct {
	size          int
	target        []string
	population    []DNA
	mating_pool   []DNA
	generations   int
	curr_best     string
	target_score  float64
	finished      bool
	mutation_rate float64
	curr_max_fit  float64
}

func generate_population(pop Population) Population {
	for i := 0; i < pop.size; i++ {
		member := DNA{length: len(pop.target)}
		member = generate_genes(member)
		member = calc_fitness(member, pop.target)
		//fmt.Println(member.fitness)
		pop.population = append(pop.population, member)
	}
	return pop
}

func generate_mating_pool(pop Population) Population {
	max_fit := 0.0
	pop.mating_pool = nil

	for i := 0; i < len(pop.population); i++ {
		if max_fit < pop.population[i].fitness {
			max_fit = pop.population[i].fitness
		}
	}

	for i := 0; i < len(pop.population); i++ {
		fit := pop.population[i].fitness / max_fit
		num_times := int(fit * 100)
		for j := 0; j < num_times; j++ {
			pop.mating_pool = append(pop.mating_pool, pop.population[i])
		}
	}

	return pop
}

func create_children(pop Population) Population {

	for i := 0; i < len(pop.population); i++ {

		rand1 := rand.Intn(len(pop.population))
		rand2 := rand.Intn(len(pop.population))

		parent1 := pop.mating_pool[rand1]
		parent2 := pop.mating_pool[rand2]

		child := DNA{length: parent1.length}
		child = crossover(parent1, parent2)
		child = mutate(child, pop.mutation_rate)
		child = calc_fitness(child, pop.target)

		pop.population[i] = child
	}

	return pop
}

func find_most_fit(pop Population) Population {
	max_fit := 0.0
	best_genes := 0

	for i := 0; i < len(pop.population); i++ {
		if pop.population[i].fitness > max_fit {
			max_fit = pop.population[i].fitness
			best_genes = i
		}
	}

	pop.curr_best = strings.Join(pop.population[best_genes].genes, "")
	pop.curr_max_fit = max_fit

	if max_fit == pop.target_score {
		pop.finished = true
	}

	return pop
}

func is_finished(pop Population) bool {
	return pop.finished
}
