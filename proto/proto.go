package proto

import (
	"bytes"
	"encoding/binary"
	l4g "github.com/alecthomas/log4go"
	"io"
	"net"
)

const (
	PACKAGE_LESS = iota
	PACKAGE_FULL
	PACKAGE_ERROR
)

func NewLoginBody() []byte {
	compid := [64]byte{0: 0x20}

	buff := bytes.NewBuffer([]byte{})

	binary.Write(buff, binary.BigEndian, compid)

	return buff.Bytes()
}

func Recv(conn net.Conn) {
	buffer := make([]byte, 1024*4)
	var currBuffer []byte
	var n int
	var err error
	for {
		n, err = conn.Read(buffer)
		if err != nil {
			netErr, ok := err.(net.Error)
			if ok && netErr.Timeout() && netErr.Temporary() {
				continue //no data,not error
			}

			if _, ok := err.(*net.OpError); ok {
				l4g.Error("Net OpError", conn.RemoteAddr())
				conn.Close()
				return
			}

			if err == io.EOF {
				l4g.Error("Connection closed by remote", conn.RemoteAddr())
			} else {
				l4g.Error("Read buffer error", err)
			}

			conn.Close()
			return
		}

		currBuffer := append(currBuffer, buffer[:n]...)

		for {
			pkgLen, status := ParsePackage(currBuffer)
			if status == PACKAGE_LESS {
				break
			}

			if status == PACKAGE_FULL {
				pkg := make([]byte, pkgLen)
				copy(pkg, currBuffer[:pkgLen])
				currBuffer := currBuffer[pkgLen:]
				//TODO:go routine parse message
				if len(currBuffer) > 0 {
					continue
				}
				currBuffer = nil
				break
			}
		}

		l4g.Error("Parse package error")
		conn.Close()
		return
	}
}

func ParsePackage(buff []byte) (pkgLen, status int) {

	length := len(buff)
	if length < 8 {
		return 0, PACKAGE_LESS
	}

	bodylength := binary.BigEndian.Uint32(buff[4:8])

	//TODO:
	/*
		if bodylength >104857600 || len(buff) > 104857600 {
			return 0,PACKAGE.ERROR
		}
	*/
	if len(buff) < int(8+bodylength+4) {
		return 0, PACKAGE_LESS
	}

	return int(8 + bodylength + 4), PACKAGE_FULL
}

func ParseMessage(pkg []byte) {
	msgType := binary.BigEndian.Uint32(pkg[:4])

	switch msgType {
	case 1: // Logon
	case 2: // Logout
	case 3: // Heartbeat
	case 390095: //行情数据频道心跳
	case 390094: //重传消息(逐笔行情数据或公告消息重传)
	case 390090: //快照行情频道统计消息
	case 8: //业务拒绝消息
	case 390019: //市场实时状态
	case 390013: //证券实时状态
	case 390012: //公告消息
	case 300111: //集中竞价交易业务行情
	case 300611: //盘后定价交易业务行情
	case 309011: //指数行情
	case 309111: //成交量统计指标行情
	case 306311: //港股实时行情
	case 300192: //集中竞价业务逐笔委托行情
	case 300592: //协议交易业务逐笔委托行情
	case 300792: //转融通证券出借业务逐笔委托行情
	case 300191: //集中竞价业务逐笔成交行情
	case 300591: //协议交易业务逐笔成交行情
	case 300791: //转融通证券出借业务逐笔成交行情
	}
}
