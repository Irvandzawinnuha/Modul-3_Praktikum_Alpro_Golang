package main

import "fmt"

func main() {
  var masukan, jumlah, perulangan, rerata, tertinggi, terendah int = 0, 0, 0, 0, 0, 0

  fmt.Scan(&masukan)
  tertinggi = masukan
  terendah = masukan

  for masukan > 0 {
    jumlah += masukan
    perulangan++
    if tertinggi < masukan {
      tertinggi = masukan
    }
    if terendah > masukan {
      terendah = masukan
    }
    fmt.Scan(&masukan)
  }

  rerata = jumlah / perulangan

  fmt.Println("Jumlah = ", jumlah)
  fmt.Println("Rata - rata = ", rerata)
  fmt.Println("Uang tertinggi = ", tertinggi)
  fmt.Println("Uang terendah = ", terendah)
} 