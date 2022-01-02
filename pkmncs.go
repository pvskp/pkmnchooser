package main

import ( 
    "io/ioutil"
    "net/http"
    "fmt"
    "encoding/json"
    "os"
)

// getMoveInfo returns a slice of maps (pokemons) that can learn
// the move specified in the parameters.
func getMoveInfo (  ) {
    //TODO: Parse via "learned by"
    //TODO: Parse via "learn via TM/HM"
    //TODO: Parse via "egg move"
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

    pkmnMap := make ( map[string]( []map[string]string ) )
    json.Unmarshal ( responseData, &pkmnMap )
    learnedBy := pkmnMap["learned_by_pokemon"]
    
    for key := range learnedBy { // list all pokémons that can learn the move
        fmt.Printf( "%s ", learnedBy[key]["name"] )
    }
}
 
func main (){
    getMoveInfo()
}
