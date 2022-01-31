package cache

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
    _, err := os.Stat(CACHEDIR+content)
    if os.IsNotExist(err) {
        return false
    }

    return true
}

func GetCache (content string) ([]byte) {
    cacheFile := CACHEDIR+content
    cacheByte, err := ioutil.ReadFile(cacheFile)

    if err != nil{
        fmt.Println("Error while reading cache:", err)
        os.Exit(1)
        
    }
    return cacheByte
}

func CacheContent (filename string, content []byte) {
    os.WriteFile(CACHEDIR+filename, content, 0700)
}
