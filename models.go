package TPBlog

type BaseData struct {
	Title      string
	StaticPath string
	Line       string
}

type Session struct {
	isOpen bool
	MyUser User
}

type User struct {
	Name     string
	Password string
}
