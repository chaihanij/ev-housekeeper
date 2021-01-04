package main

import (
	"fmt"

	"github.com/pepsi7959/ev-housekeeper/pkg/fs"
	"github.com/spf13/viper"
)

//EVHouseKeeper contains all parameters to write reports
type EVHouseKeeper struct {
	conf *viper.Viper
}

func main() {
	hk := &EVHouseKeeper{}
	hk.conf = ReadConf()

	w := fs.Writer{FileName: hk.conf.GetString("filename"), Perm: 0644}
	_, err := w.Write("pepsi\n", 6)
	if err != nil {
		fmt.Println(err.Error())
	}
}
