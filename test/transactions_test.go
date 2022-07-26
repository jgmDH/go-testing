package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/bootcamp-go/cmd/server/handler"
	"github.com/bootcamp-go/internal/transactions"
	"github.com/bootcamp-go/pkg/store"
	"github.com/bootcamp-go/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var s = createServer()

func createServer() *gin.Engine {
	err := os.Setenv("TOKEN", "123456")
	if err != nil {
		panic(err)
	}

	db := store.NewFileStore(store.FileType, "./transactions.json")
	repo := transactions.NewRepository(db)
	serv := transactions.NewService(repo)

	ts := handler.NewTransaction(serv)

	// Para evitar que se vean las rutas en los log de tests
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/transactions", ts.GetAll())
	r.PUT("/transactions/:id", ts.Update())
	r.POST("/transactions", ts.Store())
	r.DELETE("/transactions/:id", ts.Delete())

	return r
}

func createServerHTTP() *http.Server {
	db := store.NewFileStore(store.FileType, "./transactions.json")
	repo := transactions.NewRepository(db)
	serv := transactions.NewService(repo)

	ts := handler.NewTransaction(serv)

	mux := http.NewServeMux()
	mux.HandleFunc("/transactions", ts.GetAllHttp())
	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	return &server
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", os.Getenv("TOKEN"))

	return req, httptest.NewRecorder()
}

func Test_GetTransactions_Ok(t *testing.T) {
	req, rw := createRequestTest(http.MethodGet, "/transactions", "")
	s.ServeHTTP(rw, req)

	objRes := &web.Response{}
	assert.Equal(t, 200, rw.Code)
	err := json.Unmarshal(rw.Body.Bytes(), &objRes)

	data := reflect.ValueOf(objRes.Data).Len() // Obteniendo la cantidad de transactions de Data
	assert.Nil(t, err)
	assert.True(t, data > 0)
}

func Test_GetTransactionsHTTP_Ok(t *testing.T) {
	req, rw := createRequestTest(http.MethodGet, "/transactions", "")
	srv := createServerHTTP()
	srv.Handler.ServeHTTP(rw, req)

	assert.Equal(t, 200, rw.Code)
	objRes := &web.Response{}
	err := json.Unmarshal(rw.Body.Bytes(), &objRes)

	require.NotNil(t, objRes.Data)
	data := reflect.ValueOf(objRes.Data).Len() // Obteniendo la cantidad de transactions de Data
	assert.Nil(t, err)
	assert.True(t, data > 0)
}

func Benchmark_GetTransactions_Ok(b *testing.B) {
	req, rw := createRequestTest(http.MethodGet, "/transactions", "")
	//srv := createServerHTTP()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		//srv.Handler.ServeHTTP(rw, req)
		s.ServeHTTP(rw, req)
	}
}

func Test_StoreTransactions_Ok(t *testing.T) {
	body := `{"emisor":"Evelin Torres", "receptor":"Joseline Charls", "monto":1233.99, "moneda":"dollar", "codigo":"134sdfs"}`
	req, rw := createRequestTest(http.MethodPost, "/transactions", body)

	s.ServeHTTP(rw, req)

	assert.Equal(t, 200, rw.Code)
}

func Test_UpdateTransactions_Ok(t *testing.T) {
	body := `{"emisor":"José María Alonso", "receptor":"Alvaro José", "monto":1233.99, "moneda":"dollar", "codigo":"134sdfs"}`
	url := fmt.Sprintf("/transactions/%d", 1)
	req, rw := createRequestTest(http.MethodPut, url, body)

	s.ServeHTTP(rw, req)
	assert.Equal(t, 200, rw.Code)
}

func Test_DeleteTransactions_Ok(t *testing.T) {
	// s := createServer()
	store := store.NewFileStore(store.FileType, "./transactions.json")
	ts := []*transactions.Transaction{{Id: 1, Codigo: "abc123", Moneda: "peso", Emisor: "Juan Manuel", Receptor: "Gissel Rivas", Monto: 129.99}}
	store.Write(ts)

	url := fmt.Sprintf("/transactions/%d", 1)
	req, rw := createRequestTest(http.MethodDelete, url, "")
	s.ServeHTTP(rw, req)

	assert.Equal(t, 200, rw.Code)

	// Creamos la transaction eliminada para poder ejercutar el test las veces que sean necesarias.
	store.Write(ts)
}
