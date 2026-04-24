module modernc.org/cc/v4

go 1.25

require (
	github.com/dustin/go-humanize v1.0.1
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2
	modernc.org/ccorpus2 v1.6.0
	modernc.org/mathutil v1.7.1
	modernc.org/opt v0.2.0
	modernc.org/sortutil v1.2.1
	modernc.org/strutil v1.2.1
	modernc.org/token v1.1.0
)

require github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect

retract v4.27.2 // The fix belongs elsewhere
