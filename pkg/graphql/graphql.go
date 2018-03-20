package graphql

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	ggg "github.com/graph-gophers/graphql-go"
	"github.com/nomkhonwaan/myblog-server/pkg/generated"
	"github.com/nomkhonwaan/myblog-server/pkg/graphql/resolver"
)

// Handler is an HTTP handler for GraphQL which serves the following these
// - [GET] /graphql - Returns a Graphiql HTML page
// - [POST] /graphql - Serves a GraphQL endpoint
type Handler struct {
	*resolver.Resolver `inject:"pkg/graphql/resolver.Resolver"`
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	schema := ggg.MustParseSchema(
		string(Schema()),
		h.Resolver,
	)

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(Graphiql())
		return
	case http.MethodPost:
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
		respJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Accept", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

// Graphiql returns a Graphiql HTML page as a byte array
func Graphiql() []byte {
	return generated.MustAsset("pkg/graphql/graphiql/index.html")
}

// Schema returns a GraphQL's schema as a byte array
func Schema() []byte {
	buf := bytes.Buffer{}
	for _, s := range generated.AssetNames() {
		if strings.HasPrefix(s, "pkg/graphql/schema") {
			buf.Write(generated.MustAsset(s))
			buf.WriteByte('\n')
		}
	}
	return buf.Bytes()
}
