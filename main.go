package main
import (
	"io/ioutil"
	gs "github.com/huin/goserial"
	"strings"
	"time"
	"encoding/binary"
	"io"
	"bytes"
	"github.com/go-vgo/robotgo"
)

func main(){
	c := &gs.Config{Name: findArduino(), Baud: 9600}
	s, _ := gs.OpenPort(c)
	time.Sleep(1 * time.Second)
	time.Sleep(2 * time.Second)
	sendArduinoCommand('a', 1.0, s)
	time.Sleep(5 * time.Second)
	sendArduinoCommand('a', 3.0, s)
	time.Sleep(5 * time.Second)
	sendArduinoCommand('a', 5.0, s)
}


func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")
	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbserial") ||
		strings.Contains(f.Name(), "ttyUSB") ||
		strings.Contains(f.Name(), "ttyACM") /*for ubuntu*/{
			return "/dev/" + f.Name()
		}
	}
	return ""
}

func sendArduinoCommand(
command byte, arg float32, serialPort io.ReadWriteCloser) error {
	if serialPort == nil {
		return nil
	}
	bufOut := new(bytes.Buffer)
	if err := binary.Write(bufOut, binary.LittleEndian, arg); err != nil {
		return nil
	}
	for _, v := range [][]byte{[]byte{command}, bufOut.Bytes()} {
		if _, err := serialPort.Write(v); err != nil {
			return err
		}
	}
	return nil
}
