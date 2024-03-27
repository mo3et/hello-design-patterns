package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
责任链模式 请假审批
https://kamacoder.com/problempage.php?pid=1100
*/

// 处理者：定义接口
type LeaveHandler interface {
	HandleRequest(request LeaveRequest)
}

// 具体处理者：可以有多个，负责具体处理，这里分为 Supervisor、Manager、Director
type Supervisor struct {
	nextHandler LeaveHandler
}

const maxDaysSupervisorCanApprove = 3

func NewSupervisor(nextHandler LeaveHandler) *Supervisor {
	return &Supervisor{nextHandler: nextHandler}
}

func (s *Supervisor) HandleRequest(request LeaveRequest) {
	if request.Days <= maxDaysSupervisorCanApprove {
		fmt.Println(request.Name + " Approved by Supervisor.")
	} else if s.nextHandler != nil {
		s.nextHandler.HandleRequest(request)
	} else {
		fmt.Println(request.Name + " Denied by Supervisor.")
	}
}

type Manager struct {
	nextHandler LeaveHandler
}

const maxDaysManagerCanApprove = 7

func NewManager(nextHandler LeaveHandler) *Manager {
	return &Manager{nextHandler: nextHandler}
}

func (m *Manager) HandleRequest(request LeaveRequest) {
	if request.Days <= maxDaysManagerCanApprove {
		fmt.Println(request.Name + " Approved by Manager.")
	} else if m.nextHandler != nil {
		m.nextHandler.HandleRequest(request)
	} else {
		fmt.Println(request.Name + " Denied by Manager.")
	}
}

type Director struct{}

const maxDaysDirectorCanApprove = 10

func (d *Director) HandleRequest(request LeaveRequest) {
	if request.Days <= maxDaysDirectorCanApprove {
		fmt.Println(request.Name + " Approved by Director.")
	} else {
		fmt.Println(request.Name + " Denied by Director.")
	}
}

// 请求类
type LeaveRequest struct {
	Name string
	Days int
}

// 主函数
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取用户输入
	scanner.Scan()
	var n int
	fmt.Sscanf(scanner.Text(), "%d", &n)

	// 组装职责链
	director := &Director{}
	manager := NewManager(director)
	supervisor := NewSupervisor(manager)

	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Fields(scanner.Text())

		if len(input) == 2 {
			name := input[0]
			days, err := strconv.Atoi(input[1])
			if err != nil {
				fmt.Println("Invalid input")
				return
			}

			request := LeaveRequest{Name: name, Days: days}
			supervisor.HandleRequest(request)
		} else {
			fmt.Println("Invalid input")
			return
		}
	}
}
