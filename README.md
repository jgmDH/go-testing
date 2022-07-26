### Linters - Golangci-lint

Vet examina el código fuente de Go e informa construcciones sospechosas, como Printf llamadas cuyos argumentos no se alinean con la cadena de formato. https://pkg.go.dev/cmd/vet

```go
go vet  
// o si hay varios directorios
go vet ./...
```

Staticcheck, utilizando la recomendación de `go install honnef.co/go/tools/cmd/staticcheck@latest` https://staticcheck.io/
```go
// ejecuta este comando para analizar todas los paquetes.
staticcheck ./... 
```

Golangci-lint, lo podemos instalar siguiendo las instrucciones de su web https://golangci-lint.run/usage/install/. También importante es que es posible aplicar Github Actions con golangci para nuestro proyecto.
```go
golangci-lint run
```

<hr>

### Coverage

Mediante la herramienta que nos provee go, podemos mediante la ejecución de ciertos comando saber que porcentaje de codigo se encuentra testeado.

```go 
// Obtenemos un coverage de los paguetes que contienen test
go test -v ./... -cover
```



<hr>

### Benchmarks

Los benchmark nos permiten realizar pruebas que nos facilitan ver el rendimiento de uno o varios procesos en nuestra aplicación, y así detectar posibles mejoras de rendimiento. https://pkg.go.dev/cmd/go#hdr-Testing_flags

Para ejecutar los benchmark de un solo paquete:

```go
go test -bench=. github.com/bootcamp-go/conversion
goos: linux
goarch: amd64
pkg: github.com/bootcamp-go/conversion
cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
BenchmarkFmtSprint-8    21264790                54.39 ns/op
BenchmarkStrconv-8      63416014                17.21 ns/op
PASS
ok      github.com/bootcamp-go/conversion       3.868s
```

Para ejecutar todos los benchmarks del Proyecto:
```go
go test -v -bench=. ./... 
// o para ver detalles de memoria
go test -v -bench=. ./... -benchmem
```


```go
go test -bench=BenchmarkFibonacci$ github.com/bootcamp-go/testing/fibonacci -count=5 > old.txt 
go test -bench=BenchmarkFibonacci$ github.com/bootcamp-go/testing/fibonacci -count=5 > new.txt

// benchstat [options] old.txt [new.txt] 
// Tool para comparar benchmarks 
go install golang.org/x/perf/cmd/benchstat@latest
benchstat old.txt new.txt
// output:
Fibonacci-8   274ns ± 1%     6ns ± 1%  -97.72%  (p=0.008 n=5+5)
```

### Tips:

    - Using higher -count values if the benchmark numbers aren't stable
    - Usingn -benchmem to also get stats on allocated objects and space
    - Running the benchmarks on an idle machine not running on battery (and with power management off?)
    - Adding -run='$^' or -run=- to each go test command to avoid running the tests too

Links utiles: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go