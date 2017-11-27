package algo

import (
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"testing"
)

var (
	// prod prime[i] for i from 1000000 to 1000005
	n6, _ = new(big.Int).SetString("13791771367562127265550426827257002816955167", 10)
	// prod prime[i] for i from 10000000 to 10000005
	n7, _ = new(big.Int).SetString("33365182390916237933213738438338631228810007682817", 10)
	// prod prime[i] for i from 100000000 to 100000005
	n8, _ = new(big.Int).SetString("71667239492150606021054927440744481129086374361123029323", 10)
	// prod prime[i] for i from 1000000001 to 1000000005
	n9, _ = new(big.Int).SetString("1018228583743274635502006979268459963072627838100972471581", 10)
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func benchmarkParallel(n *big.Int, j int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorizeParallel(n, j)
	}
}

func BenchmarkFactorizeParallel6_1(b *testing.B) { benchmarkParallel(n6, 1, b) }
func BenchmarkFactorizeParallel6_2(b *testing.B) { benchmarkParallel(n6, 2, b) }
func BenchmarkFactorizeParallel6_3(b *testing.B) { benchmarkParallel(n6, 3, b) }
func BenchmarkFactorizeParallel6_4(b *testing.B) { benchmarkParallel(n6, 4, b) }

func BenchmarkFactorizeParallel7_1(b *testing.B) { benchmarkParallel(n7, 1, b) }
func BenchmarkFactorizeParallel7_2(b *testing.B) { benchmarkParallel(n7, 2, b) }
func BenchmarkFactorizeParallel7_3(b *testing.B) { benchmarkParallel(n7, 3, b) }
func BenchmarkFactorizeParallel7_4(b *testing.B) { benchmarkParallel(n7, 4, b) }

func BenchmarkFactorizeParallel8_1(b *testing.B) { benchmarkParallel(n8, 1, b) }
func BenchmarkFactorizeParallel8_2(b *testing.B) { benchmarkParallel(n8, 2, b) }
func BenchmarkFactorizeParallel8_3(b *testing.B) { benchmarkParallel(n8, 3, b) }
func BenchmarkFactorizeParallel8_4(b *testing.B) { benchmarkParallel(n8, 4, b) }

func BenchmarkFactorizeParallel9_1(b *testing.B) { benchmarkParallel(n9, 1, b) }
func BenchmarkFactorizeParallel9_2(b *testing.B) { benchmarkParallel(n9, 2, b) }
func BenchmarkFactorizeParallel9_3(b *testing.B) { benchmarkParallel(n9, 3, b) }
func BenchmarkFactorizeParallel9_4(b *testing.B) { benchmarkParallel(n9, 4, b) }
