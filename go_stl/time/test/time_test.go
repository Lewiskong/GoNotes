package test

import (
	"testing"
	time "../src"
	"fmt"
)

func Test_Ticker(t *testing.T){
	ticker:=time.NewTicker(time.Second*5)
	for tm :=range ticker.C {
		fmt.Println(tm.String())
	}


}

func Test_map(){

}

