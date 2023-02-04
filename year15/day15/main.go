// Package day15 - Solution for Advent of Code 2015/15
// Problem Link: http://adventofcode.com/2015/day/15
package day15

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (maxScore int, maxScore500Cals int) {
	nutrients := parse(input)

	for sprinkles := 0; sprinkles <= 100; sprinkles++ {
		for peanutButter := 0; peanutButter <= 100; peanutButter++ {
			for frosting := 0; frosting <= 100; frosting++ {
				if sugar := 100 - sprinkles - peanutButter - frosting; sugar > 0 {
					teaspoons := map[string]int{
						"Sprinkles":    sprinkles,
						"PeanutButter": peanutButter,
						"Frosting":     frosting,
						"Sugar":        sugar,
					}
					score, calories := nutrients.score(teaspoons)
					// Part 1 - just get the score
					if score > maxScore {
						maxScore = score
					}
					// Part 2 - sacrifice some score for 500 calories
					if calories == 500 && score > maxScore500Cals {
						maxScore500Cals = score
					}
				}
			}
		}
	}

	return maxScore, maxScore500Cals
}

func parse(input string) (data Ingredients) {
	ingredients := Ingredients{}
	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, ":")
		var capacity, durability, flavor, texture, calories int
		fmt.Sscanf(
			tokens[1],
			" capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&capacity, &durability, &flavor, &texture, &calories)

		ingredients[tokens[0]] = Ingredient{
			capacity,
			durability,
			flavor,
			texture,
			calories,
		}
	}

	return ingredients
}

type Ingredients map[string]Ingredient

func (ingredients Ingredients) score(teaspoons map[string]int) (int, int) {
	nutrientValues := map[string]int{}
	for ingredient, teaspoon := range teaspoons {
		composition := ingredients[ingredient].nutrientsForQuantity(teaspoon)
		for nutrient, score := range composition {
			nutrientValues[nutrient] += score
		}
	}

	totalScore := 1

	for nutrient, value := range nutrientValues {
		if nutrient != "calories" {
			if value < 0 {
				value = 0
			}
			totalScore *= value
		}
	}

	return totalScore, nutrientValues["calories"]
}

type Ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func (ingredient Ingredient) nutrientsForQuantity(teaspoon int) map[string]int {
	return map[string]int{
		"capacity":   ingredient.capacity * teaspoon,
		"durability": ingredient.durability * teaspoon,
		"flavor":     ingredient.flavor * teaspoon,
		"texture":    ingredient.texture * teaspoon,
		"calories":   ingredient.calories * teaspoon,
	}
}
