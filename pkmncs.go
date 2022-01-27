package main

import ( 
    "io/ioutil"
    "net/http"
    "fmt"
    "encoding/json"
    "os"
    "sort"
    "pkmncs/helper"
)

//TODO: create help function

//TODO: analyze when the listed pokemons can learn the desired moves

// getMoveLearnedBy receives a byte response from a API call and convert it and returns a map. 
func getMoveLearnedBy (responseData []byte) (pkmnList []string) {
    pkmnMap := make (map[string]([]map[string]string))
    json.Unmarshal (responseData, &pkmnMap)
    learnedBy := pkmnMap["learned_by_pokemon"]

    for key := range learnedBy {
        pkmnList = append(pkmnList, learnedBy[key]["name"])
    }

    sort.Strings (pkmnList)

    return pkmnList
}

func getMoveInfo (move string) ([]string) {
    moveData := "https://pokeapi.co/api/v2/move/" + move

    moveByte, callError  :=  http.Get(moveData)

    if callError != nil || moveByte.StatusCode == 404 {
        fmt.Println(callError)
        fmt.Printf ("'%s' move not found. Check spelling.\n", move)
        os.Exit ( 1 )
    }

    responseData, readError := ioutil.ReadAll(moveByte.Body)

    if readError != nil{
        panic ("Error while reading byte body")
    }

     return getMoveLearnedBy (responseData)
}
// parseArgs deals with system args to pkmncs
func parseArgs (){
    var candidates []string

    if len(os.Args) <= 1 {
        println("You should specify at least one move")
        os.Exit(1)
    }

    for i := 1; i < len(os.Args); i++{
        pkmnnList :=  getMoveInfo(os.Args[i])
        
        if i == 1{
            candidates = make([]string, len(pkmnnList))
            copy(candidates, pkmnnList)
           
        }else{
             candidates = parseIntersections (candidates, pkmnnList)
         }
    }
    helper.PrintSlice(candidates)   
}

//  parseIntersections identifies intersections between two slices. Returns a slice with the intersections.
func parseIntersections(candidates []string, newMoveList []string) ([]string) {

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

func main (){
    parseArgs()
}
