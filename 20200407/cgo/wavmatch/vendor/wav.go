package wav

import "encoding/binary"

type WavHead struct {
	RIFF         string `json:"RIFF"`       //should be "RIFF", 4 byte
	Length       uint32 `json:"Length"`     //should be length of data + 44 - 8, 4 byte
	WAVE         string `json:"WAVE "`      //should be "WAVE", 4 byte
	fmt          string `json:"fmt"`        //should be "fmt ", 4 byte
	size1        uint32 `json:"size1"`      //should be 16, 4 byte
	formatTag    uint16 `json:"formatTag"`  //should be 1, 2 byte
	channel      uint16 `json:"channel"`    //should be number of tracks, 2 byte
	sampleRate   uint32 `json:"sampleRate"` //should be sampleRate, 4 byte
	bytePerSec   uint32 `json:"bytePerSec"` //should be sampleRate*bitPerSample/8, 4 byte
	blockAlign   uint16 `json:"formatTag"`  //should be number of byte for a sample, bitPerSample/8, 2 byte
	bitPerSample uint16 `json:"formatTag"`  //should be number bit for a sample, 8 or 16, 2 byte
	data         string `json:"data"`       //should be "data", 4 byte
	size2        uint32 `json:"size2"`      //should be length of data, 4 byte
}

func (head WavHead) ToByte() []byte {
	var data []byte

	data = append(data, []byte(head.RIFF)...)

	buf32 := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf32, head.Length)
	data = append(data, buf32...)

	data = append(data, []byte(head.WAVE)...)
	data = append(data, []byte(head.fmt)...)

	binary.LittleEndian.PutUint32(buf32, head.size1)
	data = append(data, buf32...)

	buf16 := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf16, head.formatTag)
	data = append(data, buf16...)

	binary.LittleEndian.PutUint16(buf16, head.channel)
	data = append(data, buf16...)

	binary.LittleEndian.PutUint32(buf32, head.sampleRate)
	data = append(data, buf32...)

	binary.LittleEndian.PutUint32(buf32, head.bytePerSec)
	data = append(data, buf32...)

	binary.LittleEndian.PutUint16(buf16, head.blockAlign)
	data = append(data, buf16...)

	binary.LittleEndian.PutUint16(buf16, head.bitPerSample)
	data = append(data, buf16...)

	data = append(data, []byte(head.data)...)

	binary.LittleEndian.PutUint32(buf32, head.size2)
	data = append(data, buf32...)
	return data
}

func NewWavHead() (head WavHead) {
	head.RIFF = "RIFF"
	//Length wait for set
	head.WAVE = "WAVE"
	head.fmt = "fmt "
	head.size1 = 16
	head.formatTag = uint16(1)
	//channel wait for set
	//sampleRate wait for set
	//bytePerSec wait for set
	//blockAlign wait for set
	//bitPerSample wait for set
	head.data = "data"
	//size2 wait for set

	return head
}

func (head *WavHead) SetDataLength(length uint32) {
	head.size2 = length
	head.Length = length + 44 - 8
}

func (head *WavHead) SetChannel(channel uint32) {
	head.channel = uint16(channel)
}
func (head *WavHead) SetSampleRate(sampleRate uint32) {
	head.sampleRate = sampleRate
	head.bytePerSec = head.sampleRate * uint32(head.bitPerSample) / 8
}

func (head *WavHead) SetBitPerSample(bitPerSample uint32) {
	head.bitPerSample = uint16(bitPerSample)
	head.bytePerSec = head.sampleRate * uint32(head.bitPerSample) / 8
	head.blockAlign = uint16(head.bitPerSample / 8)
}
