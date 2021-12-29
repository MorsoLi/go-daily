## database/sql 接口
* Go 官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发相应的数据库驱动
  
```go

// https://github.com/mattn/go-sqlite3 驱动
func init() {
    sql.Register("sqlite3", &SQLiteDriver{})
}

// https://github.com/mikespook/mymysql 驱动
// Driver automatically registered in database/sql
var d = Driver{proto: "tcp", raddr: "127.0.0.1:3306"}
func init() {
    Register("SET NAMES utf8")
    sql.Register("mymysql", &d)
}
```

* 通过调用Register函数来注册自己的数据库驱动名称以及相应的 driver 实现，在 database/sql 内部通过一个 map 来存储用户定义的相应驱动。
   
```go

var drivers = make(map[string]driver.Driver)

drivers[name] = driver
```
* 在 database/sql 内部通过一个 map 来存储用户定义的相应驱动，可以同时注册多个不重复的数据库驱动

1. dirver.Driver 
   - 定义了一个 method： Open (name string)，这个方法返回一个数据库的 Conn 接口;第三方驱动都会定义这个函数，它会解析 name 参数来获取相关数据库的连接信息，解析完成后，它将使用此信息来初始化一个 Conn 并返回它
```go
type Driver interface {
    Open(name string) (Conn, error)
}
```

1. driver.Conn
   - Conn 是一个数据库连接的接口定义, 返回的 Conn 只能用来进行一次 goroutine 的操作, 它定义了一系列method: Prepare 函数返回与当前连接相关的执行 Sql 语句的准备状态，可以进行查询、删除等操作。Close 函数关闭当前的连接，执行释放连接拥有的资源等清理工作。因为驱动实现了 database/sql 里面建议的 conn pool，所以不用再去实现缓存 conn 之类的，这样会容易引起问题。Begin 函数返回一个代表事务处理的 Tx，通过它你可以进行查询，更新等操作，或者对事务进行回滚、递交。

```go

type Conn interface {
    Prepare(query string) (Stmt, error)
    Close() error
    Begin() (Tx, error)
}
```

3. driver.Stmt
   - 说明
```go

type Stmt struct{
    Close() error
    finalClose() error
    QueryRow(args ...interface{}) *Row
    Exec(args ...interface{}) (Result, error)
    ExecContext(ctx context.Context, args ...interface{}) (Result, error)
    ...
}
```

1. driver.Tx
   - 事务处理一般就两个过程，递交或者回滚
```go

type Tx struct {
    Commit() error
    Rollback() error
    ...
}
```

5. driver.Execer
   - Conn 可选择实现的接口，如果这个接口没有定义，那么在调用 DB.Exec, 就会首先调用 Prepare 返回 Stmt，然后执行 Stmt 的 Exec，然后关闭 Stmt。
```go
type Execer interface {
	Exec(query string, args []Value) (Result, error)
}
```

6. driver.Result
   -- LastInsertId 函数返回由数据库执行插入操作得到的自增 ID 号;RowsAffected 函数返回 query 操作影响的数据条目数。
```go
type Result interface {
    LastInsertId() (int64, error)
    RowsAffected() (int64, error)
}
```

7. driver.Rows
```go

type Rows struct {
    Columns() ([]string, error)
    ColumnTypes() ([]*ColumnType, error)
    Next() bool 
    awaitDone(ctx, txctx context.Context)
    ...
}
```