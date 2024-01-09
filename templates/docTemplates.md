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

### category.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
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

### login.gohtml
```go
data := struct {
    Base     BaseData {
        Title      string
        StaticPath string
    }
    Message string // message if there is a problem logging (username or password): <div class="message">Wrong username or password!</div>
}
```

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

### admin.gohtml
```go
data := struct {
    Base BaseData {
        Title      string
        StaticPath string
    }
    User User {
        Name        string
        Password    string
    }
}
```

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

### about.gohtml
```go
data := struct {
    Base BaseData {
        Title      string
        StaticPath string
    }
}
```

### error404.gohtml
```go
data := struct {
    Base BaseData {
        Title      string
        StaticPath string
    }
}
```
