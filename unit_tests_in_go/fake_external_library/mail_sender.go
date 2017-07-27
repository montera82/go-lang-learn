package mail_sender

type Message struct {
	Email, Subject string
	Body           []byte
}

//A message instance has the ability to send a message.
func (m *Message) Send() (string, error) {

	return "Message sent successfully by thirdparty library", nil
}
