package models

// User is the main user model.
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"index:length:50,unique"`
}

// GetUsers queries the database for all users.
func GetUsers(users *[]User) (err error) {
	if err = DB.Find(users).Error; err != nil {
		return err
	}

	return nil
}

// CreateUser creates a new user.
func CreateUser(user *User) (err error) {
	if err = DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUser creates a new user.
func UpdateUser(user *User) (err error) {
	if err = DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

// GetUser queries the database for all users.
func GetUser(user *User, id string) (err error) {
	if err = DB.First(user, id).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser queries the database for all users.
func DeleteUser(user *User, id string) (err error) {
	if err = DB.Delete(user, id).Error; err != nil {
		return err
	}

	return nil
}
