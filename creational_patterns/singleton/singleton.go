package singleton

// Singleton contract
type Singleton interface {
	Increment() int
}

// Singleton struct
type singleton struct {
	counter int
}

// implementation of singleton contract
func (s *singleton) Increment() int {
	s.counter++

	return s.counter
}

// making singleton instance
var instance *singleton

func NewInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}
