package channel

import "testing"

type largeStruct struct {
	str  string
	Data [1000]int
}

func getLargeStruct() largeStruct {
	return largeStruct{
		str: "Larger struct, how would it impact the benchmark?",
	}
}

func BenchmarkChannels_Buffered_Write(b *testing.B) {
	c := make(chan struct{}, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c <- struct{}{}
	}
}

func BenchmarkChannels_Buffered_Read(b *testing.B) {
	c := make(chan struct{}, b.N)
	for i := 0; i < b.N; i++ {
		c <- struct{}{}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-c
	}
}

func BenchmarkChannels_Buffered_Write_Message(b *testing.B) {
	c := make(chan largeStruct, b.N)
	msg := getLargeStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c <- msg
	}
}

func BenchmarkChannels_Buffered_Read_Message(b *testing.B) {
	c := make(chan largeStruct, b.N)
	msg := getLargeStruct()
	for i := 0; i < b.N; i++ {
		c <- msg
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-c
	}
}
