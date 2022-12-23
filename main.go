package main

import "fmt"

func main() {
	// Contoh 1
	fmt.Println("Contoh 1, Goku")
	profile := MakeProfile("Goku")
	fmt.Println("Name", profile.Name)
	fmt.Println("Health", profile.Health)
	fmt.Println("Power", profile.Power)
	fmt.Println("Exp", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
	fmt.Println("==========================================")
	fmt.Println()

	// Contoh 2
	profile = MakeProfile("Sasuke")
	fmt.Println("Name", profile.Name)
	fmt.Println("Health", profile.Health)
	fmt.Println("Power", profile.Power)
	fmt.Println("Exp", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp = PowerUp(profile, 3)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
	fmt.Println("==========================================")
	fmt.Println()

	// Contoh 3
	profile = MakeProfile("Naruto")
	fmt.Println("Name", profile.Name)
	fmt.Println("Health", profile.Health)
	fmt.Println("Power", profile.Power)
	fmt.Println("Exp", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp = PowerUp(profile, 4)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
}

// membuat struct profile
type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(namaHero string) Profile {
	// membuat object baru dari struct Profile
	var newProfile Profile

	// jika input nama heronnya adalah Sasuke, isi variabel newProfile dengan isian berikut
	if namaHero == "Sasuke" {
		newProfile.Name = namaHero
		newProfile.Health = 200
		newProfile.Power = 100
		newProfile.Exp = 0
	} else if namaHero == "Goku" { // jika input nama heronnya adalah Goku, isi variabel newProfile dengan isian berikut
		newProfile.Name = namaHero
		newProfile.Health = 400
		newProfile.Power = 300
		newProfile.Exp = 100
	} else if namaHero == "Naruto" { // jika input nama heronnya adalah Naruto, isi variabel newProfile dengan isian berikut
		newProfile.Name = namaHero
		newProfile.Health = 150
		newProfile.Power = 200
		newProfile.Exp = 50
	}

	// kirim variabel newProfile sebagai nilai return
	return newProfile
}

func PowerUp(profile Profile, multiplier int) Profile {
	// membuat object baru dari cetakan struct Profile, definisikan isinya, lalu kirim sebagai nilai kembalian
	return Profile{
		// pada bagian nama, isi sesuai dengan object profile yang diinputkan pada parameter
		Name: profile.Name,
		// mengisi nilai property-property berikut dengan perhitungan antara property-property dari profile dan multiplier yang diinputkan pada parameter
		Health: profile.Health + (profile.Health * multiplier),
		Power:  profile.Power + (profile.Power * multiplier),
		Exp:    profile.Exp + (profile.Exp * multiplier),
	}
}
