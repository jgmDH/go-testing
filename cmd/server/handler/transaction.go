package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bootcamp-go/internal/transactions"
	"github.com/bootcamp-go/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Codigo   string  `json:"codigo"`
	Monto    float64 `json:"monto"`
	Moneda   string  `json:"moneda"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{service: s}
}

func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		transactions, err := t.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error())) // 500
			return
		}

		if len(transactions) <= 0 {
			c.JSON(http.StatusOK, web.NewResponse(200, "No se encontraron transacciones", "")) // 500
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transactions, ""))
	}
}

func (t *Transaction) GetAllHttp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		token := r.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			message, err := json.Marshal(web.NewResponse(401, nil, "token inválido"))
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, err = w.Write(message)
			if err != nil {
				panic(err)
			}
			return
		}

		transactions, err := t.service.GetAll()
		if err != nil {
			message, err := json.Marshal(web.NewResponse(500, nil, "token inválido"))
			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write(message)
			if err != nil {
				panic(err)
			}
			return
		}

		dataJson, err := json.Marshal(web.NewResponse(200, transactions, ""))
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-type", "application/json")
		if _, err := w.Write(dataJson); err != nil {
			panic(err)
		}
	}
}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error())) //400
			return
		}

		transaction, err := t.service.Store(req.Codigo, req.Moneda, req.Emisor, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error())) //500
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transaction, ""))
	}
}
