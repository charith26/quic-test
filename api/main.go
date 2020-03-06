package main

import (
    "log"
    "net/http"
    "math/rand"
    "time"
)
const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("{\n"))
    w.Write([]byte("\t\"key\": \""+ string(key) +"\",\n"))

    for i:=1;i<500;i++ {
      w.Write([]byte("\t\""+StringWithCharset(5, charset)+"\": \""+ StringWithCharset(255, charset) +"\",\n"))
    }

    w.Write([]byte("\t\""+StringWithCharset(5, charset)+"\": \""+ StringWithCharset(255, charset) +"\"\n"))
    w.Write([]byte("}"))
}

func main() {
    s := &server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
