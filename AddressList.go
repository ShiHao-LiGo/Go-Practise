package main

import "fmt"

type Person struct {
	userName     string
	addressPhone map[string]string
}

var personList = make([]Person, 0)

func addPerson() {
	//1、定义结构体表示多人信息s
	var exit string //表示退出电话的录入
	var name string
	var address string
	var phone string
	var addressPhone = make(map[string]string) //保存电话类型和电话
	fmt.Println("请输入姓名")
	fmt.Scan(&name)
	for {
		fmt.Println("请输入电话类型")
		fmt.Scan(&address)
		fmt.Println("请输入电话号码")
		fmt.Scan(&phone)
		addressPhone[address] = phone
		fmt.Println("如果结束电话的录入，请按Q")
		fmt.Scan(&exit)
		if exit == "Q" {
			break
		} else {
			continue
		}
	}

	//将联系人的信息存储到切片中
	personList = append(personList, Person{userName: name, addressPhone: addressPhone})
	//fmt.Println(personList)
	//调用函数显示联系人信息
	showPersonList()
}

func removePerson() {
	var name string
	var index int = -1
	fmt.Println("请输入需要删除的联系人的姓名:")
	fmt.Scan(&name)
	for i := 0; i < len(personList); i++ {
		if personList[i].userName == name {
			index = i
			break
		}
	}
	if index != -1 {
		personList = append(personList[:index], personList[index+1:]...) //append函数第二个参数如果是切片后面要加三个点
	}

	showPersonList()

}

func findPerson() *Person {
	var name string
	var index = -1
	fmt.Println("请输入要查询的联系人的姓名：")
	fmt.Scan(&name)
	for k, value := range personList {
		if value.userName == name {
			index = k
			fmt.Println("联系人姓名：", value.userName)
			for key, value := range value.addressPhone {
				fmt.Printf("%s:%s\n", key, value)
			}
		}
	}
	if index == -1 {
		fmt.Println("没有找到联系人信息")
		return nil
	} else {
		return &personList[index]
	}
}

func editPerson() {
	var name string
	var p *Person
	var num int                  // 存储修改的数据的类型
	var menu = make([]string, 0) //保存电话类型，方便修改
	var pNum int
	p = findPerson()
	if p != nil {
		for {
			fmt.Println("编辑用户名称请按:5，编辑电话请按:6,退出请按:7")
			fmt.Scan(&num)
			switch num {
			case 5:
				fmt.Println("请输入新的姓名:")
				fmt.Scan(&name)
				p.userName = name
				showPersonList()
			case 6:
				var j int
				for key, value := range p.addressPhone {
					fmt.Println("编辑(", key, ")", value, "请按:")
					menu = append(menu, key)
					j++
				}
				fmt.Println("请输入编辑号码的类型:")
				fmt.Scan(&pNum)

				var phone string
				for index, v := range menu {
					if index == pNum {
						fmt.Println("请输入新的电话号码：")
						fmt.Scan(&phone)
						p.addressPhone[v] = phone
					}
				}

			}
		}
	} else {
		fmt.Println("没有找到要编辑的人的信息")
	}
}

func showPersonList() {
	if len(personList) == 0 {
		fmt.Println("暂时没有联系人信息")
	} else {
		for _, value := range personList {
			fmt.Println("姓名：", value.userName)
			for k, v := range value.addressPhone {
				fmt.Println("电话类型:", k)
				fmt.Println("电话号码", v)
			}
		}
	}
}

func main() {
	for {
		scanNum()
	}
}

func scanNum() {
	//1、给出相应的提示
	fmt.Println("添加联系人信息，请按1")
	fmt.Println("删除联系人信息，请按2")
	fmt.Println("查询联系人信息，请按3")
	fmt.Println("编辑联系人信息，请按4")
	//2、对用户输入的数字进行判断
	var num int
	fmt.Scan(&num)
	switchType(num)

}
func switchType(n int) {
	switch n {
	case 1:
		//添加联系人
		addPerson()
	case 2:
		//删除联系人信息
		removePerson()
	case 3:
		//查询联系人信息
		findPerson()
	case 4:
		//修改联系人信息
		editPerson()

	}
}
