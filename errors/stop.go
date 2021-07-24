package errors

type TStop struct{}

var Stop TStop

func (self TStop) Error() string {
    return "Great success!"
}
