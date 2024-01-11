package TPBlog

// BaseData stores all basic data used in the base.gohtml template.
type BaseData struct {
	Title      string
	StaticPath string
}

// Session stores all User info.
type Session struct {
	IsOpen bool
	MyUser User
}

// User login info.
type User struct {
	Name     string
	Password string
}

// Article stores all info and content used in the website.
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
