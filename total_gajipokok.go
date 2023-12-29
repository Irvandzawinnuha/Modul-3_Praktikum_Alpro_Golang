package main

import "fmt"

func main() {
  var golongan string
  var jam_regular, jam_lembur, gaji_regular, gaji_lembur int
  var total_gaji int = 0

  fmt.Scan(&golongan, &jam_regular, &jam_lembur)

  for golongan != "Z" {
    if golongan == "A" {
      gaji_regular = (jam_regular * 75000) + (jam_lembur * 90000)
      total_gaji += gaji_regular
      fmt.Println(gaji_regular)
    } else if golongan == "B" {
      gaji_lembur = (jam_regular * 125000) + (jam_lembur * 70000)
      total_gaji += gaji_lembur
      fmt.Println(gaji_lembur)
    } else {
      fmt.Println("Golongan tidak dikenali")
    }

    fmt.Scan(&golongan, &jam_regular, &jam_lembur)
  }

  fmt.Print("Biaya yang dikeluarkan perusahaan untuk gaji karyawan ", total_gaji)
}