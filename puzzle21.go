package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func intersect(d1, d2 []string) []string {
	tmp := make(map[string]struct{}, len(d2))
	for _, val := range d2 {
		tmp[val] = struct{}{}
	}
	var result []string
	for _, val := range d1 {
		if _, ok := tmp[val]; ok {
			result = append(result, val)
		}
	}
	return result
}

func remove(d1 []string, item string) ([]string, int) {
	result := make([]string, 0, len(d1)-1)
	var count int
	for _, val := range d1 {
		if val == item {
			count++
			continue
		}
		result = append(result, val)
	}
	return result, count
}

func puzzle21(data []string) int {
	allergens := make(map[string][]string)
	foods := make(map[string]int)
	expr := regexp.MustCompile(`^(.*) \(contains (.*)\)$`)
	for _, line := range data {
		matches := expr.FindStringSubmatch(line)
		sfoods := strings.Split(matches[1], " ")
		for _, food := range sfoods {
			foods[food]++
		}
		for _, allergen := range strings.Split(matches[2], ", ") {
			if cfoods, ok := allergens[allergen]; ok {
				allergens[allergen] = intersect(sfoods, cfoods)
				continue
			}
			allergens[allergen] = sfoods
		}
	}

	for {
		var changed bool
		for i, sfoods := range allergens {
			if len(sfoods) == 1 {
				for j, cfoods := range allergens {
					if i == j {
						continue
					}
					var count int
					allergens[j], count = remove(cfoods, sfoods[0])
					if count > 0 {
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
	}

	for _, sfoods := range allergens {
		if len(sfoods) != 1 {
			panic(fmt.Sprintf("expected 1, got %d", len(sfoods)))
		}
		delete(foods, sfoods[0])
	}

	var total int
	for _, count := range foods {
		total += count
	}
	return total
}

func Puzzle21() int {
	data := readFile("input/input21")
	return puzzle21(data)
}

func puzzle21b(data []string) string {
	allergens := make(map[string][]string)
	foods := make(map[string]int)
	expr := regexp.MustCompile(`^(.*) \(contains (.*)\)$`)
	for _, line := range data {
		matches := expr.FindStringSubmatch(line)
		sfoods := strings.Split(matches[1], " ")
		for _, food := range sfoods {
			foods[food]++
		}
		for _, allergen := range strings.Split(matches[2], ", ") {
			if cfoods, ok := allergens[allergen]; ok {
				allergens[allergen] = intersect(sfoods, cfoods)
				continue
			}
			allergens[allergen] = sfoods
		}
	}

	for {
		var changed bool
		for i, sfoods := range allergens {
			if len(sfoods) == 1 {
				for j, cfoods := range allergens {
					if i == j {
						continue
					}
					var count int
					allergens[j], count = remove(cfoods, sfoods[0])
					if count > 0 {
						changed = true
					}
				}
			}
		}
		if !changed {
			break
		}
	}

	danger := make([][2]string, 0, len(allergens))
	for allergen, sfoods := range allergens {
		if len(sfoods) != 1 {
			panic(fmt.Sprintf("expected 1, got %d", len(sfoods)))
		}
		danger = append(danger, [2]string{allergen, sfoods[0]})
	}
	sort.Slice(danger, func(i, j int) bool {
		return danger[i][0] < danger[j][0]
	})
	rdanger := make([]string, len(danger))
	for i, d := range danger {
		rdanger[i] = d[1]
	}
	return strings.Join(rdanger, ",")
}

func Puzzle21b() string {
	data := readFile("input/input21")
	return puzzle21b(data)
}
