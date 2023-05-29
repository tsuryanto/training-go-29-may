package main

import (
	"fmt"
	"gobasic/employee"
)

func main() {
	// fmt.Println("Hello World")

	// // Variable
	// var str string = "Hello"
	// // var str1 = "World"
	// // str2 := "string"

	// // var (
	// // 	nomor1       int
	// // 	nomor2       int
	// // 	nama, alamat string
	// // )

	// /* const phi = 3.14
	// Ini tidak akan diekseskusi
	// */

	// fmt.Println(str)

	// // int float
	// var angka int = 1
	// var koma float32 = 2.6

	// perkalian := angka * int(koma)
	// fmt.Println(perkalian)

	// perkalianKoma := float32(angka) * koma
	// fmt.Println(perkalianKoma)

	// // bool
	// var canAccess bool = true
	// if canAccess {
	// 	fmt.Println("Bisa Akses")
	// }

	// // Operator
	// a := 10
	// var b int = 3
	// b *= a // b = b * a
	// fmt.Println(b)

	// if b%2 == 0 {
	// 	fmt.Println("genap")
	// } else if b < 0 {
	// 	fmt.Println("negative")
	// } else {
	// 	fmt.Println("ganjil")
	// }

	// var status = "open"
	// switch status {
	// case "open":
	// 	fmt.Println(1)
	// case "close":
	// 	fmt.Println(0)
	// }

	// if !(a == 0 && b != 0) || (a == 4) {
	// 	fmt.Println("Ok")
	// }

	// // array
	// var buah [3]interface{} = [3]interface{}{"Apel", 3, "pir"}

	// // if len(buah) < 3 {
	// // 	fmt.Println(buah[3])
	// // }

	// for _, b := range buah {
	// 	// bConvert := b.(string)
	// 	// manis := bConvert + "manis"
	// 	fmt.Println(b)
	// }

	// // slice
	// var buah2 []interface{} = []interface{}{"Apel", 3, "pir", "strowberry"}
	// fmt.Println(buah2[1:3])

	// // map
	// var mapBuah = make(map[string]string)
	// mapBuah["buah pasar"] = "mangga"
	// mapBuah["buah manis"] = "jeruk"

	// dataMahasiswa := map[string]string{
	// 	"Nama":   "Budi",
	// 	"Alamat": "Tangerang",
	// }
	// fmt.Println(dataMahasiswa["Nama"])

	// for key, mhsValue := range dataMahasiswa {
	// 	fmt.Println(key, mhsValue)
	// }

	// var buah2an []string = []string{"Apel", "pir", "strowberry"}
	// for _, buah := range buah2an {
	// 	fmt.Println(buah)
	// }

	// var name string  // ""
	// var usia int     // 0
	// var canFill bool // false
	// var cobaMap map[string]string

	// if !canFill {
	// 	fmt.Println("Gak bisa akses")
	// }
	// fmt.Println(name, usia, canFill, cobaMap)

	// defer
	// defer fmt.Println("Hello")
	// fmt.Println("Selamat Pagi")
	// defer fmt.Println("World")
	// fmt.Println("Hai")

	// employee := employee.Employee{
	// 	EmpNo: 1,
	// 	Name:  "Agus",
	// }
	// fmt.Println(employee)
	// fmt.Println(employee.GetName())
	// fmt.Println(employee.GetEmpNo())

	// PendaftaranCPNS(&employee)
	// fmt.Println(employee)

	// x := 10
	// var y *int = &x
	// x = 20
	// fmt.Println(*y)

	// *y = 30
	// fmt.Println(x)

	// var pointernul *int
	// fmt.Println(*pointernul)

	// employee.CountAge()
	// employee.GetEmpNo()

	// error
	// manager := employee.Manager{}
	// manager.AksesMasuk()

	// masDikta := employee.NewProgrammer("Dikta", "Vscode")
	// masDikta.AksesMasuk()

	// masDikta.PakaiPakaian(pakaian.Kemeja{Ukuran: 18})
	// masDikta.PakaiPakaian(pakaian.Kaos{Ukuran: 16})
	// masDikta.PakaiPakaian(pakaian.Sweater{Ukuran: 16})

	// fmt.Println("Hello")
	// go func() {
	// 	time.Sleep(100 * time.Millisecond)
	// 	fmt.Println("done")
	// }()

	// wg := &sync.WaitGroup{}

	// makanan := []string{"Soto", "Bakso", "Nasi Goreng"}
	// for _, m := range makanan {
	// 	wg.Add(1)
	// 	go func(food string) {
	// 		fmt.Println(food)
	// 		wg.Done()
	// 	}(m)
	// }

	// wg.Wait()

	// var once sync.Once

	// var yangmausayamakan string
	// for _, m := range makanan {
	// 	once.Do(
	// 		func() {
	// 			yangmausayamakan = m
	// 		},
	// 	)
	// }
	// fmt.Println(yangmausayamakan)

	// // Race Condition
	// var x int
	// wg := &sync.WaitGroup{}
	// mu := sync.Mutex{}

	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		for j := 0; j < 1000; j++ {
	// 			mu.Lock()
	// 			x++ // x = x + 1
	// 			mu.Unlock()
	// 		}
	// 		wg.Done()
	// 	}()
	// }

	// // saldo 1000
	// // uang masuk 1000 = 2000

	// // narik 500
	// // 1000 - 500 = 500

	// // saldo akhir 2000

	// // x -> 0 , 1
	// // x -> 0 , 1

	// wg.Wait()
	// fmt.Println(x)

	ch := make(chan string)
	defer close(ch)

	makanan := []string{"Soto", "Bakso", "Nasi Goreng"}
	for _, m := range makanan {
		go func(food string) {
			ch <- food
		}(m)
	}

	for a := range ch {
		fmt.Println(a)
	}
}

func PendaftaranCPNS(emp *employee.Employee) {
	emp.Name = "Agus Kurniawan"
}

func GetName() string {
	return ""
}

type Freelancer struct {
}

func LombaBasket(employee employee.Employee) {

}
