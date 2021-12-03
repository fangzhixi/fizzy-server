package model

/*
 * @Author       : zhixi.fang (Pop)
 * @Date         : 2021-12-03 15:45:07
 * @LastEditors  : zhixi.fang (Pop)
 * @LastEditTime : 2021-12-03 15:54:18
 */

type PWM int64

type RemoteControlCar struct {
	Motor          PWM `json:"motor"`
	SteeringEngine PWM `json:"steering_engine"`
}

func (p RemoteControlCar) GetMotorPWM() int64 {
	return int64(p.Motor)
}

func (p RemoteControlCar) GetSteeringEnginePWM() int64 {
	return int64(p.SteeringEngine)
}
