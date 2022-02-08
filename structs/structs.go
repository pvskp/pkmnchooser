package structs

type LearnedBy struct{
    Pokemons []struct{
        Name string `json:"name"`
    }`json:"learned_by_pokemon"`
}

type Pokemon struct{
    Moves []struct{
        Move struct{
            Name string `json:"name"`
        } `json:"move"`
        
        VersionDetails []struct{
            LevelLearnedAt int `json:"level_learned_at"`

            LearnMethod struct{
                Name string `json:"name"`
            }`json:"move_learn_method"`

            VersionGroup struct{
                Name string `json:"name"`
            }`json:"version_group"`

        }`json:"version_group_details"`

    }`json:"moves"`
}
