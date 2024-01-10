# Templates Documentation
## Showing structure of data sent to every template

The variable sent to any template is ```data``` and is always composed of at least one structure.

Here is an example of the template's execution:
```go
err := tmpl["index"].ExecuteTemplate(w, "base", data)
```

---

## Templates:

- [index](#indexgohtml)
- [category](#categorygohtml)
- [article](#articlegohtml)
- [search](#searchgohtml)
- [login](#logingohtml)
- [create user](#createusergohtml)
- [modify user](#modifyusergohtml)
- [admin](#admingohtml)
- [add article](#addarticlegohtml)
- [modify article](#modifyarticlegohtml)
- [delete article](#deletearticlegohtml)
- [about](#aboutgohtml)
- [error 404](#error404gohtml)

---


### index.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Articles []Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
}
```
[↑ Return to table of content](#templates)

---

### category.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    MainArticle Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
    Category []Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
}
```
[↑ Return to table of content](#templates)


---

### article.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Article Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
}
```
[↑ Return to table of content](#templates)


---

### search.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Articles []Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
    Search  string  // the word searched for
    Message string  // message if the search doesn't match any content: <div class="message">There is no article matching your research!</div>
}
```
[↑ Return to table of content](#templates)


---

### login.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Message string // message if there is a problem logging (username or password): <div class="message">Wrong username or password!</div>
}                  // also contains a message when redirected from restricted website area without login.
```
[↑ Return to table of content](#templates)


---

### createuser.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Message string // message if there is a problem signing up (username or password): <div class="message">Username already used!</div>
}
```
[↑ Return to table of content](#templates)


---

### modifyuser.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Message string // message if there is a problem modifying user info (username or password): <div class="message">Invalid data!</div>
}
```
[↑ Return to table of content](#templates)


---

### admin.gohtml
```go
data := struct {
    Base BaseData {
        Title      string
        StaticPath string
    }
    Articles []Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
    User User {
        Name        string
        Password    string
    }
}
```
[↑ Return to table of content](#templates)


---

### addarticle.gohtml
```go
data := struct {
    Base       BaseData {
        Title      string
        StaticPath string
    }
    User       User {
        Name        string
        Password    string
    }
    Categories []string     // Containing all category titles: []string{"Formule 1", "Esport", "Football"}
    Article    Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
}
```
[↑ Return to table of content](#templates)


---

### modifyarticle.gohtml
```go
data := struct {
    Base    BaseData {
        Title      string
        StaticPath string
    }
    Article Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
}
```
[↑ Return to table of content](#templates)


---

### deletearticle.gohtml
```go
data := struct {
    Base    BaseData {
        Title      string
        StaticPath string
    }
    Article Article {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      string
    }
    Message string  // message asking for confirmation: <div class="message">Do you really want to delete that article ?</div>
}
```
[↑ Return to table of content](#templates)


---

### about.gohtml
```go
data := struct {
    Base BaseData {
        Title      string
        StaticPath string
    }
}
```
[↑ Return to table of content](#templates)


---

### error404.gohtml
```go
data := struct {
    Base BaseData {
        Title      string
        StaticPath string
    }
}
```
[↑ Return to table of content](#templates)

