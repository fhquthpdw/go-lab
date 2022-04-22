package main

const GunTypeMusket = "musket"

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun{
			name:  "musket",
			power: 1,
		},
	}
}
