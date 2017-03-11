package exphandler

import "github.com/Junbong/mankato-server/db/documents"

type ExpirationHandler interface {
	/*
	Prepare this handler if necessary,
	open connection to queue for example.
	 */
	Open() (error)
	
	/*
	Close this handler to make it clear all requests.
	 */
	Close() (error)
	
	HandleDocument(doc *document.Document) (error)
	HandleDocuments(docs ...*document.Document) (unhandled []interface{}, err error)
}
