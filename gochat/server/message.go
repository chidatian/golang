package server


// 向用户 发送数据
func (this *Server) SendData(index uint16, info []byte) {
	conn := this.LinkPool.Get(index)
	number := this.Msg.Uint32ToByte(info)
	_, err := conn.Write(number)
	if err != nil {
		panic(err)
	}

	_, err = conn.Write(info)
	if err != nil {
		panic(err)
	}
}

// 接收用户的数据
func (this *Server) ReadData(index uint16) (string, error){
	conn := this.LinkPool.Get(index)
	resp := make([]byte, 10240)
	number := make([]byte, 5)
	// Read(b []byte) (n int, err error)
	_, err := conn.Read(number[:4])
	if err != nil {
		return "", err
	}

	right := this.Msg.ToUint32(number[:4])
	n, err := conn.Read(resp[:right])
	if err != nil {
		return "", err
	}
	return string(resp[:n]), nil
}