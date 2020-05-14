package wavmatch

// #cgo LDFLAGS: -L ./lib -lmkl_intel_lp64 -lmkl_core -lmkl_sequential -lmkl_vml_avx2 -lmkl_avx2 -lyispeech -lwavmatch -lstdc++
// #cgo CFLAGS: -I ./
//#include <lib/wavmatch.h>
//#include <stdlib.h>
import "C"
import (
	"fmt"
	"os"
	"strings"
	wav "test/cgo/wavmatch/vendor"
	"unsafe"
)

const (
	CONNECTED           string = "100" //已接通
	OFFLINE             string = "1"   //关机
	IN_CONNECTED        string = "2"   //通话中
	OUT_OF_SERVICE      string = "3"   //不在服务区
	ARREARS             string = "4"   //欠费
	TIMEOUT             string = "5"   //无人接听
	TEMP_OUT_OF_CONNECT string = "6"   //暂时无法接通
	HALTED              string = "7"   //停机
	EMPTY_NUMBER        string = "8"   //空号
	BUSY                string = "9"   //电路正忙
	CUSTOMER_BUSY       string = "10"  //用户正忙
	OTHER               string = "11"  //其他原因
	BEYOND_PRIVILEGE    string = "12"  //请勿越权使用
	SYSTEM_BUSY         string = "17"  //系统正忙
	SERVICE_STOPPED     string = "18"  //暂停服务
)

type WavMatch interface {
	ReloadTemplate(path string) error //path: Template path
	Match(voicd string) (string, error)
}

type ZyCosWavMatch struct {
}

func (ZyCosWavMatch) ReloadTemplate(path string) error { //path: Template path
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	ret := int(C.ReloadTemplate(cstr))
	if ret != 0 {
		return fmt.Errorf("ReloadTemplate failed, ret:%d", ret)
	}

	return nil
}

func (ZyCosWavMatch) Match(sid, voice string) (string, error) {
	var cstrMatchedPath *C.char

	save_path := "./rings/" + sid + ".wav"

	if !isExist("./rings/") {
		err := os.MkdirAll("./rings/", os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("Match: %s", err)
		}
	}

	if err := saveVoicetoWav(save_path, voice); err != nil {
		return "", fmt.Errorf("Match:%s", err)
	}

	defer os.Remove(save_path)

	cstrSavePath := C.CString(save_path)
	defer C.free(unsafe.Pointer(cstrSavePath))

	ret := C.WavMatchCos(cstrSavePath, &cstrMatchedPath)
	if ret != 0 {
		return "", fmt.Errorf("Match failed, sid:%s, ret:%d", sid, ret)
	}

	defer C.free(unsafe.Pointer(cstrMatchedPath))

	return MatchEndReason(C.GoString(cstrMatchedPath)), nil
}

func MatchEndReason(fullpath string) string {
	dirs := strings.Split(fullpath, "/")

	path := dirs[len(dirs)-1]

	if 0 == strings.Index(path, "bnjthygj") { //不能接听或已关机
		return OFFLINE
	} else if 0 == strings.Index(path, "bxz") { //被限制
		return BEYOND_PRIVILEGE
	} else if 0 == strings.Index(path, "gj") { //关机
		return OFFLINE
	} else if 0 == strings.Index(path, "gq") { //过期
		return SERVICE_STOPPED
	} else if 0 == strings.Index(path, "hjsx") { //呼叫受限
		return BEYOND_PRIVILEGE
	} else if 0 == strings.Index(path, "hjzy") { //呼叫转移
		return CUSTOMER_BUSY
	} else if 0 == strings.Index(path, "kh") { //空号
		return EMPTY_NUMBER
	} else if 0 == strings.Index(path, "ldtx") { //来电提醒
		return CUSTOMER_BUSY
	} else if 0 == strings.Index(path, "thz") { //通话中
		return IN_CONNECTED
	} else if 0 == strings.Index(path, "tj") { //停机
		return HALTED
	} else if 0 == strings.Index(path, "wfjt") { //暂时无法接通
		return TEMP_OUT_OF_CONNECT
	} else if 0 == strings.Index(path, "wrjt") { //无人接听
		return TIMEOUT
	} else if 0 == strings.Index(path, "xtzm") { //系统正忙
		return SYSTEM_BUSY
	} else if 0 == strings.Index(path, "yhzm") { //用户正忙
		return CUSTOMER_BUSY
	} else if 0 == strings.Index(path, "ztfw") { //暂停服务
		return SERVICE_STOPPED
	} else {
		return OTHER
	}
}

func saveVoicetoWav(path, voice string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("saveVoicetoWav: open file failed, %s", err)
	}
	defer file.Close()

	wavHead := wav.NewWavHead()
	wavHead.SetDataLength(uint32(len(voice)))
	wavHead.SetChannel(1)
	wavHead.SetSampleRate(8000)
	wavHead.SetBitPerSample(16)

	if _, err = file.Write(wavHead.ToByte()); err != nil {
		return fmt.Errorf("saveVoicetoWav: write file failed, %s", err)
	}

	if _, err = file.Write([]byte(voice)); err != nil {
		return fmt.Errorf("saveVoicetoWav: write file failed, %s", err)
	}

	return nil
}
func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
