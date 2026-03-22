package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f bit got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})
}

/*
	Para rodar o teste, use o comando:
		- go mod test .

	E também temos o comando:
		- go test -v

*/

/*
	Comando:
		- go test -fuzz=. -run=^#

	Resultado:
	fuzz: elapsed: 0s, gathering baseline coverage: 0/6 completed
	fuzz: elapsed: 0s, gathering baseline coverage: 6/6 completed, now fuzzing with 8 workers
	fuzz: elapsed: 0s, execs: 331 (2095/sec), new interesting: 1 (total: 7)
	--- FAIL: FuzzCalculateTax (0.17s)
		--- FAIL: FuzzCalculateTax (0.00s)
			tax_test.go:58: Reveived 5.000000 but expected 0

		Failing input written to testdata\fuzz\FuzzCalculateTax\5fb97e24f60a8962
		To re-run:
		go test -run=FuzzCalculateTax/5fb97e24f60a8962
	FAIL
	exit status 1
	FAIL    taxgo/1 0.458s

*/

/*
	Comandos:
		- go test -fuzz=. -fuzztime=5s -run=^#
*/
