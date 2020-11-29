package miller_rabin

import (
	"math/big"
	"reflect"
	"testing"
)

const Reps = 16

var one *big.Int
var two *big.Int
var ten *big.Int
var prime16 *big.Int
var prime39 *big.Int
var prime44 *big.Int
var prime79 *big.Int
var prime157 *big.Int
var prime969 *big.Int
var prime6987 *big.Int
var prime13395 *big.Int
var prime25962 *big.Int
var prime39751 *big.Int
var prime65050 *big.Int
var prime65087 *big.Int
var composite16 *big.Int
var composite83 *big.Int
var composite340 *big.Int
var composite1350 *big.Int
var composite20382 *big.Int
var composite65713 *big.Int

func TestMain(m *testing.M) {
	one = big.NewInt(1)
	two = big.NewInt(2)
	ten = big.NewInt(10)

	prime16 = big.NewInt(9007199254740997)
	prime39 = mersenne(127)

	prime44 = big.NewInt(2)
	prime44.Exp(prime44, big.NewInt(148), nil)
	prime44.Add(prime44, one)
	prime44.Div(prime44, big.NewInt(17))

	prime79 = mersenne(127)
	prime79.Exp(prime79, two, nil)
	prime79.Mul(prime79, big.NewInt(180))
	prime79.Add(prime79, one)

	prime157 = mersenne(521)
	prime969 = mersenne(3217)
	prime6987 = mersenne(23209)
	prime13395 = mersenne(44497)
	prime25962 = mersenne(86243)
	prime39751 = mersenne(132049)
	prime65050 = mersenne(216091)

	prime65087 = new(big.Int).Exp(two, big.NewInt(216193), nil)
	prime65087.Mul(prime65087, big.NewInt(391581))
	prime65087.Sub(prime65087, one)

	composite16 = big.NewInt(8635844967113809)
	composite83 = new(big.Int).Mul(prime39, prime44)
	composite340 = new(big.Int).Mul(mersenne(521), mersenne(607))
	composite1350 = new(big.Int).Mul(mersenne(2203), mersenne(2281))
	composite20382 = new(big.Int).Mul(prime6987, prime13395)
	composite65713 = new(big.Int).Mul(prime25962, prime39751)

	m.Run()
}

func TestIsPrime(t *testing.T) {
	type args struct {
		bigInt *big.Int
	}
	tests := []struct {
		desc string
		args args
		want bool
	}{
		{
			desc: "Case 1",
			args: args{bigInt: one},
			want: false,
		},
		{
			desc: "Case 2",
			args: args{bigInt: two},
			want: true,
		},
		{
			desc: "Case prime 16 digits",
			args: args{bigInt: prime16},
			want: true,
		},
		{
			desc: "Case prime 39 digits",
			args: args{bigInt: prime39},
			want: true,
		},
		{
			desc: "Case prime 44 digits",
			args: args{bigInt: prime44},
			want: true,
		},
		{
			desc: "Case prime 79 digits",
			args: args{bigInt: prime79},
			want: true,
		},
		{
			desc: "Case prime 157 digits",
			args: args{bigInt: prime157},
			want: true,
		},
		{
			desc: "Case prime 969 digits",
			args: args{bigInt: prime969},
			want: true,
		},
		{
			desc: "Case composite 16 digits",
			args: args{bigInt: composite16},
			want: false,
		},
		{
			desc: "Case composite 83 digits",
			args: args{bigInt: composite83},
			want: false,
		},
		{
			desc: "Case composite 340 digits",
			args: args{bigInt: composite340},
			want: false,
		},
		{
			desc: "Case composite 1350 digits",
			args: args{bigInt: composite16},
			want: false,
		},
		/*
			{
				desc: "Case composite 20382 digits",
				args: args{bigInt: composite20382},
				want: false,
			},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := tt.args.bigInt.ProbablyPrime(Reps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProbablyPrime = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrimeAll(t *testing.T) {
	type args struct {
		digit int64
	}
	tests := []struct {
		desc string
		args args
		want int64
	}{
		{
			desc: "25 / 10^2",
			args: args{digit: 2},
			want: 25,
		},
		{
			desc: "168 / 10^3",
			args: args{digit: 3},
			want: 168,
		},
		{
			desc: "1229 / 10^4",
			args: args{digit: 4},
			want: 1229,
		},
		{
			desc: "9592 / 10^5",
			args: args{digit: 5},
			want: 9592,
		},
		{
			desc: "78498 / 10^6",
			args: args{digit: 6},
			want: 78498,
		},
		{
			desc: "664579 / 10^7",
			args: args{digit: 7},
			want: 664579,
		},
		/*
			{
				desc: "5761445 / 10^8",
				args: args{digit: 8},
				want: 5761445,
			},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			max := new(big.Int).Exp(ten, big.NewInt(tt.args.digit), nil)
			n := big.NewInt(1)
			got := int64(0)
			for n.Cmp(max) == -1 {
				if n.ProbablyPrime(Reps) {
					got++
				}
				n.Add(n, one)
			}
			if got != tt.want {
				t.Errorf("Prime num = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAll5(b *testing.B) {
	e := big.NewInt(5)
	max := new(big.Int).Exp(ten, e, nil)
	n := big.NewInt(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n.Cmp(max) == -1 {
			n.ProbablyPrime(Reps)
			n.Add(n, one)
		}
	}
}

func BenchmarkAll6(b *testing.B) {
	e := big.NewInt(6)
	max := new(big.Int).Exp(ten, e, nil)
	n := big.NewInt(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n.Cmp(max) == -1 {
			n.ProbablyPrime(Reps)
			n.Add(n, one)
		}
	}
}

func BenchmarkAll7(b *testing.B) {
	e := big.NewInt(7)
	max := new(big.Int).Exp(ten, e, nil)
	n := big.NewInt(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n.Cmp(max) == -1 {
			n.ProbablyPrime(Reps)
			n.Add(n, one)
		}
	}
}

func BenchmarkPrime16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime16.ProbablyPrime(Reps)
	}
}

func BenchmarkPrime39(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime39.ProbablyPrime(Reps)
	}
}
func BenchmarkPrime157(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime157.ProbablyPrime(Reps)
	}
}

func BenchmarkPrime969(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime969.ProbablyPrime(Reps)
	}
}

func BenchmarkPrime6987(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime6987.ProbablyPrime(Reps)
	}
}

func BenchmarkPrime13395(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime13395.ProbablyPrime(Reps)
	}
}

func BenchmarkPrime25962(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime25962.ProbablyPrime(Reps)
	}
}

func BenchmarkPrime39751(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime39751.ProbablyPrime(Reps)
	}
}

func BenchmarkPrime65050(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime65050.ProbablyPrime(Reps)
	}
}

func BenchmarkPrime65087(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prime65087.ProbablyPrime(Reps)
	}
}

func BenchmarkComposite16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		composite16.ProbablyPrime(Reps)
	}
}

func BenchmarkComposite83(b *testing.B) {
	for i := 0; i < b.N; i++ {
		composite83.ProbablyPrime(Reps)
	}
}

func BenchmarkComposite340(b *testing.B) {
	for i := 0; i < b.N; i++ {
		composite340.ProbablyPrime(Reps)
	}
}

func BenchmarkComposite1350(b *testing.B) {
	for i := 0; i < b.N; i++ {
		composite1350.ProbablyPrime(Reps)
	}
}

func BenchmarkComposite20382(b *testing.B) {
	for i := 0; i < b.N; i++ {
		composite20382.ProbablyPrime(Reps)
	}
}

func BenchmarkComposite65713(b *testing.B) {
	for i := 0; i < b.N; i++ {
		composite65713.ProbablyPrime(Reps)
	}
}

func mersenne(m int64) *big.Int {
	e := new(big.Int).Exp(two, big.NewInt(m), nil)
	return new(big.Int).Sub(e, one)
}
