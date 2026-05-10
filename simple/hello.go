package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello SayHello
}

type SayHelloImpln struct{}

func (hello *SayHelloImpln) Hello(name string) string {

	return "Hello" + name
}

func NewSayHelloImpln() *SayHelloImpln {
	return &SayHelloImpln{}
}
func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{
		SayHello: sayHello,
	}
}
