package main

import (
	"fmt"
	"math"
)

/**
변수를 선언을 위해 var 을 사용 합니다.
함수의 매겨변수 처럼 타임은 문장 끝에 명시 합니다.
*/
var x, y, z = 1, 2, 3                    //변수 선언과 함께 변수 각각을 초기화를 할수 있습니다.
var c, python, java = true, false, "no!" //초기화를 하는 경우 타입(type)을 생략할수 있습니다. 변수는 초기화 하고자 하는 값에 따라 타입이 결정 됩니다.

//상수는 const 키워드와 함께 변수 처럼 선언합니다.
const Pi = 3.14

//숫자형 상수는 정밀한 값(values) 을 표현 할수 있습니다.
//타입을 지정 하지 않은 상수는 문맥(context)에 따라 타입을 가지게 됩니다.
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string, string) {
	return y, x, x + y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello", math.Pi, "World")
	fmt.Println("Now you have %g problems.", math.Nextafter(2, 3))
	fmt.Println(add(4, 2))
	a, b, c := swap("hello", "world")
	fmt.Println(a, b, c)
	fmt.Println(split(17))
	var x, y, z int = 1, 2, 3
	c1, python, java := true, false, "no!" //함수 내에서 := 을 사용하면 var 과 명시적인 타입 (e.g. int. bool) 을 생략 할수 있습니다.
	//그러나 함수 빡에서는 := 선언을 할 수 없습니다.
	fmt.Println(x, y, z, c1, python, java)

	//상수는 문자(character), 문자열(string), 부울(boolean), 숫자 타임 중의 하나가 될수 있습니다.
	const World = "안녕"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	//코드를 통해 needInt(big) 는 어떤 결과를 출력 할지 실험 해보세요
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
