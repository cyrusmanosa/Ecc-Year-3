package main

import "fmt"

type myStruct struct {
    field int
}

func main() {
    // 構造体は値型であるため、s1とs2は別々のメモリに値を格納する
    s1 := myStruct{1}
    s2 := s1
    s2.field = 2
    fmt.Println(s1, s2)  // 別々の値が表示される

    // ポインタを用いると同じメモリ領域が使用される
    s3 := &s1
    //ポインタ型でも、「.」で構造体のフィールドを参照可能
    s3.field = 3         // 「(*s3).field = 3」と同じ意味
    fmt.Println(s1, *s3) // 同じ値が表示される
}
