package iteration

func Repeat(character string, times ...int) (repeated string) {
    repeatTimes := 2

    if len(times) > 0 {
        repeatTimes = times[0]
    }

    for i:=0; i < repeatTimes; i++ {
        repeated += character
    }
    return
}
