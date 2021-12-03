/*
 * @Author       : zhixi.fang (Pop)
 * @Date         : 2021-12-03 11:09:02
 * @LastEditors  : zhixi.fang (Pop)
 * @LastEditTime : 2021-12-03 16:15:24
 */
package server_test

import (
	"fmt"
	"testing"

	"github.com/fangzhixi/fizzy-server/server/config"
	"github.com/fangzhixi/fizzy-server/server/core"
	"github.com/fangzhixi/fizzy-server/server/utils"
)

func TestLongConnectionServerMyself(t *testing.T) {
	var (
		logId    = utils.NewLogId()
		address  = config.Config.RConfig.ListenAddress
		poolSize = config.Config.RConfig.MaxLongConnectionPoolSize
	)
	fmt.Println("服务端")

	server, err := core.NewLongConnServer(logId, address, poolSize)
	if err != nil {
		t.Fatal(err)
	}
	err = server.CreateTcpListering()
	if err != nil {
		t.Fatal(err)
	}
}
