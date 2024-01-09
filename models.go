package TPBlog

type BaseData struct {
	Title      string
	StaticPath string
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
	Id           int    `json:"id"`
	Category     string `json:"category"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	Date         string `json:"date"`
	BigImg       string `json:"big_img"`
	SmallImg     string `json:"small_img"`
	Introduction string `json:"introduction"`
	Content      string `json:"content"`
}
