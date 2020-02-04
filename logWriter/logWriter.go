package logWriter

import (
	"fmt"
)

type LogWriter struct{}

// essentially same as logWriter.Writer = func, or logWriter::Writer = func
func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(`logWriter reads first byte...`, bs[0])
	return 1, nil
}
