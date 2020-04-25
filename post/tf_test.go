package badgo

// import (
// 	tf "github.com/tensorflow/tensorflow/tensorflow/go"
// 	"testing"
// )

// func BenchmarkNewTensor(b *testing.B) {
// 	vec := make([][][]float32, 100)
// 	for i := range vec {
// 		vec[i] = make([][]float32, 100)
// 		for j := range vec[i] {
// 			vec[i][j] = make([]float32, 100)
// 		}
// 	}
// 	b.ReportAllocs()
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			tf.NewTensor(vec)
// 		}
// 	})
// }
