package tferror

var (
	// Errors
	err error
)

// Panic : Show fatal errors
func Panic(err error) {

	if err != nil {
		panic(err)
	}
}
