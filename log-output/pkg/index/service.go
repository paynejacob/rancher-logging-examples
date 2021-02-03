package index

import (
	"bufio"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
	"text/template"
)

type Service struct {
	m       *sync.RWMutex
	indices map[string]Index
}

func NewIndexService() *Service {
	return &Service{
		m:       &sync.RWMutex{},
		indices: make(map[string]Index),
	}
}

func (s Service) ListLogs(w http.ResponseWriter, r *http.Request) {
	var err error
	var logs [][]byte
	var indexName = mux.Vars(r)["index_name"]

	s.m.RLock()

	// check if this is a valid index
	if _, exists := s.indices[indexName]; exists {
		logs = s.indices[indexName].Logs
	}

	s.m.RUnlock()

	// write logs to response
	for i := range logs {
		if _, err = w.Write(logs[i]); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if _, err = w.Write([]byte("\n")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (s Service) WriteLog(w http.ResponseWriter, r *http.Request) {
	var index Index
	var indexName = mux.Vars(r)["index_name"]

	s.m.Lock()
	defer s.m.Unlock()
	// check if this index exists
	if _, exists := s.indices[indexName]; exists {
		index = s.indices[indexName]
	} else {
		index = NewIndex(indexName)
	}

	scanner := bufio.NewScanner(r.Body)
	for scanner.Scan() {
		WriteLog(&index, []byte(scanner.Text()))
	}

	s.indices[indexName] = index
}

func (s Service) ListIndices(w http.ResponseWriter, r *http.Request) {
	var indices []Index

	s.m.RLock()

	for _, index := range s.indices {
		indices = append(indices, index)
	}

	s.m.RUnlock()

	_template, err := template.New("listPage").Funcs(template.FuncMap{"DateFormat": DateFormat}).Parse(listPageTemplate)
	_ = err
	_ = _template.Execute(w, indices)
	w.Header().Add("Content-Type", "text/html")
}

