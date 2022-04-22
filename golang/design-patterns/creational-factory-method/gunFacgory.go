package main

func getGun(gunType string) iGun {
	switch gunType {
	case GunTypeAk47:
		return newAk47()
	case GunTypeMusket:
		return newMusket()
	default:
		return nil
	}
}
