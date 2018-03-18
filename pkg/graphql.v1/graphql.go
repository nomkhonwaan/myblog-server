package graphql

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	graphqlgo "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/generated"
	"github.com/nomkhonwaan/myblog-server/pkg/graphql/resolver"
)

type GraphQL struct {
	Schema *graphqlgo.Schema
}

func New() *GraphQL {
	return &GraphQL{
		Schema: graphqlgo.MustParseSchema(Schema(), resolver.New()),
	}
}

func (h *GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(generated.MustAsset("pkg/graphql/graphiql/graphiql.html"))

		return
	} else if r.Method == http.MethodPost {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := h.Schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)

		responsJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.Write(responsJSON)

		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func Schema() string {
	buf := bytes.Buffer{}

	for _, s := range generated.AssetNames() {
		if strings.HasPrefix(s, "pkg/graphql/schema") {
			buf.Write(generated.MustAsset(s))
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
