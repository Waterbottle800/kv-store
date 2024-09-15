package main
import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	"sync"
	"log"
)

type Store struct {
	m sync.RWMutex
	data map[string]string
}

func InitKV() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (store *Store) Set(key, value string) {
	store.m.Lock()
	defer store.m.Unlock()
	store.data[key]=value
}

func (store *Store) Get(key string) (string, bool) {
	store.m.RLock()
	defer store.m.RUnlock()
	k,v := store.data[key]
	return k,v
}

func (store *Store) Search(prefix, suffix string) []string {
	store.m.RLock()
	defer store.m.RUnlock()

	var result []string
	for key:= range store.data {
		if (prefix == "" || strings.HasPrefix(key,prefix)) && (suffix == "" || strings.HasSuffix(key,suffix)) {
			result = append(result,key)

		}

	}
	return result
}

func main() {
	
	kv := InitKV()

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}
	
		var requestData map[string]string
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	
		key, keyExists := requestData["key"]
		value, valueExists := requestData["value"]
	
		if !keyExists || !valueExists {
			http.Error(w, "Key and value are required", http.StatusBadRequest)
			return
		}
	
		kv.Set(key, value)
		fmt.Fprintf(w, "Key %s set successfully", key)
	})

	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.TrimPrefix(r.URL.Path, "/get/")
		if key == "" {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}
	
		value, exists := kv.Get(key)
		if !exists {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}
	
		fmt.Fprintf(w, "Value: %s", value)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		prefix := r.URL.Query().Get("prefix")
		suffix := r.URL.Query().Get("suffix")
	
		results := kv.Search(prefix, suffix)
	
		json.NewEncoder(w).Encode(results)
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

	