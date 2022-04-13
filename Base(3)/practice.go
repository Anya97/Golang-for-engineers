package main

import (
	"fmt"
)

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type HealthCheck struct {
	ServiceID int
	status    string
}

func main() {
	fmt.Println("Тут будет выведен идентификатор")
	HealthCheckSlice := GenerateCheck()
	for _, v := range HealthCheckSlice {
		if v.status == PassStatus {
			fmt.Println(v.ServiceID)
		}
	}
}

func GenerateCheck() []HealthCheck {
	HealthCheckSlice := make([]HealthCheck, 0)
	var status string
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			status = PassStatus
		} else {
			status = FailStatus
		}
		HealthCheckSlice = append(HealthCheckSlice, HealthCheck{ServiceID: i, status: status})
	}
	return HealthCheckSlice
}
