package proto

import (
	"bytes"
	"encoding/binary"
	"ggquotation/protobuf"
	l4g "github.com/alecthomas/log4go"
	"io"
	"net"
	"strings"
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

		currBuffer = append(currBuffer, buffer[:n]...)
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
	// TODO: use checkSum to validate whether PACKAGE_FULL
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

func ParseMessage3xxx11(pkg []byte) {
	quotation := &protobuf.Quotation300111{}
	quotation.Head = &protobuf.Head3Xxx11{}

	binary.Read(bytes.NewBuffer(pkg[0:4]), binary.BigEndian, &quotation.Head.MsgType)
	binary.Read(bytes.NewBuffer(pkg[8:16]), binary.BigEndian, &quotation.Head.OrigTime)
	binary.Read(bytes.NewBuffer(pkg[16:18]), binary.BigEndian, &quotation.Head.ChannelNo)
	quotation.Head.MDStreamID = strings.TrimSpace(string(pkg[18:21]))
	quotation.Head.SecurityID = strings.TrimSpace(string(pkg[21:29]))
	quotation.Head.SecurityIDSource = strings.TrimSpace(string(pkg[29:33]))

	quotation.Head.TradingPhaseCode = make([]byte, 8)
	copy(quotation.Head.TradingPhaseCode, pkg[33:41])

	binary.Read(bytes.NewBuffer(pkg[41:49]), binary.BigEndian, &quotation.Head.PrevClosePx)
	binary.Read(bytes.NewBuffer(pkg[49:57]), binary.BigEndian, &quotation.Head.NumTrades)
	binary.Read(bytes.NewBuffer(pkg[57:65]), binary.BigEndian, &quotation.Head.TotalVolumeTrade)
	binary.Read(bytes.NewBuffer(pkg[65:73]), binary.BigEndian, &quotation.Head.TotalValueTrade)
}

func ParseMessage300111(quotation protobuf.Quotation300111, pkg []byte, index int) {
	binary.Read(bytes.NewBuffer(pkg[index:index+4]), binary.BigEndian, &quotation.NoMDEntries)
	for i := 0; i < int(quotation.NoMDEntries); i++ {
		mdEntry := &protobuf.MDEntry300111{}
		mdEntry.MDEntryType = strings.TrimSpace(string(pkg[index+4 : index+6]))
		binary.Read(bytes.NewBuffer(pkg[index+6:index+14]), binary.BigEndian, mdEntry.MDEntryPx)
		binary.Read(bytes.NewBuffer(pkg[index+6:index+14]), binary.BigEndian, mdEntry.MDEntrySize)
		binary.Read(bytes.NewBuffer(pkg[index+14:index+22]), binary.BigEndian, mdEntry.MDEntryPx)
		binary.Read(bytes.NewBuffer(pkg[index+22:index+24]), binary.BigEndian, mdEntry.MDPriceLevel)
		binary.Read(bytes.NewBuffer(pkg[index+24:index+32]), binary.BigEndian, mdEntry.NumberOfOrders)
		binary.Read(bytes.NewBuffer(pkg[index+32:index+36]), binary.BigEndian, mdEntry.NoOrders)

		quotation.MDEntry = append(quotation.MDEntry, mdEntry)
	}
}
