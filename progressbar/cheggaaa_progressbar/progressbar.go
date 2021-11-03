package main

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	count := 100
	bar := pb.StartNew(count)
	// start bar from 'default' template
	// bar := pb.Default.Start(count)

	// start bar from 'simple' template
	// bar := pb.Simple.Start(count)

	// start bar from 'full' template
	// bar := pb.Full.Start(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond * 40)
	}
	bar.Finish()

	var limit int64 = 1024 * 1024 * 50000
	// we will copy 500 MiB from /dev/rand to /dev/null
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard

	// start new bar
	bar2 := pb.Full.Start64(limit)

	// create proxy reader
	barReader := bar2.NewProxyReader(reader)

	// copy from proxy reader
	io.Copy(writer, barReader)

	// finish bar
	bar2.Finish()
}
