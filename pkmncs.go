package main

import ( 
    "io/ioutil"
    "net/http"
    "fmt"
    "encoding/json"
    "os"
    "sort"
)

func getMoveGet (  ){
    panic ( "getMoveGet not implemented yet" )
}

// getMoveLearnedBy receives a byte response from a API call and convert it and returns a map. 
func getMoveLearnedBy ( responseData []byte ) ( pkmnList []string ) {
    pkmnMap := make ( map[string]( []map[string]string ) )
    json.Unmarshal ( responseData, &pkmnMap )
    learnedBy := pkmnMap["learned_by_pokemon"]
    
    // var pkmnList []string

    for key := range learnedBy {
        pkmnList = append(pkmnList, learnedBy[key]["name"])
    }

    sort.Strings ( pkmnList )

    fmt.Println ( pkmnList )

    return pkmnList
}

func getMoveInfo () {
    moveData := "https://pokeapi.co/api/v2/move/" + os.Args[1]

    moveByte, callError  :=  http.Get ( moveData )

    if callError != nil || moveByte.StatusCode == 404 {
        fmt.Printf ( "'%s' move not found. Check spelling.\n", os.Args[1] )
        os.Exit ( 1 )
    }

    responseData, readError := ioutil.ReadAll( moveByte.Body )

    if readError != nil{
        panic ( "Error while reading byte body" )
    }

     getMoveLearnedBy ( responseData )


}
 
func main (){
    getMoveInfo()
}
