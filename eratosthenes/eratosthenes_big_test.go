package eratosthenes

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGenarateBig(t *testing.T) {
	type args struct {
		max int64
	}
	tests := []struct {
		desc string
		args args
		want []*big.Int
	}{
		{
			desc: "Case 10",
			args: args{max: int64(10)},
			want: int64ArrToBigIntarr([]int64{2, 3, 5, 7}),
		},
		{
			desc: "Case 100",
			args: args{max: int64(100)},
			want: int64ArrToBigIntarr([]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}),
		},
		{
			desc: "Case 1000",
			args: args{max: int64(1000)},
			want: int64ArrToBigIntarr([]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := GenarateBig(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenarateBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func int64ArrToBigIntarr(arr []int64) []*big.Int {
	r := make([]*big.Int, len(arr))
	for i, a := range arr {
		r[i] = big.NewInt(a)
	}

	return r
}

func BenchmarkGenerateBig6(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenarateBig(1000000)
	}
}

func BenchmarkGenerateBig7(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenarateBig(10000000)
	}
}

func BenchmarkGenerateBig8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenarateBig(100000000)
	}
}