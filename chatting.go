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
	Disetujui bool
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

//Procedure Registrasi (Hellmina Enjelina Fitri)
func register(username, password string){
	User_Baru := User{
		ID: jumlah_User + 1,
		Username: username,
		Password: password,
		Disetujui: false,
	}

	users[jumlah_User] = User_Baru
	jumlah_User++
	fmt.Printf("Registrasi akun %s berhasil silahkan tunggu persetujuan dari Admin.\n", username)
}

//Procedur Disetujui pembuatan akun oleh Admin (Hellmina Enjelina Fitri)
func disetujui(username string){
	for i := 0; i < jumlah_User; i++{
		if users[i].Username == username{
			users[i].Disetujui = true
			fmt.Printf("Akun %s telah disetujui oleh Admin.\n", users[i].Username)
			return
		}
	}
	fmt.Println("Username tidak valid")
}

//Procedure penolakan pembuatan akun oleh Admin (Hellmina Enjelina Fitri)
func ditolak(username string){
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
	fmt.Println("Username tidak valid")
}

//Procedure untuk print daftar user yang sudah disetujui oleh Admin (Hellmina Enjelina Fitri)
func printUsers(){
	fmt.Println("Daftar user yang sudah disetujui :")
	for i := 0; i < jumlah_User; i++{
		if users[i].Disetujui{
			fmt.Printf("Id : %d, Username : %s\n", users[i].ID, users[i].Username)
		}
	}
}

//Procedure untuk kirim pesan pribadi (Hellmina Enjelina Fitri)
func kirimPesan(username_Pengirim, username_Penerima, konten string){
	user_Disetujui, penerima_Disetujui := false, false
	for i := 0; i < jumlah_User; i++ {
		if users[i].Username == username_Pengirim && users[i].Disetujui {
			user_Disetujui = true
		}
		if users[i].Username == username_Penerima && users[i].Disetujui {
			penerima_Disetujui = true
		}
	}
	if user_Disetujui && penerima_Disetujui {
		newMessage := Message{
			Username_Pengirim: username_Pengirim,
			Username_Penerima: username_Penerima,
			Konten: konten,
		}
		messages[jumlah_Message] = newMessage
		jumlah_Message++
		fmt.Println("Pesan berhasil dikirim.")
	} else {
		fmt.Println("Salah satu akun atau keduanya belum disetujui.")
	}
}


//fungsi main (Hellmina Enjelina Fitri)
func main(){
	var choice int
	for{
		fmt.Println()
		fmt.Println("-------- Menu Utama --------")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Admin")
		fmt.Println("4. Keluar")
		fmt.Println("----------------------------")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&choice)

		if choice == 1 {
			var username, password string
			fmt.Print("Masukkan username: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scan(&password)
			register(username, password)
		} else if choice == 2{
			var username, password string
			fmt.Print("Masukan username: ")
			fmt.Scan(&username)
			fmt.Print("Masukan password: ")
			fmt.Scan(&password)

			validLogin := false
			for i := 0; i < jumlah_User; i++ {
				if users[i].Username == username && users[i].Password == password && users[i].Disetujui {
					validLogin = true
					break
				}
			}

			if validLogin{
				var subChoice int
				for {
					fmt.Println()
					fmt.Println("-------- Menu Pengguna --------")
					fmt.Println("1. Kirim Pesan")
					fmt.Println("2. Buat Grup")
					fmt.Println("3. Tambah Anggota Grup")
					fmt.Println("4. Kirim Pesan ke Grup")
					fmt.Println("5. Lihat Anggota Grup")
					fmt.Println("6. Keluar")
					fmt.Println("-------------------------------")
					fmt.Print("Pilih Opsi: ")
					fmt.Scan(&subChoice)

					if subChoice == 6{
						break
					}

					if subChoice == 1{
						var username_Penerima, konten string
						fmt.Print("Masukan username tujuan: ")
						fmt.Scan(&username_Penerima)
						fmt.Print("Masukan pesan: ")
						fmt.Scan(&konten)
						kirimPesan(username, username_Penerima, konten)
					} else if subChoice == 2{
						// Buat Grup
					} else if subChoice == 3{
						// Tambah anggota grup
					} else if subChoice == 4{
						// Kirim pesan ke grup
					} else if subChoice == 5{
						// Lihat anggota grup
					} else {
						fmt.Println("Opsi tidak valid, silakan pilih lagi.")
					}
				}
			} else {
				fmt.Println("Username tidak valid, Password salah, atau Akun belum disetujui.")
			}
		} else if choice == 3{
			var admin_Choice int
			for {
				fmt.Println()
				fmt.Println("-------- Menu Admin --------")
				fmt.Println("1. Setujui Akun")
				fmt.Println("2. Tolak Akun")
				fmt.Println("3. Cetak Daftar Akun yang Disetujui")
				fmt.Println("4. Kembali ke Menu Utama")
				fmt.Println("----------------------------")
				fmt.Print("Pilih Opsi: ")
				fmt.Scan(&admin_Choice)

				if admin_Choice == 4{
					break
				}

				if admin_Choice == 1{
					var username string
					fmt.Print("Masukan username yang ingin di setujui: ")
					fmt.Scan(&username)
					disetujui(username)
				} else if admin_Choice == 2{
					var username string
					fmt.Print("Masukan username yang ingin di tolak: ")
					fmt.Scan(&username)
					ditolak(username)
				} else if admin_Choice == 3{
					printUsers()
				} else{
					fmt.Print("Opsi tidak valid, silahkan pilih lagi.")
				}
			}
		} else if choice == 4{
			fmt.Println("Keluar dari program.")
			return
		} else {
			fmt.Println("Opsi tidak valid, silahkan pilih lagi.")
		}
	}
}