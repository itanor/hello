package arrays

func Sum(numbers []int) int {
  sum := 0
  for _, number := range numbers {
    sum += number
  }
  return sum
}

func SumAll(numbersToSum ...[]int) (sum []int) {
  for _, numbers := range numbersToSum {
    total := Sum(numbers)
    sum = append(sum, total)
  }
  return
}

func SumAllTails(numbersToSum ...[]int) (sum []int) {
  for _, numbers := range numbersToSum {
    idx := 0
    if len(numbers) > 1 {
      idx = 1
    }
    total := Sum(numbers[idx:])
    sum = append(sum, total)
  }
  return
}

