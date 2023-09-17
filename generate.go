package gogit

//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative gogit.proto
//go:generate stringer -type=Mode -linecomment
