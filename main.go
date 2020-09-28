package package1
//Function sum provides an interface for as many as supplied ints to be summed/added together.
func Sum(xi ...int) int{
 sum := 0
 for _, v := range xi{
  sum += v
 }
  return sum;
}
