module GoStudy

go 1.15

require (
	github.com/cavaliercoder/go-rpm v0.0.0-20200122174316-8cb9fd9c31a8
	github.com/isbm/go-deb v0.0.0-20200606113352-45f79b074aa5
	github.com/mattn/go-sqlite3 v1.14.8
	github.com/mitchellh/go-ps v1.0.0
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
)

replace github.com/cavaliercoder/go-rpm => ./go-rpm

replace github.com/isbm/go-deb => ./go-deb
