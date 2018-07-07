package p2p

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	log "github.com/cihub/seelog"
	"io"
	"net"
	"strings"
)

type PeerInfo struct {
	Id      string
	Addr    string
	Port    string
	ChainId string
}

func (a *PeerInfo) Equal(b PeerInfo) bool {
	return (a.Id == b.Id && a.Id != "" && b.Id != "") || (a.Addr == b.Addr && a.Port == b.Port)
}

func (a *PeerInfo) IsIncomplete() bool {
	return a.Id == "" || a.Addr == "" || a.Port == "" || a.ChainId == ""
}

//Bigger 1 a > b;  0 a = b ; -1 a < b
func (a *PeerInfo) Bigger(b PeerInfo) int {
	return strings.Compare(a.Id, b.Id)
}

type PeerData struct {
	Id    string
	Index uint16
}

type PeerDataSet []PeerData

func (s PeerDataSet) Len() int {
	return len(s)
}

func (s PeerDataSet) Less(i, j int) bool {
	return s[i].Id > s[j].Id
}

func (s PeerDataSet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Peer struct {
	Info  PeerInfo
	Index uint16

	/*peer state*/
	State int

	conn   net.Conn
	isconn bool
	reader *bufio.Reader

	In bool

	sendup SendupCb
}

func CreatePeer(info PeerInfo, conn net.Conn, in bool, sendup SendupCb) *Peer {
	return &Peer{
		Info:   info,
		State:  PEER_STATE_INIT,
		conn:   conn,
		isconn: true,
		reader: bufio.NewReader(conn),
		In:     in,
		sendup: sendup,
	}
}

func (p *Peer) Start() {
	go p.recvRoutine()
}

func (p *Peer) Stop() {
	p.conn.Close()
}

func (p *Peer) Send(packet Packet) error {
	var length uint32
	var head Head
	headsize := uint32(binary.Size(head))

	if packet.Data == nil {
		length = headsize
	} else {
		length = headsize + uint32(len(packet.Data))
	}

	if length > MAX_PACKET_LEN {
		log.Errorf("p2p Send packet length large than max packet length")
		return errors.New("large than max packet length")
	}

	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, length)
	if err != nil {
		log.Error("p2p send write packet length error")
		return err
	}

	err = binary.Write(buf, binary.BigEndian, packet.H)
	if err != nil {
		log.Error("p2p send write packet protocolType error")
		return err
	}

	_, err = buf.Write(packet.Data)
	if err != nil {
		log.Error("p2p send write packet Data error")
		return err
	}

	_, err = p.conn.Write(buf.Bytes())
	return err
}

func (p *Peer) recvRoutine() {
	defer p.conn.Close()

	bl := make([]byte, 4)
	var packetLen uint32
	var len int
	var head Head
	readerr := false
	headsize := uint32(binary.Size(head))

	for {
		_, err := io.ReadFull(p.reader, bl)
		if err != nil {
			log.Errorf("p2p recvRoutine read head error:%s", err)
			p.isconn = false
			return
		}

		packetLen = binary.BigEndian.Uint32(bl)
		if packetLen < headsize || packetLen > MAX_PACKET_LEN {
			log.Errorf("p2p recvRoutine drop packet wrong packet lenght %d", packetLen)
			continue
		}

		buf := make([]byte, packetLen)
		len, err = io.ReadFull(p.reader, buf)
		if err != nil {
			log.Errorf("p2p recvRoutine read data error:%s", err)
			p.isconn = false
			break
		}

		if uint32(len) < packetLen {
			for {
				length, err := io.ReadFull(p.reader, buf[len:])
				if err != nil {
					log.Errorf("p2p recvRoutine continue read data error:%s", err)
					p.isconn = false
					return
				}

				len += length

				if uint32(len) < packetLen {
					continue
				} else if uint32(len) == packetLen {
					break
				} else {
					log.Errorf("p2p recvRoutine continue read data length wrong packet length:%d, read:%d", packetLen, len)
					readerr = true
					break
				}
			}
		}

		if readerr {
			readerr = false
			continue
		}

		var packet Packet

		packet.H.ProtocolType = uint16(binary.BigEndian.Uint16(buf))
		packet.H.PacketType = uint16(binary.BigEndian.Uint16(buf[2:]))

		if packetLen > headsize {
			packet.Data = buf[headsize:packetLen]
		}

		p.sendup(p.Index, &packet)
	}
}
