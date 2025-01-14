// Code generated by goctl. DO NOT EDIT.
package types

type GetUserReq struct {
	Username string `path:"username"` // user name
}

type GetUserRes struct {
	Result Result `json:"result"` //  response result
	User   User   `json:"user"`   //  info user
}

type GetUsersReq struct {
}

type GetUsersRes struct {
	Result Result `json:"result"` //  response result
	Users  []User `json:"users"`  //  list info user
}

type LoginReq struct {
	Username string `form:"username"` //  username
	Password string `form:"password"` //  password
}

type LoginRes struct {
	Result         Result `json:"result"`         //          response result
	Token          string `json:"authToken"`      //       jwttoken for api
	BaseStorageUrl string `json:"baseStorageUrl"` //  BaseStorageUrl
	UserInfo       User   `json:"user"`           //  info user
}

type RegisterReq struct {
	Role        int    `json:"role,optional"`        // Role user
	UserName    string `json:"userName,optional"`    // User name
	FullName    string `json:"fullName,optional"`    // Full name
	Email       string `json:"email,optional"`       // Email user
	Gender      int    `json:"gender,optional"`      // Gender user
	Avatar      string `json:"avatar,optional"`      // Avatar user
	PhoneNumber string `json:"phoneNumber,optional"` // Phone number of user
	CreatedAt   int64  `json:"createdAt,optional"`   // Time user created
	UpdatedAt   int64  `json:"updatedAt,optional"`   // Time user updated
}

type Result struct {
	Code    int    `json:"code"`    //	Result code: 0 is success. Otherwise, getting an error
	Message string `json:"message"` //	Result message: detail response code
}

type UpdateDeviceTokenReq struct {
	UserID int64  `header:"User-Id"`
	Token  string `json:"token"`
}

type UpdateDeviceTokenRes struct {
	Code    int    `json:"code"`    //    Result code: 0 is success. Otherwise, getting an error
	Message string `json:"message"` // Result message: detail response code
}

type User struct {
	ID          int64  `json:"id,optional"`
	Role        int    `json:"role,optional"`        // Role user
	UserID      int64  `json:"userId,optional"`      // ID of user
	UserName    string `json:"userName,optional"`    // User name
	FullName    string `json:"fullName,optional"`    // Full name
	Email       string `json:"email,optional"`       // Email user
	Gender      int    `json:"gender,optional"`      // Gender user
	Avatar      string `json:"avatar,optional"`      // Avatar user
	PhoneNumber string `json:"phoneNumber,optional"` // Phone number of user
	CreatedAt   int64  `json:"createdAt,optional"`   // Time user created
	UpdatedAt   int64  `json:"updatedAt,optional"`   // Time user updated
}
