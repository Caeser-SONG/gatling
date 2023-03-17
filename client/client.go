package client


type Client interface {
	Send() (interface{}, error)
}


