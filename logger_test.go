package dotlog

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestLogger_Info(t *testing.T) {
	err := StartLogService("d:/gotmp/golog/log.conf")
	fmt.Println(err)

	log1 := GetLogger("log1")
	log1.Info("Test_Info")
	log1.Error(errors.New("test error"), "Test_Error")

	time.Sleep(time.Second * 10)
}

func TestLogger_InfoS(t *testing.T) {
	err := StartLogService("d:/gotmp/golog/log.conf")
	fmt.Println(err)

	log1 := GetLogger("log1")
	log1.InfoS("Test_Info", 1, true, "test2")
	time.Sleep(time.Second * 10)
}

func TestLogger_InfoF(t *testing.T) {

}

func BenchmarkTest_Info(b *testing.B) {
	err := StartLogService("d:/gotmp/golog/log.conf")
	fmt.Println(err)

	log1 := GetLogger("log1")

	for i := 0; i < b.N; i++ {
		log1.Info("Test_Info")
	}

	time.Sleep(time.Second * 10)
}
