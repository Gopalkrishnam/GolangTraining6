package subscribers

import "errors"

type Subscriber struct {
	SId  int
	Name string
}

var (
	subscribers = make([]Subscriber, 3, 3)
)

func init() {
	subscribers[0] = Subscriber{SId: 1, Name: "Anand"}
	subscribers[1] = Subscriber{SId: 2, Name: "Vijay"}
	subscribers[2] = Subscriber{SId: 3, Name: "Kumar"}
}

func GetAll() []Subscriber {
	subs := subscribers
	return subs
}

func Get(sid int) (s Subscriber, e error) {
	for _, v := range subscribers {
		if sid == v.SId {
			s = v
			return
		}
	}
	e = errors.New("subscriber not found")
	return
}
