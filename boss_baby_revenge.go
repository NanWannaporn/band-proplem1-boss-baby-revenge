package main

/*
check Boss Baby Behavior is a Good boy
1. string not start with R
2. after shot need to revenge
3. every single shot must be revenged by Boss Baby at least one with not intiated shooting
*/

func BossBabayRevenge(input string) string {
	var shotCnt, revengeCnt = 0, 0

	if len(input) < 1 || len(input) > 1000000 {
		return "invalid legnth"
	}

	for i, cha := range input {
		char := string(cha)

		if char == "S" {
			shotCnt++
		} else if char == "R" {
			if i == 0 { //can't start by revenge
				return "Bad boy"
			}

			if shotCnt > 0 {
				revengeCnt++
			}

			if revengeCnt > shotCnt { //reset --ถ้า r มากกว่า s แล้ว เริ่มนับใหม่ถือว่า แก้แค้นรอบนั้นครบแล้ว
				revengeCnt = shotCnt
				// revengeCnt = 0
				// shotCnt = 0
			}
		} else {
			return "Invalid" //invalid characters
		}
	}

	//at the end, must have revenged every shot
	if revengeCnt < shotCnt { //have to revenge at least one shot for every S
		return "Bad boy"
	}
	return "Good boy"

}
