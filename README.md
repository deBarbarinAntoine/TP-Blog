# TP-Blog

## Introduction

This TP-Blog project is an assignment done in our first semester studying development. It shows some content formatted in JSON in a golang server using ```net/http``` package.

This project was made by:
- [Alexandre](https://github.com/alexxpgb) (F1 articles, design with Figma and oral presentation);
- [Antoine](https://github.com/debarbarinantoine) (Organization, Golang backend and some HTML & CSS);
- [Marin](https://github.com/harelmarin) (Football articles and images, terms and conditions, HTML and CSS);
- [Thierry](https://github.com/thierry-mous) (mainly HTML and CSS);
- [Yulan](https://github.com/yulannn) (Esport articles, CSS and HTML).

---

## Organization

We followed a [Trello Roadmap](https://trello.com/invite/b/FL92zoB0/ATTIa251e5ce304a27c302c96f73d82cd60447C5154A/repartition-travail) to divide and coordinate our teamwork.

You can also see the [Figma map](https://www.figma.com/file/Uz7Io6pV5LruYHPN2t2KKH/maquette_siteTP?type=design&node-id=0%3A1&mode=dev).

---

## How to execute it

---

### First method (running the executable file - only for windows users) :

Download the executable ```SportPulse.exe``` clicking on **<kbd>SportPulse v1.0</kbd>** in the ```Release``` section, then run it on your computer.

---

### Second method (running the go project as it is):

Download the **<kbd>repository</kbd>** as a **<kbd>.zip</kbd>** file clicking on **<kbd>Code</kbd>** and then on **<kbd>Download ZIP</kbd>**.

Unzip it and then go to ```TP-Blog/exec``` and right click in it. Then, click on **<kbd>Open in the terminal</kbd>**.

In the terminal, write that line and press <kbd>Enter</kbd>:
```
go run ./main.go
```

Then, it should display that:
```
Server is running...
If the navigator didn't open on its own, just go to  http://localhost:8080/  on your browser.
If you want to end the server, type 'stop' here :
```
_Don't close the terminal, otherwise, the server will stop automatically._

When you open your browser at the right *url*, you can then access the website.

To stop the server, return to the terminal and type ```stop```, then press <kbd>Enter</kbd>.

A lot of information should be displayed in the terminal, because the program is currently quite verbose. Just don't worry about it.
