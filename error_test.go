package errors_test

import (
	"fmt"
	"testing"

	"github.com/bdlm/errors"
	"github.com/bdlm/log"

	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		//DisableTTY: true,
		//ForceTTY: true,
		FieldMap: log.FieldMap{
			"data": "_",
		},
		//EnableTrace: true,
	})
}

func TestErrors(t *testing.T) {
	assert := assert.New(t)
	var e error

	//e = errors.New("error 1")
	//e = errors.Wrap(e, "error 2")
	//e = errors.Wrap(e, "error 3")

	e = errors.New("error 1")
	e = errors.Track(e)
	e = errors.Track(e)
	e = errors.Wrap(e, "error 2")
	e = errors.Track(e)
	e = errors.Wrap(e, "error 3")
	e = errors.Track(e)
	e = errors.Track(e)

	//spew.Printf("spew:\n%#+v\n\n", e)

	// '#' modeJSON = true
	// ' ' flagFormat = true
	// '-' flagTrace = true
	// '+' flagDetail = true

	//fmt.Printf("%%v:\n%v\n\n", e)
	//fmt.Printf("%%-v:\n%-v\n\n", e)
	//fmt.Printf("%%+v:\n%+v\n\n", e)
	//fmt.Printf("%%+-v:\n%+-v\n\n", e)

	//fmt.Printf("%% v:\n% v\n\n", e)
	//fmt.Printf("%% -v:\n% -v\n\n", e)
	//fmt.Printf("%% +v:\n% +v\n\n", e)
	//fmt.Printf("%% +-v:\n% +-v\n\n", e)

	//fmt.Printf("%%#v:\n%#v\n\n", e)
	//fmt.Printf("%%#-v:\n%#-v\n\n", e)
	//fmt.Printf("%%#+v:\n%#+v\n\n", e)
	//fmt.Printf("%%#+-v:\n%#+-v\n\n", e)

	//fmt.Printf("%% #v:\n% #v\n\n", e)
	fmt.Printf("%% #-v:\n% #-v\n\n", e)
	//fmt.Printf("%% #+v:\n% #+v\n\n", e)
	//fmt.Printf("%% #+-v:\n% #+-v\n\n", e)

	//fmt.Printf("%%v:\n%v\n\n", e)
	//fmt.Printf("%%+v:\n%+v\n\n", e)
	//fmt.Printf("%%#v:\n%#v\n\n", e)
	//fmt.Printf("%%#+v:\n%#+v\n\n", e)
	//fmt.Printf("%% v:\n% v\n\n", e)
	//fmt.Printf("%% +v:\n% +v\n\n", e)
	//fmt.Printf("%% #v:\n% #v\n\n", e)
	//fmt.Printf("%% #+v:\n% #+v\n\n", e)
	//fmt.Printf("%%- v:\n%- v\n\n", e)
	//fmt.Printf("%%- +v:\n%- +v\n\n", e)
	//fmt.Printf("%%- #v:\n%- #v\n\n", e)
	//fmt.Printf("%%- #+v:\n%- #+v\n\n", e)

	//log.WithError(e).Info("log test")

	//byts, _ := json.Marshal(e)
	//assert.Equal(2, 1, string(byts))

	assert.Equal(2, 1, "stop")
}