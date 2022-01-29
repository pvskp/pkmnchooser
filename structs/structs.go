package structs

type LearnedBy struct{
    Pokemons []struct{
        Name string `json:"name"`
    }`json:"learned_by_pokemon"`
}

type Pokemon struct{
    Moves []struct {
        Move struct{
            Name string `json:"name"`
        } `json:"move"`
    }`json:"moves"`
}


