package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pkmncs/cache"
	"pkmncs/helper"
	"pkmncs/structs"
	"sort"
)

//TODO: create help function
func _help() {
	//print instructions
	panic("_help not implemented yet.")
}

//TODO: analyze when the listed pokemons can learn the desired moves
func getPokeInfo(pokemonList []string) {
	var pkmnStruct structs.Pokemon
	for _, pokemon := range pokemonList {
		if cache.CacheExists(pokemon) {
			cache := cache.GetCache(pokemon)
			json.Unmarshal(cache, &pkmnStruct)

		} else {
			r, _ := helper.ApiConsume("https://pokeapi.co/api/v2/pokemon/" + pokemon)
			json.Unmarshal(r, &pkmnStruct)
			cache.CacheContent(pokemon, r)
		}

		// for _, m := range(pkmnStruct.Moves){
		// fmt.Println(m.Move.Name)
		//     for _, n := range(m.VersionDetails) {
		//         if n.VersionGroup.Name == "ultra-sun-ultra-moon" {
		//             fmt.Println("\t", n.LearnMethod.Name, n.VersionGroup.Name)
		//         }
		//     }
		// }
	}
}

// getMoveLearnedBy receives a byte response from a API call and convert it and returns a struct.
func getMoveLearnedBy(responseData []byte) (pkmnList []string) {
	var pkmnLearnedBy structs.LearnedBy
	json.Unmarshal(responseData, &pkmnLearnedBy)
	for _, key := range pkmnLearnedBy.Pokemons {
		pkmnList = append(pkmnList, key.Name)
	}
	sort.Strings(pkmnList)
	return pkmnList
}

func getMoveInfo(move string) []string {
	if cache.CacheExists(move) {
		// fmt.Printf("Using cache for %s!\n", move)
		return getMoveLearnedBy(cache.GetCache(move))
	}
	moveData := "https://pokeapi.co/api/v2/move/" + move
	responseData, _ := helper.ApiConsume(moveData)
	cache.CacheContent(move, responseData)
	return getMoveLearnedBy(responseData)
}

// parseArgs deals with system args to pkmncs
func parseArgs() {
	var candidates []string
	if len(os.Args) <= 1 {
		println("You should specify at least one move")
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		pkmnnList := getMoveInfo(os.Args[i])

		if i == 1 {
			candidates = make([]string, len(pkmnnList))
			copy(candidates, pkmnnList)

		} else {
			candidates = parseIntersections(candidates, pkmnnList)
		}
	}
	helper.PrintSlice(candidates)
}

//  parseIntersections identifies intersections between two slices. Returns a slice with the intersections.
func parseIntersections(candidates []string, newMoveList []string) []string {
	newCandidates := make([]string, len(candidates))
	j := 0
	for i := 0; i < len(candidates); i++ {
		// conferir se o dado elemento da linha está presente na nova linha
		if helper.BinarySearch(candidates[i], newMoveList, 0, len(newMoveList)-1) {
			// remover elementos que não estão presentes nas duas linhas
			newCandidates[j] = candidates[i]
			j++
		}
	}
	newCandidates = newCandidates[:j]
	return newCandidates
}

func main() {
	// parseArgs()
	// parseLearnedAt("pikachu")
	pkmnList := []string{"pikachu", "ditto"}
	getPokeInfo(pkmnList)
}
