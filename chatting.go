package main

import "fmt"

const (
	MAX_USERS    = 100
	MAX_MESSAGES = 200
)

type User struct {
	ID       int
	Name     string
	Username string
	Password string
	Approved bool
}

type Message struct {
	Sender   string
	Receiver string
	Content  string
}

type Group struct {
	Name    string
	Members []string
}

var users [MAX_USERS]User
var userCount int
var messages [MAX_MESSAGES]Message
var messageCount int
var groups []Group

func main() {
	var choice int
	for {
		fmt.Println()
		fmt.Println("-------- Menu Utama --------")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Admin")
		fmt.Println("4. Keluar")
		fmt.Println("----------------------------")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, username, password string
			fmt.Print("Masukkan Nama: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan username: ")
			fmt.Scan(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scan(&password)
			register(name, username, password)
		case 2:
			login()
		case 3:
			adminMenu()
		case 4:
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Opsi tidak valid, silahkan pilih lagi.")
		}
	}
}

func userMenu(username string) {
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

		switch subChoice {
		case 1:
			var receiver, content string
			fmt.Print("Masukan username tujuan: ")
			fmt.Scan(&receiver)
			fmt.Print("Masukan pesan: ")
			fmt.Scan(&content)
			sendPrivateMessage(username, receiver, content)
		case 2:
			var groupName string
			fmt.Print("Masukkan nama grup: ")
			fmt.Scan(&groupName)
			createGroup(groupName, username)
		case 3:
			var groupName, memberUsername string
			fmt.Print("Masukkan nama grup: ")
			fmt.Scan(&groupName)
			fmt.Print("Masukkan username anggota baru: ")
			fmt.Scan(&memberUsername)
			addMemberToGroup(groupName, memberUsername)
		case 4:
			var groupName, content string
			fmt.Print("Masukkan nama grup: ")
			fmt.Scan(&groupName)
			fmt.Print("Masukkan pesan: ")
			fmt.Scan(&content)
			sendGroupMessage(groupName, username, content)
		case 5:
			var groupName string
			fmt.Print("Masukkan nama grup: ")
			fmt.Scan(&groupName)
			viewGroupMembers(groupName)
		case 6:
			return
		default:
			fmt.Println("Opsi tidak valid, silakan pilih lagi.")
		}
	}
}

func adminMenu() {
	var adminChoice int
	for {
		fmt.Println()
		fmt.Println("-------- Menu Admin --------")
		fmt.Println("1. Setujui Akun")
		fmt.Println("2. Tolak Akun")
		fmt.Println("3. Cetak Daftar Akun yang Approved")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Println("----------------------------")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&adminChoice)

		switch adminChoice {
		case 1:
			var username string
			fmt.Print("Masukan username yang ingin di setujui: ")
			fmt.Scan(&username)
			approved(username)
		case 2:
			var username string
			fmt.Print("Masukan username yang ingin di tolak: ")
			fmt.Scan(&username)
			rejected(username)
		case 3:
			printUsers()
		case 4:
			return
		default:
			fmt.Print("Opsi tidak valid, silahkan pilih lagi.")
		}
	}
}

//Fungsi untuk melakukan registrasi user (Hellmina Enjelina Fitri)
func register(name, username, password string) {
	for i := 0; i < userCount; i++ {
		if users[i].Username == username {
			fmt.Printf("Username %s sudah digunakan. Silakan buat dengan username lain.\n", username)
			return
		}
	}

	new_User := User{
		ID:       userCount + 1,
		Name:     name,
		Username: username,
		Password: password,
		Approved: false,
	}

	users[userCount] = new_User
	userCount++
	fmt.Printf("Registrasi akun %s berhasil silahkan tunggu persetujuan dari Admin.\n", username)
}

//Fungsi untuk melakukan login user (Hellmina Enjelina Fitri)
func login() {
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

	if validLogin {
		userMenu(username)
	} else {
		fmt.Println("Username tidak valid, Password salah, atau Akun belum Approved.")
	}
}

//Fungsi untuk menyetujui akun user (Hellmina Enjelina Fitri)
func approved(username string) {
	for i := 0; i < userCount; i++ {
		if users[i].Username == username {
			users[i].Approved = true
			fmt.Printf("Akun %s telah disetujui oleh Admin.\n", users[i].Username)
			return
		}
	}
	fmt.Println("Username tidak ditemukan")
}

//Fungsi untuk menolak akun user (Hellmina Enjelina Fitri)
func rejected(username string) {
	for i := 0; i < userCount; i++ {
		if users[i].Username == username {
			fmt.Printf("Akun %s ditolak.\n", users[i].Username)
			for j := i; j < userCount-1; j++ {
				users[j] = users[j+1]
			}
			userCount--
			return
		}
	}
	fmt.Println("Username tidak ditemukan")
}

//Fungsi untuk menampilkan daftar user (Hellmina Enjelina Fitri)
func printUsers() {
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
func sendPrivateMessage(sender, receiver, content string) {
	userApproved, receiverApproved := false, false
	for i := 0; i < userCount; i++ {
		if users[i].Username == sender && users[i].Approved {
			userApproved = true
		}
		if users[i].Username == receiver && users[i].Approved {
			receiverApproved = true
		}
	}
	if userApproved && receiverApproved {
		newMessage := Message{
			Sender:   sender,
			Receiver: receiver,
			Content:  content,
		}
		messages[messageCount] = newMessage
		messageCount++
		fmt.Println("Pesan berhasil dikirim.")
	} else {
		fmt.Println("Salah satu akun atau keduanya belum disetujui.")
	}
}

//Fungsi untuk membuat group (Muhammad Haiqal Patria Iskandar)
func createGroup(groupName, creatorUsername string) {
	for _, group := range groups {
		if group.Name == groupName {
			fmt.Printf("Group %s already exists.\n", groupName)
			return
		}
	}

	newGroup := Group{
		Name:    groupName,
		Members: []string{creatorUsername}, // Creator is automatically a member of the group
	}
	groups = append(groups, newGroup)
	fmt.Printf("Group %s created successfully.\n", groupName)
}

//Fungsi untuk menambahkan anggota group (Muhammad Haiqal Patria Iskandar)
func addMemberToGroup(groupName, memberUsername string) {
	for i, group := range groups {
		if group.Name == groupName {
			for _, member := range group.Members {
				if member == memberUsername {
					fmt.Printf("User %s is already a member of %s.\n", memberUsername, groupName)
					return
				}
			}
			groups[i].Members = append(groups[i].Members, memberUsername)
			fmt.Printf("User %s added to group %s successfully.\n", memberUsername, groupName)
			return
		}
	}
	fmt.Printf("Group %s not found.\n", groupName)
}

//Fungsi untuk mengirim pesan ke group (Muhammad Haiqal Patria Iskandar)
func sendGroupMessage(groupName, sender, content string) {
	for _, group := range groups {
		if group.Name == groupName {
			found := false
			for _, member := range group.Members {
				if member == sender {
					found = true
					break
				}
			}
			if found {
				newMessage := Message{
					Sender:   sender,
					Receiver: groupName, // Group name as receiver
					Content:  content,
				}
				messages[messageCount] = newMessage
				messageCount++
				fmt.Printf("Message sent to group %s successfully.\n", groupName)
				return
			} else {
				fmt.Printf("Sender %s is not a member of group %s.\n", sender, groupName)
				return
			}
		}
	}
	fmt.Printf("Group %s not found.\n", groupName)
}

//Fungsi untuk menampilkan anggota di dalam group (Muhammad Haiqal Patria Iskandar)
func viewGroupMembers(groupName string) {
	for _, group := range groups {
		if group.Name == groupName {
			fmt.Printf("Members of %s:\n", groupName)
			for _, member := range group.Members {
				fmt.Println(member)
			}
			return
		}
	}
	fmt.Printf("Group %s not found.\n", groupName)
}