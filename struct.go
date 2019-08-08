package main

import "fmt"

type Books struct {
    title string
    auther string
    subject string
    book_id int
}

func output_book(book Books){
    fmt.Println(book.title)
    fmt.Println(book.auther)
    fmt.Println(book.subject)
    fmt.Println(book.book_id)
}

func output_book2(book *Books){
    fmt.Println(book.title)
    fmt.Println(book.auther)
    fmt.Println(book.subject)
    fmt.Println(book.book_id)
}

func main(){
    book1 := Books {"Go language", "Jeremy", "Go language course", 33548970}
    book2 := Books {title: "Flower all", auther: "xiaoxie", subject: "Flower all over the world", book_id: 43219870}
    fmt.Println(book1)
    fmt.Println(book2)

    var book3 Books
    book3.title = "Machine Learning"
    book3.auther = "who are you"
    book3.subject = "ML course"
    book3.book_id = 5499832
    fmt.Println(book3.subject)

    output_book(book2)
    output_book2(&book1)
}
