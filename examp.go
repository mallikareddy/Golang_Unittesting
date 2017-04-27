package main
import "fmt"
func main()
{
   var num = []int{10,20,30,40,50,60,70,80,90,200}
   var sum int
   var m int
   m := len(num)
   for i := 0; i <= m; i++
   {
       sum+=num[i]
       i--
   }
   fmt.printf("sum is %d",sum)
}
