package shared

import "time"

// Next 下一个复习的时间点
// https://stackoverflow.com/questions/49047159/spaced-repetition-algorithm-from-supermemo-sm-2/49047160#49047160
// spaced repetition 间隔重复：一种学习技巧
//  repetions 用户已经复现的次数，0 ~ N。默认0
//	quality 用户感觉的难度
// 		5 - perfect response;
// 		4 - correct response; 	after a hesitation
// 		3 - correct response; 	recalled with serious difficulty
// 		2 - incorrect response; where the correct one seemed easy to recall
// 		1 - incorrect response; the correct one remembered
// 		0 - complete blackout
//	interval 每个重复重复之间的时长。 默认1
// 	easiness 舒适度，用来调节 interval。默认2.5
//
// EF':=EF+(0.1-(5-q)*(0.08+(5-q)*0.02))
//
// 艾宾浩斯曲线只是 repetions => interval

// 返回间隔系数
// func Next(repetions int64, quality int8) float64 {
// 	if quality < 3 { // 错了
// 		return 1
// 	}

// 	if repetions == 0 { // 从来没有复现过
// 		return 1
// 	}

// 	if repetions == 1 { // 复习过一次
// 		return 6
// 	}

// 	q := float64(5 - quality)
// 	// easiness := 2.6 - q*(0.08+q*0.02)
// 	if t := (8*q + 2*q*q) / 100; t > 1.3 {
// 		return t
// 	}

// 	return 1.3
// }

// Next 下一个复习的时间点
//	quality 用户感觉的难度
// 		0 容易：按计划进行， repetion+1
// 		1 一般：停留在当前的地点
//		2 困难：repetion -1
// 		3 重来：repetion = 0
func Next(curStage, quality int8) (time.Duration, int8) {
	// 这个是艾宾浩斯记忆曲线的复现计划
	// 这个没有考虑到当前用户记忆情况。比如用户觉得很容易和觉得很难都一样
	list := []time.Duration{
		time.Minute * 5,
		time.Minute * 30,
		time.Minute * 60 * 12,
		time.Minute * 60 * 24,
		time.Minute * 60 * 24 * 2,
		time.Minute * 60 * 24 * 4,
		time.Minute * 60 * 24 * 7,
		time.Minute * 60 * 24 * 15,
	}

	if quality == 0 {
		curStage++
	} else if quality == 1 {
		// DO nothing
	} else if quality == 2 {
		curStage--
	} else if quality == 3 {
		curStage = 0
	}

	if curStage < 0 {
		return list[0], 0
	}

	if curStage > 7 {
		return 0, -1
	}

	return list[curStage], curStage
}
