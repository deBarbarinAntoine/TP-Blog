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

## Header links

- Formule 1: ``/category?category=Formule 1``
- Esport: ``/category?category=Esport``
- Foot: ``/category?category=Football``
- Login: ``/login``
- Sign up: ``/adduser``
- Search: ``/search?q=<search>``
- Logout: ``/logout``
- Admin: ``/admin``

---

## Footer links

- About: ``/about``
- Terms&Conditions: ``/about#<legal>``

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
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```

- Articles: ``/article?article=<article-id>``

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
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```
- Articles: ``/article?article=<article-id>``

[↑ Return to table of content](#templates)


---

### article.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Article ArticleHTML {
        Id           int
        Category     string
        Title        string
        Author       string
        Date         string
        BigImg       string
        SmallImg     string
        Introduction string
        Content      template.HTML
    }
    Recommended []Article {
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
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
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
    Search  string         // the word searched for
    Message template.HTML  // message if the search doesn't match any content: <div class="message">There is no article matching your research!</div> 
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```
- Articles: ``/article?article=<article-id>``

[↑ Return to table of content](#templates)


---

### login.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Message template.HTML // message if there is a problem logging (username or password): <div class="message">Wrong username or password!</div>
    Session Session {     // also contains a message when redirected from restricted website area without login.
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```

- Submit: ``/login/treatment``
- Register: ``/createuser``

[↑ Return to table of content](#templates)


---

### createuser.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Message template.HTML // message if there is a problem signing up (username or password): <div class="message">Username already used!</div>
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```

- Submit: ``/createuser/treatment``
- Login: ``/login``

[↑ Return to table of content](#templates)


---

### modifyuser.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Message template.HTML // message if there is a problem modifying user info (username or password): <div class="message">Invalid data!</div>
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```

- Submit: ``/createuser/treatment``
- Cancel: ``/admin``

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
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```

- Add article: ``/addarticle``
- Article: ``/article?article=<article-id>``
- Modify article: ``/modifyarticle?article=<article-id>``
- Delete article: ``/deletearticle?article=<article-id>``

[↑ Return to table of content](#templates)


---

### addarticle.gohtml
```go
data := struct {
    Base       BaseData {
        Title      string
        StaticPath string
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
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```

- Submit: ``/addarticle/treatment``
- Cancel: ``/admin``

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
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```

- Submit: ``/modifyarticle/treatment?article=<article-id>``
- Cancel: ``/admin``

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
    Message template.HTML  // message asking for confirmation: <div class="message">Do you really want to delete that article ?</div>
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```

- Submit: ``/deletearticle/treatment?article=<article-id>``
- Cancel: ``/admin``

[↑ Return to table of content](#templates)


---

### about.gohtml
```go
data := struct {
    Base BaseData {
        Title      string
        StaticPath string
    }
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
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
    Session Session {
        IsOpen bool
        MyUser User {
            Name string
            Password string
        }
    }
}
```
[↑ Return to table of content](#templates)

