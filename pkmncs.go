package main

import ( 
    "io/ioutil"
    "net/http"
    "fmt"
    "encoding/json"
    "os"
    "sort"
)

func printSlice (slice []string) {
    for i := 0; i < len(slice); i++ {
        fmt.Printf("%s ", slice[i])
    }
}

func sliceRemove (slice []string, index int) ([]string) {
    slice[index] = slice[len(slice)-1]
    slice[len(slice)-1] = ""
    return slice[:len(slice)-1]
}

// binarySearch checks if a specified name is in a vector/slice. If it is, returns true. Else, returns false.
func binarySearch (name string, list []string, start int, end int) (cotains bool) {

    mid := ( start + end ) / 2
    mid = int( mid )

    if start <= end {

        if list[mid] == name {
            return true
        }
        
        if list[mid] > name {
            return binarySearch(name, list, start, mid - 1)
        }

        return binarySearch(name, list, mid + 1, end)
    }

    return false
}

func getMoveGet (){
    panic ("getMoveGet not implemented yet")
}

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

func parseArgs (){
    var candidates []string

    if len(os.Args) <= 1 {
        println("You should specify at least one move")
        os.Exit(1)
    }

    for i := 1; i < len(os.Args); i++{
        pkmnnList :=  getMoveInfo(os.Args[i])
        
        if i == 1{
            copy(candidates, pkmnnList)
           
        }else{
            parseIntersections (candidates, pkmnnList)
        }
    }
    
}

func parseIntersections(candidates []string, newMoveList []string) {

    for i := 0; i < len(candidates); i++ {
    // conferir se o dado elemento da linha está presente na nova linha
        if !binarySearch(candidates[i], newMoveList , 0, len(newMoveList)) {
            // remover elementos que não estão presentes nas duas linhas
            candidates = sliceRemove(candidates, i)
        }
    }
}

func main (){
    parseArgs()
}
