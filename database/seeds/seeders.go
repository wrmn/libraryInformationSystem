package seeds

func SeedAll() {
	ddcSeed()
	guestSeed()
	userSeed()

	assetRecordSeed()
	bookSeeder()

	borrowSeeder()
	employeeSeeder()
}
