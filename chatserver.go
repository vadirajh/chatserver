package main 
import (
	"encoding/json"
	"net/http"
	"time"
)

// Structure to hold all messages
type Result struct {
	Timestamp int64  `json:"timestamp"`
	User      string `json:"user"`
	Text      string `json:"text"`
}
type Results []Result
type ResultJson struct {
	Messages Results `json:"messages"`
}

//Structure to unmarshal POST /message
type IncomingMessage struct {
	User string `json:"user"`
	Text string `json:"text"`
}

//JSON ready slice of users
type UniqueUsers struct {
	Username []string `json:"users"`
}

//Use a map to add to unique users
func AppendIfUnique(slice []string, user string) []string {
	if _, ok := uu_map[user]; ok {
		return slice
	}
	uu_map[user] = struct{}{}
	return append(slice, user)
}
const GetMaxMessagesLen = 100
func messagesHandler(c http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(c, "405 Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	c.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(c)
	enc.SetIndent("", "    ")
	resultsLen := len(results.Messages)
	if resultsLen < GetMaxMessagesLen {
		if err := enc.Encode(results); err != nil {
			panic(err)
		}
	} else {
		if err := enc.Encode(ResultJson{results.Messages[resultsLen-GetMaxMessagesLen:]}); err != nil {
			panic(err)
		}
	}
}
func messagePostHandler(c http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(c, "405 Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var m IncomingMessage
	if err := json.NewDecoder(req.Body).Decode(&m); err != nil {
		http.Error(c, err.Error(), http.StatusBadRequest)
		return
	}
	results.Messages = append(results.Messages, Result{Timestamp: time.Now().Unix(), User: m.User, Text: m.Text})
	// Add to list of unique users
	uu.Username = AppendIfUnique(uu.Username, m.User)
	enc := json.NewEncoder(c)
	enc.SetIndent("", "    ")
	success := map[string]bool{"ok": true}
	if err := enc.Encode(success); err != nil {
		panic(err)
	}
}

func usersHandler(c http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(c, "405 Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	c.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(c)
	enc.SetIndent("", "    ")
	if err := enc.Encode(uu); err != nil {
		panic(err)
	}

}

func makeMap() map[string]struct{} {
	return make(map[string]struct{})
}

var results ResultJson
var uu UniqueUsers
var uu_map map[string]struct{} = makeMap()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status",
		func(c http.ResponseWriter, req *http.Request) {
			c.Write([]byte("alive"))
		})
	mux.HandleFunc("/messages", messagesHandler)
	mux.HandleFunc("/message", messagePostHandler)
	mux.HandleFunc("/users", usersHandler)
	err := http.ListenAndServe(":8081", mux)
	panic(err)
}
