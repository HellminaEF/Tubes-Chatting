package main

import "fmt"

const(
	MAX_USERS = 100
	MAX_MESSAGES = 200
)

type User struct{
	ID int
	Name string
	Username string
	Password string
	Approved bool
}

type Message struct{
	Sender string
	Receiver string
	Content string
}

var users [MAX_USERS]User
var userCount int
var messages [MAX_MESSAGES]Message
var messageCount int

//Fungsi main (Hellmina Enjelina Fitri)
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
			var name, username, password string
			fmt.Print("Masukkan Nama: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan username: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scan(&password)
			register(name, username, password)
		} else if choice == 2{
			var username, password string
			fmt.Print("Masukan username: ")
			fmt.Scan(&username)
			fmt.Print("Masukan password: ")
			fmt.Scan(&password)

			validLogin := false
			for i := 0; i < userCount; i++ {
				if users[i].Username == username && users[i].Password == password && users[i].Approved {
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
						var receiver, content string
						fmt.Print("Masukan username tujuan: ")
						fmt.Scan(&receiver)
						fmt.Print("Masukan pesan: ")
						fmt.Scan(&content)
						sendPrivateMessage(username, receiver, content)
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
				fmt.Println("Username tidak valid, Password salah, atau Akun belum Approved.")
			}
		} else if choice == 3{
			var admin_Choice int
			for {
				fmt.Println()
				fmt.Println("-------- Menu Admin --------")
				fmt.Println("1. Setujui Akun")
				fmt.Println("2. Tolak Akun")
				fmt.Println("3. Cetak Daftar Akun yang Approved")
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
					approved(username)
				} else if admin_Choice == 2{
					var username string
					fmt.Print("Masukan username yang ingin di tolak: ")
					fmt.Scan(&username)
					rejected(username)
				} else if admin_Choice == 3{
					printUsers()
				} else{
					fmt.Print("Opsi tidak valid, silahkan pilih lagi.")
				}
			}
		} else if choice == 4{
			fmt.Println("Keluar dari program...")
			return
		} else {
			fmt.Println("Opsi tidak valid, silahkan pilih lagi.")
		}
	}
}

//Fungsi Registrasi (Hellmina Enjelina Fitri)
func register(name, username, password string){
	for i := 0; i < userCount; i++ {
		if users[i].Username == username {
			fmt.Printf("Username %s sudah digunakan. Silakan buat dengan username lain.\n", username)
			return
		}
	}

	new_User := User{
		ID: userCount + 1,
		Name: name,
		Username: username,
		Password: password,
		Approved: false,
	}

	users[userCount] = new_User
	userCount++
	fmt.Printf("Registrasi akun %s berhasil silahkan tunggu persetujuan dari Admin.\n", username)
}

//Fungsi disetujui pembuatan akun oleh Admin (Hellmina Enjelina Fitri)
func approved(username string){
	for i := 0; i < userCount; i++{
		if users[i].Username == username{
			users[i].Approved = true
			fmt.Printf("Akun %s telah disetujui oleh Admin.\n", users[i].Username)
			return
		}
	}
	fmt.Println("Username tidak ditemukan")
}

//Fungsi penolakan pembuatan akun oleh Admin (Hellmina Enjelina Fitri)
func rejected(username string){
	for i := 0; i < userCount; i++{
		if users[i].Username == username{
			fmt.Printf("Akun %s ditolak.\n", users[i].Username)
			for j := i; j < userCount - 1; j++{
				users[j] = users[j - 1]
			}
			userCount--
			return
		}
	}
	fmt.Println("Username tidak ditemukan")
}

//Fungsi untuk print daftar user yang sudah Approved oleh Admin (Hellmina Enjelina Fitri)
func printUsers(){
	fmt.Println("Daftar user:")
	for i := 0; i < userCount; i++ {
		status := "Belum disetujui"
		if users[i].Approved {
			status = "Sudah disetujui"
		}
		fmt.Printf("ID: %d, Nama: %s, Username: %s, Status: %s\n", users[i].ID, users[i].Name, users[i].Username, status)
	}
}

//Fungsi untuk kirim pesan pribadi (Hellmina Enjelina Fitri)
func sendPrivateMessage(sender, receiver, content string){
	user_Approved, penerima_Approved := false, false
	for i := 0; i < userCount; i++ {
		if users[i].Username == sender && users[i].Approved {
			user_Approved = true
		}
		if users[i].Username == receiver && users[i].Approved {
			penerima_Approved = true
		}
	}
	if user_Approved && penerima_Approved {
		newMessage := Message{
			Sender: sender,
			Receiver: receiver,
			Content: content,
		}
		messages[messageCount] = newMessage
		messageCount++
		fmt.Println("Pesan berhasil dikirim.")
	} else {
		fmt.Println("Salah satu akun atau keduanya belum disetujui.")
	}
}