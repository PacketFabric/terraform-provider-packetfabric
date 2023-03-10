package packetfabric

import (
	"errors"
	"fmt"
)

const UsersURI = "/v2/users"
const PasswordUpdateURI = "/v2/users/%s/password"

// This struct represents a User
// https://docs.packetfabric.com/api/v2/swagger/#/Users/user_post
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Timezone  string `json:"timezone"`
	Group     string `json:"group"`
}

type UserUpdate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Login     string `json:"login"`
	Timezone  string `json:"timezone"`
	Group     string `json:"group"`
}

// This struct represents a User response
// https://docs.packetfabric.com/api/v2/swagger/#/Users/user_post
type UserResponse struct {
	UUID                string `json:"uuid"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Phone               string `json:"phone"`
	Timezone            string `json:"timezone"`
	Email               string `json:"email"`
	Login               string `json:"login"`
	Group               string `json:"group"`
	MFAEnabled          bool   `json:"mfa_enabled"`
	TimeLastLogin       string `json:"time_last_login,omitempty"`
	ResetPasswordBefore string `json:"reset_password_before,omitempty"`
}

// This struct represents a User delete response
// https://docs.packetfabric.com/api/v2/swagger/#/Users/user_delete_by_login
type UserDelResp struct {
	Message string `json:"message"`
}

// This function represents the Action to create a new User
// https://docs.packetfabric.com/api/v2/swagger/#/Users/user_post
func (c *PFClient) CreateUsers(user User) (*UserResponse, error) {
	resp := &UserResponse{}
	_, err := c.sendRequest(UsersURI, postMethod, user, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Retrieve an existing User by ID
// https://docs.packetfabric.com/api/v2/swagger/#/Users/user_get_by_login
func (c *PFClient) ReadUsers(userID string) (*UserResponse, error) {
	formatedURI := fmt.Sprintf("%s/%s", UsersURI, userID)
	resp := &UserResponse{}
	_, err := c.sendRequest(formatedURI, getMethod, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to update an existing User
// https://docs.packetfabric.com/api/v2/swagger/#/Users/user_patch
func (c *PFClient) UpdateUser(user UserUpdate, userID string) (*UserResponse, error) {
	formatedURI := fmt.Sprintf("%s/%s", UsersURI, userID)
	resp := &UserResponse{}
	_, err := c.sendRequest(formatedURI, patchMethod, user, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to update user's password
// https://docs.packetfabric.com/api/v2/swagger/#/Users/user_update_password
func (c *PFClient) UserPasswordUpdate(userID string, oldPassword string, newPassword string) (*UserDelResp, error) {
	resp := &UserDelResp{}
	formatedURI := fmt.Sprintf(PasswordUpdateURI, userID)
	type UserPasswodUpdate struct {
		NewPassword string `json:"new_password"`
		OldPassword string `json:"old_password"`
	}
	_, err := c.sendRequest(formatedURI, patchMethod, UserPasswodUpdate{NewPassword: newPassword, OldPassword: oldPassword}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// This function represents the Action to Delete an existing User
// https://docs.packetfabric.com/api/v2/swagger/#/Users/user_delete_by_login
func (c *PFClient) DeleteUsers(userID string) (*UserDelResp, error) {
	if userID == "" {
		return nil, errors.New(errorMsg)
	}
	formatedURI := fmt.Sprintf("%s/%s", UsersURI, userID)
	expectedResp := &UserDelResp{}
	_, err := c.sendRequest(formatedURI, deleteMethod, nil, expectedResp)
	if err != nil {
		return expectedResp, err
	}
	return expectedResp, nil
}
