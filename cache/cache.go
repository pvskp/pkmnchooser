package cache

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
)

const TTL = 15 //days
var CACHEDIR string = getHomeDir() + ".cache/pkmncs/"

// getHomeDir returns the user's home directory path. For example: /home/user/
func getHomeDir() string {
	user, _ := user.Current()
	username := user.Username
	return ("/home/" + username + "/")
}

// CacheFolderExists identifies if a pkmncs cache folder exists on ~/.cache.
func CacheFolderExists() bool {
	_, err := os.Stat(CACHEDIR)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateCacheFolder creates a cache folder for pkmncs on ~/.cache
func CreateCacheFolder() {
	if !CacheFolderExists() {
		cmd := exec.Command("mkdir", "pkmncs")
		cmd.Dir = getHomeDir() + ".cache/"
		cmd.Run()
	}
}

// CacheExists verify if a desired content is cached for use.
func CacheExists(content string) bool {
	_, err := os.Stat(CACHEDIR + content)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

// GetCache parses a cached content.
func GetCache(content string) []byte {
	cacheFile := CACHEDIR + content
	cacheByte, err := ioutil.ReadFile(cacheFile)

	if err != nil {
		fmt.Println("Error while reading cache:", err)
		os.Exit(1)

	}
	return cacheByte
}

// CacheContent caches a content on ~/.cache
func CacheContent(filename string, content []byte) {
	os.WriteFile(CACHEDIR+filename, content, 0700)
}
