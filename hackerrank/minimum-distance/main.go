package main
import "fmt"

func main() {
    var nb int
    fmt.Scan(&nb)    
    mp := make(map[int]int)
    minDist := nb+1
    num :=0
    for i:=0;i<nb;i++{
        fmt.Scan(&num)
        if v,ok:=mp[num];ok{
            if dist:=i-v;dist<minDist {
                minDist = dist
            }
        }
        mp[num] = i
    }
    if minDist>nb {
        minDist = -1
    }
    fmt.Println(minDist)
}