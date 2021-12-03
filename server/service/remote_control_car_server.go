package service

import (
	"fmt"

	"github.com/fangzhixi/fizzy-server/server/model"
	"github.com/fangzhixi/fizzy-server/server/utils"
)

/*
 * @Author       : zhixi.fang (Pop)
 * @Date         : 2021-12-03 16:00:39
 * @LastEditors  : zhixi.fang (Pop)
 * @LastEditTime : 2021-12-03 16:11:55
 */

func (r *RemoteControlService) RcCarControl(signal *model.RemoteControlCar) error {
	var (
		logId = utils.NewLogId()
	)
	fmt.Println(logId)
	return nil
}
