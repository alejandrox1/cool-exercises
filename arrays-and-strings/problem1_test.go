package main

import (
	"testing"
)

var input = `1234567890-=~!@#$%^*()_+qwertyuiop[]\asdfghjkl;'zxcvbnm,./QWERTYUIOP{}|
ASDFGHJKL:"ZXCVBNM<>?電电電買买車车車紅红紅無无東东馬马風风時时鳥鸟語语頭头魚鱼園园長长島岛愛爱紙纸書书見见假佛仏德拜黑冰兔兎妒每毎壤步巢區区五十音,ごじゅうおん外来語,がいらい`

func BenchmarkBruteForce(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bruteForce(input)
	}
}

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isUniqueSort(input)
	}
}

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isUniqueMap(input)
	}
}
