package helper

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func ApiConsume(endpoint string) ([]byte, error) {
    // TODO: Implement returns to errors
    r, err := http.Get(endpoint)
    
    if err != nil{
        panic("Error while consuming on endpoint"+endpoint)
    }
    
    responseData, err := ioutil.ReadAll(r.Body)

    if err != nil{
        panic("Error while parsing byte data on API endpoint" + endpoint)
    }
    
    return responseData, nil
}


func PrintSlice (slice []string) {
    if (len(slice) == 0) {
        fmt.Println("No results")
        return
    }

    for i := 0; i < len(slice); i++ {
        fmt.Printf("%s\n", slice[i])
    }
}

// BinarySearch checks if a specified name is in a vector/slice. If it is, returns true. Else, returns false.
func BinarySearch (name string, list []string, start int, end int) (cotains bool) {

    mid := ( start + end ) / 2

    if start <= end {

        if list[mid] == name {
            return true
        }
        
        if list[mid] > name {
            return BinarySearch(name, list, start, mid - 1)
        }

        return BinarySearch(name, list, mid + 1, end)
    }

    return false
}

