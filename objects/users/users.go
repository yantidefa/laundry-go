package users

type User struct {
	Id        int    `form:"id" json:"id"`
	Nama_user string `form:"nama_user" json:"nama_user"`
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	Id_outlet int    `form:"id_outlet" json:"id_outlet"`
	// Role      string `form:"role" json:"role"`
}
