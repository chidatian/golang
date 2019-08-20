package client


func (this *Client) SendData(info []byte) {
	// Write(b []byte) (n int, err error)
	number := this.Msg.Uint32ToByte(info)
	_, err := this.Conn.Write(number)
	if err != nil {
		panic(err)
	}

	_, err = this.Conn.Write(info)
	if err != nil {
		panic(err)
	}
}

func (this *Client) ReadData() string{
	resp := make([]byte, 10240)
	number := make([]byte, 5)
	// Read(b []byte) (n int, err error)
	_, err := this.Conn.Read(number[:4])
	right := this.Msg.ToUint32(number[:4])
	if err != nil {
		panic(err)
	}

	n, err := this.Conn.Read(resp[:right])
	if err != nil {
		panic(err)
	}
	return string(resp[:n])
}