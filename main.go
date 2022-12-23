package main

import "fmt"

func main() {
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

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(namaHero string) Profile {
	var newProfile Profile

	if namaHero == "Sasuke" {
		newProfile.Name = namaHero
		newProfile.Health = 200
		newProfile.Power = 100
		newProfile.Exp = 0

	} else if namaHero == "Goku" {
		newProfile.Name = namaHero
		newProfile.Health = 400
		newProfile.Power = 300
		newProfile.Exp = 100
	} else if namaHero == "Naruto" {
		newProfile.Name = namaHero
		newProfile.Health = 150
		newProfile.Power = 200
		newProfile.Exp = 50

	}

	return newProfile
}

func PowerUp(profile Profile, multiplier int) Profile {
	return Profile{
		Name:   profile.Name,
		Health: profile.Health + (profile.Health * multiplier),
		Power:  profile.Power + (profile.Power * multiplier),
		Exp:    profile.Exp + (profile.Exp * multiplier),
	}
}
