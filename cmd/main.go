package main

import (
	"github.com/dli-invest/finreddit/pkg/reddit"
)

func main() {

	// Ready to make API calls!
	reddit.ScanSRs("cmd/scan_sr/simple.yml")
}