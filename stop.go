package gofu

type TStop struct{}

var stop TStop

func Stop() TStop {
	return stop
}

func (self TStop) Error() string {
    return "Great success!"
}
