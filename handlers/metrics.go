// handlers/metrics.go
package handlers

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/Viru9029/URL-shortener-golang/storage"
)

func Metrics(storage *storage.MemoryStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		counts := storage.GetTopDomains()
		type kv struct {
			Key   string
			Value int
		}
		var ss []kv
		for k, v := range counts {
			ss = append(ss, kv{k, v})
		}
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Value > ss[j].Value
		})

		top3 := map[string]int{}
		for i := 0; i < len(ss) && i < 3; i++ {
			top3[ss[i].Key] = ss[i].Value
		}

		json.NewEncoder(w).Encode(top3)
	}
}
