package pac

import (
	"github.com/GregorioMonari/testfix-sepago/sepa"
	"github.com/GregorioMonari/testfix-sepago/sepa/sparql"
)

type consumer struct {
	query string
	app Application
}

func (c consumer)Consume(handler func(notification *sparql.Notification),data interface{})(sepa.Subscription,error)  {
	ql,err := c.app.profile.GetQuery(c.query,data)
	if err != nil {
		return sepa.Subscription{}, err
	}
	return c.app.client.Subscribe(ql,handler)
}
