package main

import (
    "fmt"
    "os/user"
    "os/exec"
    "os"
    "io/ioutil"
)

const TTL = 15 //days
var CACHEDIR string = getHomeDir()+".cache/pkmncs/"

func getHomeDir () (string) {
    user, _ := user.Current()
    username := user.Username
    return ("/home/"+username+"/")
}

func CacheFolderExists () (bool){
    _, err := os.Stat(CACHEDIR)
    if os.IsNotExist(err) {
        return false
    }

    return true
}

func CreateCacheFolder () {
    if !CacheFolderExists() {
        cmd := exec.Command("mkdir", "pkmncs")
        cmd.Dir = getHomeDir()+".cache/"
        cmd.Run()
    }
}

func CacheExists (content string) (bool) {
    _, err := os.Stat(CACHEDIR+content+".json")
    if os.IsNotExist(err) {
        return false
    }

    return true
}

func GetCache (content string) (string) {
    cacheFile := CACHEDIR+content+".json"
    cacheByte, err := ioutil.ReadFile(cacheFile)

    if err != nil{
        fmt.Println("Error while reading cache:", err)
        os.Exit(1)
        
    }
    return string(cacheByte)
}

func main (){
}

// Se a pasta de cache não existir, criar
// Se existir, buscar cache
// Se cache tiver expirado, recriar
// Senao, importar com base no cache
