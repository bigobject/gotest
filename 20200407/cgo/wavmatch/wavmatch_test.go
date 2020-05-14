package wavmatch_test

import (
	"fmt"
	"os"
	"sync"
	"test/cgo/wavmatch"
	"testing"
)

func TestZyCosWavMatch_ReloadTemplate(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		z       wavmatch.ZyCosWavMatch
		args    args
		wantErr bool
	}{
		{name: "invalid path", z: wavmatch.ZyCosWavMatch{}, args: args{path: "./sdsdsds/"}, wantErr: true},
		{name: "valid path", z: wavmatch.ZyCosWavMatch{}, args: args{path: "../wav/"}, wantErr: false},
	}
	for _, tt := range tests {
		if err := tt.z.ReloadTemplate(tt.args.path); (err != nil) != tt.wantErr {
			t.Errorf("%q. ZyCosWavMatch.ReloadTemplate() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestZyCosWavMatch_Match(t *testing.T) {
	type args struct {
		sid   string
		voice string
	}
	tests := []struct {
		name    string
		z       wavmatch.ZyCosWavMatch
		args    args
		want    string
		wantErr bool
	}{
		{name: "wrjt", z: wavmatch.ZyCosWavMatch{}, args: args{sid: "wss", voice: getVoice("../rings/yhzm_5s.wav")}, want: wavmatch.MatchEndReason("../rings/yhzm_5s.wav"), wantErr: false},
	}
	for _, tt := range tests {
		tt.z.ReloadTemplate("../wav/")

		got, err := tt.z.Match(tt.args.sid, tt.args.voice)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. ZyCosWavMatch.Match() error = %v, want= %v, wantErr %v", tt.name, err, tt.want, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. ZyCosWavMatch.Match() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkZyCosWavMatch_Match(b *testing.B) {
	type args struct {
		sid   string
		voice string
	}
	tests := []struct {
		name    string
		z       wavmatch.ZyCosWavMatch
		args    args
		want    string
		wantErr bool
	}{
		{name: "wrjt", z: wavmatch.ZyCosWavMatch{}, args: args{sid: "wrjt", voice: getVoice("../rings/yhzm_5s.wav")}, want: wavmatch.MatchEndReason("../rings/yhzm_5s.wav"), wantErr: false},
	}

	pass := make(chan bool, 4)
	for i := 0; i < 4; i++ {
		pass <- true
	}

	z := wavmatch.ZyCosWavMatch{}
	z.ReloadTemplate("../wav/")

	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			go func() {
				defer wg.Done()
				<-pass
				got, err := tt.z.Match(fmt.Sprintf("%s_%d", tt.args.sid, i), tt.args.voice)
				if (err != nil) != tt.wantErr {
					b.Errorf("%q. ZyCosWavMatch.Match() error = %v, wantErr %v", tt.name, err, tt.wantErr)
					pass <- true
					return
				}
				if got != tt.want {
					b.Errorf("%q. ZyCosWavMatch.Match() = %v, want %v", tt.name, got, tt.want)
				}
				pass <- true
			}()
		}
	}
	wg.Wait()
}

func getVoice(path string) string {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return ""
	}
	defer file.Close()

	fileInfo, err := os.Stat(path)
	if err != nil {
		return ""
	}

	if fileInfo.Size() < 44 {
		return ""
	}

	Voice := make([]byte, fileInfo.Size())
	if n, err := file.Read(Voice); err != nil || int64(n) != fileInfo.Size() {
		return ""
	}

	return string(Voice[44:])
}
