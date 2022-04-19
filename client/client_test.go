package client

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	header := map[string][]string{
		"Content-Length": {"8"},
		"Cookie":         {"qqqq=123OB;BAIDUID=D66EB4351D00F3F540C5B197F81F76D2:FG=1; BIDUPSID=D66EB4351D00F3F57C7E32DB09FCFB07; H_PS_PSSID=36174_36019_36167_34584_36120_36073_36125_35864_36233_26350_36094_36061; PSTM=1650254443; BDSVRTM=21; BD_HOME=1"},
	}
	h := NewHttpClient("GET", "http://www.baidu.com", nil, header)
	e := h.Send()
	fmt.Println(e)
}
