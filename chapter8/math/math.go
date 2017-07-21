package math

//Calculates the Average of a set of numbers.
func Average (numbers ...float64) float64 {

  total := 0.0
  for _, n:= range numbers {
    total += n
  }

  return total
}
