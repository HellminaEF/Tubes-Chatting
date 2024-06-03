package main

import "fmt"

const(
	MAX_USERS = 100
	MAX_MESSAGES = 200
)

type User struct{
	ID int
	Username string
	Password string
	Menyetujui bool
}

type Message struct{
	Username_Pengirim string
	Username_Penerima string
	Konten string
}

var users [MAX_USERS]User
var jumlah_User int
var messages [MAX_MESSAGES]Message
var jumlah_Message int

//Procedure Registrasi
func registrasi(username, password string){
	User_Baru := User{
		ID: jumlah_User + 1,
		Username: username,
		Password: password,
		Menyetujui: false,
	}

	users[jumlah_User] = User_Baru
	jumlah_User++
	fmt.Printf("Registrasi akun %s berhasil silahkan tunggu persetujuan dari Admin.\n", username)
}

//Procedur menyetujui pembuatan akun oleh Admin
func menyetujui(username string){
	for i := 0; i < jumlah_User; i++{
		if users[i].Username == username{
			users[i].Menyetujui = true
			fmt.Printf("Akun %s telah disetujui oleh Admin.\n", users[i].Username)
			return
		}
	}
	fmt.Println("Username tidak valid")
}

//Procedure penolakan pembuatan akun oleh Admin
func menolak(username string){
	for i := 0; i < jumlah_User; i++{
		if users[i].Username == username{
			fmt.Printf("Akun %s ditolak.\n", users[i].Username)
			for j := i; j < jumlah_User - 1; j++{
				users[j] = users[j - 1]
			}
			jumlah_User--
			return
		}
	}
	fmt.Println("Username tidak valid.")
}