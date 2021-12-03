package controller

/*
 * @Author       : zhixi.fang (Pop)
 * @Date         : 2021-12-03 16:00:39
 * @LastEditors  : zhixi.fang (Pop)
 * @LastEditTime : 2021-12-03 16:13:09
 */

import (
	"fmt"

	"github.com/fangzhixi/fizzy-server/server/model"
	"github.com/fangzhixi/fizzy-server/server/utils"
)

func (r *RemoteControlServer) RcCarControl(signal *model.RemoteControlCar) error {
	var (
		logId = utils.NewLogId()
	)
	fmt.Println(logId)
	return nil
}
