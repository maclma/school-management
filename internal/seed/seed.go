package seed

import "gorm.io/gorm"

func Run(db *gorm.DB) error {
	if err := SeedRoles(db); err != nil {
		return err
	}

	if err := SeedAdminUser(db); err != nil {
		return err
	}

	if err := SeedClasses(db); err != nil {
		return err
	}

	return nil
}
