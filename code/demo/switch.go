package main

import "code/P"

type Gender int8

const (
	MALE   Gender = 1
	FEMALE Gender = 2
)

func selectSwitch(){
	P.PE("selectSwitch")
	gender := FEMALE
	P.D(MALE)
	P.D(gender)
	switch gender {
		case FEMALE:
			P.D("female")
		case MALE:
			P.D("male")
		default:
			P.D("unknown")
		
	}
}

func selectSwitchFallthrough(){
	P.PE("selectSwitchFallthrough")
	gender := MALE
	P.D(MALE)
	P.D(gender)
	switch gender {
		case FEMALE:
			P.D("female")
			fallthrough
		case MALE:
			P.D("male")
			fallthrough
		default:
			P.D("unknown")
		
	}
}


func main() {
	selectSwitch()
	selectSwitchFallthrough()
}

/*
docker exec go_c go run demo/switch.go
*/

/*
1. Go 语言的 switch 不需要 break，匹配到某个 case，执行完该 case 定义的行为后，
默认不会继续往下执行。如果需要继续往下执行，需要使用 fallthrough

 */
