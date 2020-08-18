module go-web

go 1.14

require (
	ami v0.0.0
	github.com/go-sql-driver/mysql v1.5.0
)

// 本地库替换
replace ami => ./ami
