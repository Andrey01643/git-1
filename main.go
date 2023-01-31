package main

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//ПАРСЕР

//func main() {
//	geziyor.NewGeziyor(&geziyor.Options{
//		StartURLs: []string{"https://www.fabrikant.ru/trades/procedure/search/#"},
//		ParseFunc: parseFabrikant,
//		Exporters: []export.Exporter{&export.JSON{}},
//	}).Start()
//}
//
//func parseFabrikant(g *geziyor.Geziyor, r *client.Response) {
//	r.HTMLDoc.Find("div.innerGrid").Each(func(i int, s *goquery.Selection) {
//		var sessions = strings.Split(s.Find(".marketplace-unit__info__name").Text(), " \n ")
//		sessions = sessions[:len(sessions)-1]
//
//		//for i := 0; i < len(sessions); i++ {
//		//	sessions[i] = strings.Trim(sessions[i], "\n ")
//		//}
//
//		var description string
//
//		if href, ok := s.Find("a.text").Attr("href"); ok {
//			g.Get(r.JoinURL(href), func(_g *geziyor.Geziyor, _r *client.Response) {
//				description = _r.HTMLDoc.Find("span.popuptext_icons").Text()
//
//				description = strings.ReplaceAll(description, "\t", "")
//				description = strings.ReplaceAll(description, "\n", "")
//				description = strings.TrimSpace(description)
//
//				g.Exports <- map[string]interface{}{
//					"title": strings.TrimSpace(s.Find("a.text").Text()),
//					//"subtitle":    strings.TrimSpace(s.Find("div.marketplace-unit__info__name").Text()),
//					"sessions":    sessions,
//					"description": description,
//				}
//			})
//		}
//	})
//}
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Требуется реализовать Rest API с двумя endpoint'ами, с применением пакета net/http или gin. "Content-Type" и "Accept" - "application/json".

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"sync"
//)
//
//// Book представляет собой книгу с идентификатором, названием, автором и годом издания.
//type Book struct {
//	ID            int    `json:"id"`
//	Title         string `json:"title"`
//	Author        string `json:"author"`
//	PublisherYear int    `json:"publisher_year"`
//}
//
//// Repository это хранилище книг в оперативной памяти.
//type Repository struct {
//	mu    sync.RWMutex
//	books []Book
//}
//
//// Add добавляет книгу в репозиторий
//func (r *Repository) Add(book Book) {
//	r.mu.Lock()
//	defer r.mu.Unlock()
//	r.books = append(r.books, book)
//}
//
//// Delete удаляет книгу из репозитория по ID
//func (r *Repository) Delete(id int) error {
//	r.mu.Lock()
//	defer r.mu.Unlock()
//	for i, book := range r.books {
//		if book.ID == id {
//			r.books = append(r.books[:i], r.books[i+1:]...)
//			return nil
//		}
//	}
//	return fmt.Errorf("book with ID %d not found", id)
//}
//
//// All возвращает все книги в репозитории.
//func (r *Repository) All() []Book {
//	r.mu.RLock()
//	defer r.mu.RUnlock()
//	return r.books
//}
//
//// BookService предоставляет доступ к хранилищу книг
//type BookService interface {
//	Add(Book)
//	Delete(int) error
//	All() []Book
//}
//
//var (
//	service BookService = &Repository{}
//	limiter             = NewLimiter(10) // 10 запросов в секунду
//)
//
//func main() {
//	service.Add(Book{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", PublisherYear: 1925})
//	service.Add(Book{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", PublisherYear: 1960})
//	service.Add(Book{ID: 3, Title: "1984", Author: "George Orwell", PublisherYear: 1949})
//
//	http.HandleFunc("/books", limiter.LimitFunc(booksHandler))
//	http.HandleFunc("/books/delete", limiter.LimitFunc(deleteBookHandler))
//	http.ListenAndServe(":8080", nil)
//}
//
//func booksHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodGet {
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//
//		return
//	}
//	books := service.All()
//	jsonBooks, err := json.Marshal(books)
//	if err != nil {
//		http.Error(w, "Failed to marshal books to JSON", http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(jsonBooks)
//}
//
//func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodDelete {
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//}
//
////NewLimiter создает новый limiter(ограничитель) с заданной ставкой
//func NewLimiter(rate int) *Limiter {
//	return &Limiter{
//		rate:   rate,
//		bucket: make(chan struct{}, rate),
//	}
//}
//
//// Limiter простой ограничитель limiter
//type Limiter struct {
//	rate   int
//	bucket chan struct{}
//}
//
//// LimitFunc ограничивает скорость запросов до заданной скорости
//func (l *Limiter) LimitFunc(fn http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		l.bucket <- struct{}{}
//		defer func() {
//			<-l.bucket
//		}()
//		fn(w, r)
//	}
//}
//
//// Limit ограничивает скорость запросов до заданной скорости
//func (l *Limiter) Limit(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		l.bucket <- struct{}{}
//		defer func() {
//			<-l.bucket
//		}()
//		next.ServeHTTP(w, r)
//	})
//}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
