package tax

import "testing"

// O coverage serve para testar a cobertura de testes
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
func TestCalculateTax(test *testing.T) {
    amount := 500.0
    expected := 5.0

    result := CalculateTax(amount)

    if result != expected {
        test.Errorf("Expected %f but got %f", expected, result)
    }
}

func TestCalculateTaxBatch(test *testing.T) {
    type calcTax struct {
        amount, expect float64
    }
        table := []calcTax{
        {500.0, 5.0},
        {1000.0, 10.0},
        {1500.0, 10.0},
        {0, 0.0},
    }
    for _, item := range table {
        result := CalculateTax(item.amount)
        if result != item.expect {
            test.Errorf("Expected %f but got %f", item.expect, result)
        }
    }
}
// O benchmark serve para testar a performance da função
// go test -bench=.
// go test -bench=. -count 5 -run=^# -benchtime=5s
func BenchmarkCalculateTax(benchmark *testing.B) {
    for i := 0; i < benchmark.N; i++ {
        CalculateTax(500.0)
    }
}

func BenchmarkCalculateTax2(benchmark *testing.B) {
    for i := 0; i < benchmark.N; i++ {
        CalculateTax2(500.0)
    }
}

// O fuzz serve para testar a função com valores aleatórios
// go test -fuzz=. -fuzztime=5s
func FuzzCalculateTax(fuzz *testing.F) {
   seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
   for _, amount := range seed {
    fuzz.Add(amount)
   }
   fuzz.Fuzz(func(test *testing.T, amount float64) {
    result := CalculateTax(amount)
    if amount <= 0 && result != 0 {
        test.Errorf("Reveived %f but expected 0", result)
    }
    if amount > 20000 && result != 5 {
        test.Errorf("Reveived %f but expected 5", result)
    }

    // if amount > 20000 && result != 20 {
    //     test.Errorf("Reveived %f but expected 20", result)
    // }
   })
}