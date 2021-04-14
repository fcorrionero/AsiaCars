package benchmark

import (
	"fmt"
	"github.com/fcorrionero/europcar/infrastructure/memory"
	"github.com/fcorrionero/europcar/test"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSave(b *testing.B) {
	repo := memory.New()
	for i := 0; i < b.N; i++ {
		v, _ := test.GetVehicleWithParams(
			RandStringChassisNumber(),
			test.ValidLicensePlate,
			test.ValidCategory, time.Now(),
		)
		v.DeviceSerialNumber = fmt.Sprintf("%d", i)
		_ = repo.Save(v)
	}
}

func RandStringChassisNumber() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 17)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
