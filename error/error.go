package tferror

var (
	// Errors
	err error
)

func Panic(err error) {

	if err != nil {
		panic(err)
	}
}
