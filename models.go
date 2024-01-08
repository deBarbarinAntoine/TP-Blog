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

type Article struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Date     string `json:"date"`
	Content  string `json:"content"`
}
