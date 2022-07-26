package domains

type Transaction struct {
	Id       int64
	Codigo   string
	Moneda   string
	Emisor   string
	Receptor string
	Monto    float64
}
