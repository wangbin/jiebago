#!/bin/sh

go test -run=XXX -v -bench=BenchmarkViterbi -benchtime 10s  -benchmem -memprofile viterbi-mem.out -cpuprofile viterbi-cpu.out
go tool pprof -png -output ~/tmp/viterbi-cpu.png posseg.test viterbi-cpu.out
go tool pprof -png -output ~/tmp/viterbi-mem.png posseg.test viterbi-mem.out
