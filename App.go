package main

import (
	"flag"
	"fmt"
	"github.com/bobby96333/goSqlHelper"
	"strings"
)

type Csvdump struct {
	Arg_e      string
	Arg_dsn    string
	Arg_h      bool
	connection *goSqlHelper.SqlHelper
}

func (this *Csvdump) Init() {

	flag.StringVar(&this.Arg_e, "e", "", "executing sql")
	flag.StringVar(&this.Arg_dsn, "dsn", "", "connecting db dsn,example:mysql://root:password@tcp(127.0.0.1:3006)/dbname")
	flag.BoolVar(&this.Arg_h, "h", false, "show command help")
	flag.Parse()

}

func (this *Csvdump) Run() {
	if this.Arg_h {
		flag.Usage()
		return
	}
	this.connection = new(goSqlHelper.SqlHelper)

	index1 := strings.Index(this.Arg_dsn, "://")
	if index1 == -1 {
		panic("dsn format error!")
	}
	driver := this.Arg_dsn[0:index1]
	if driver != "mysql" {
		panic("dsn driver error")
	}
	mysqlDsn := this.Arg_dsn[index1+3:]
	err := this.connection.Open(mysqlDsn)
	this.checkError(err)

	querying, err := this.connection.Querying(this.Arg_e)
	this.checkError(err)
	var i = 0
	var cols []string

	for row, err := querying.QueryRow(); this.checkError(err) && row != nil; row, err = querying.QueryRow() {
		i++
		if i == 1 {
			//first line
			cols, err = querying.Columns()
			if err != nil {
				panic(err)
			}
			this.outputRow(cols...)
		}
		//output data
		line := make([]string, len(cols))
		for ii := 0; ii < len(cols); ii++ {
			val := row.PString(cols[ii])
			if val == "%!V(<nil>)" {
				val = ""
			}
			line[ii] = val
		}
		this.outputRow(line...)
	}
}
func (this *Csvdump) outputRow(args ...string) {
	str := this.formatRow(args...)
	fmt.Println(str)
}
func (this *Csvdump) formatRow(args ...string) string {

	var ret = ""
	for i, arg := range args {
		if i > 0 {
			ret += ","
		}
		ret += "\""
		ret += strings.Replace(arg, "\"", "\"\"", -1)
		ret += "\""
	}

	return ret

}

func (this *Csvdump) checkError(err error) bool {
	if err != nil {
		panic(err)
	}
	return true
}
