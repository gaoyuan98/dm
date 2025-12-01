# 介绍
维护官方dm的go的驱动包，用于日常的项目引入
如果遇到驱动类问题 可提lssues

# 引入方式
```
go get github.com/gaoyuan98/dm
```

# 使用方式
```go
import (
	"database/sql"
	"fmt"
	_ "github.com/gaoyuan98/dm"
)

func main() {
	driverName := "dm"
	dataSourceName := "dm://SYSDBA:SYSDBA@127.0.0.1:5236"
	var db *sql.DB
	var err error
	if db, err = sql.Open(driverName, dataSourceName); err != nil {
		fmt.Println(err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("connect to \"%s\" succeed.\n", dataSourceName)

}

```
# 版本说明
v1.3.162版本 ==> 8.1.3.162 pack12 ENT版本

v1.4.48版本 ==> 8.1.4.48 ENT版本

V1.4.147 版本 ==> 8.1.4.147 ENT版本

V1.4.177 版本 ==> 8.1.4.177 ENT版本

V1.5.7 版本 ==> 8.1.5.7 ENT版本