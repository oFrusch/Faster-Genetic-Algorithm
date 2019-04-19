package main

import (
	"math/rand"
)

type DNA struct {
	genes   []string
	length  int
	fitness float64
}

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
	"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", " "}

func generate_genes(dna DNA) DNA {
	for i := 0; i < dna.length; i++ {
		random := rand.Intn(52)
		dna.genes = append(dna.genes, alphabet[random])
	}
	return dna
}

func calc_fitness(dna DNA, target []string) DNA {
	curr_fit := 0
	for i := 0; i < dna.length; i++ {
		if dna.genes[i] == target[i] {
			curr_fit += 1
		}
	}
	dna.fitness = float64(curr_fit) / float64(dna.length)
	// fmt.Println(dna.fitness)
	return dna
}

func crossover(par1 DNA, par2 DNA) DNA {
	child := DNA{length: par1.length}
	midpoint := rand.Intn(par1.length)
	for i := 0; i < midpoint; i++ {
		child.genes = append(child.genes, par1.genes[i])
	}
	for i := midpoint; i < child.length; i++ {
		child.genes = append(child.genes, par2.genes[i])
	}
	return child
}

func mutate(dna DNA, rate float64) DNA {
	for i := 0; i < len(dna.genes); i++ {
		random := rand.Float64()
		if random < rate {
			rand_index := rand.Intn(52)
			dna.genes[i] = alphabet[rand_index]
		}
	}
	return dna
}
