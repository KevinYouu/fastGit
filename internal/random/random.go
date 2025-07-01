package random

import (
	"math/rand/v2"
)

// FuncProbability represents a function and its corresponding probability
type FuncProbability struct {
	Function    func()  // Function is the function to execute
	Probability float64 // Probability is the probability of executing the function
}

// ExecuteRandomly executes one of the given functions based on their probabilities
func ExecuteRandomly(funcProbs []FuncProbability) {
	// Calculate the total probability
	totalProb := 0.0
	for _, fp := range funcProbs {
		totalProb += fp.Probability
	}

	// Generate a random number in the range [0, totalProb)
	r := rand.Float64() * totalProb

	// Choose the function to execute based on the random number
	for _, fp := range funcProbs {
		if r < fp.Probability {
			fp.Function()
			return
		}
		r -= fp.Probability
	}
}
