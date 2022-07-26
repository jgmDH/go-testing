package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

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
		w.Write(dataJson)
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

func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		if req.Codigo == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo codigo es requerido")) //400
			return
		}

		if req.Emisor == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo emisor es requerido")) //400
			return
		}

		if req.Receptor == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo receptor es requerido")) //400
			return
		}

		if req.Moneda == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo moneda es requerido")) //400
			return
		}

		if req.Monto <= 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "El campo monto es requerido")) //400
			return
		}

		transaction, err := t.service.Update(id, req.Codigo, req.Moneda, req.Emisor, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error())) //404
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transaction, ""))
	}
}

func (t *Transaction) UpdateReceptorYMonto() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		if req.Receptor == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		if req.Monto <= 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "")) //400
			return
		}

		transaction, err := t.service.UpdateReceptorYMonto(id, req.Receptor, req.Monto)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, "")) //404
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, transaction, ""))
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error())) //400
			return
		}

		err = t.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error())) //404
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(200, nil, ""))
	}
}
