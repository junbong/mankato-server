package amqpexphandler

import (
	"github.com/Junbong/mankato-server/db/documents"
	"github.com/streadway/amqp"
)

type ExpirationHandler struct {
	URL        string
	Queue      string
	connection *amqp.Connection
}


func New(url, queue string) (*ExpirationHandler) {
	return &ExpirationHandler{
		URL: url,
		Queue: queue,
	}
}


func (h *ExpirationHandler) Open() (error) {
	if connection, err := amqp.Dial(h.URL); err == nil {
		h.connection = connection
		return nil
	} else {
		return err
	}
}


func (h *ExpirationHandler) Close() (error) {
	return h.connection.Close()
}


func (h *ExpirationHandler) HandleDocument(doc *document.Document) (error) {
	if ch, err := h.connection.Channel(); err == nil {
		perr := ch.Publish(
			"", h.Queue, false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body: []byte(doc.Key),
			},
		)
		
		if perr == nil {
			return nil
		} else {
			return err
		}
		
	} else {
		defer ch.Close()
		return err
	}
}


func (h *ExpirationHandler) HandleDocuments(docs ...*document.Document) (unhandled []interface{}, err error) {
	return unhandled, nil
}
