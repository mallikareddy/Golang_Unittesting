package main
import "fmt"
func main() {
   var num = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 200}
   var sum int
   m := len(num)
   for i := 0; i <= m; i++ {
      fmt.Println("num[%d]:%d",i,num[i])
       sum = sum + num[i]
   }
   fmt.Printf("sum is %d",sum)
}
