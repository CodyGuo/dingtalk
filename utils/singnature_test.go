package utils

import (
	"testing"
	"time"
)

func BenchmarkComputeSignatureByte(b *testing.B) {
	timestamp := time.Now().UnixNano() / 1e6
	secret := "SEC07dc82521f02e3318d3cf99db8497d94f041c3565b41b8bea5772c93d65e854f"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ComputeSignature(timestamp, secret)
	}

}

func BenchmarkComputeSignatureFmt(b *testing.B) {
	timestamp := time.Now().UnixNano() / 1e6
	secret := "SEC07dc82521f02e3318d3cf99db8497d94f041c3565b41b8bea5772c93d65e854f"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ComputeSignatureFmt(timestamp, secret)
	}
}
