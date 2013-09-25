
func (conn *NodeZMQ) Send(msg message.Message) {
  data, err := json.Marshal(msg)
  if err != nil {
    glog.Errorf("marshal error: %s", err)
  } else {
    multipart := [][]byte{
      []byte(b.ID),
      []byte(b.roomID),
      data,
    }
    //log.Printf("sending %d bytes", len(data))
    b.zmqConn.Send <- multipart
  }
}

func (b *Bot) Receive(multipart [][]byte) *event.Event {
  //topic := multipart[0]
  data := multipart[1]
  //log.Printf("received %d bytes for %s", len(data), topic)

  event := &event.Event{}
  if err := event.Unproto(data); err != nil {
    log.Printf("unmarshal error: %s", err)
    return nil
  }

  return event
}
