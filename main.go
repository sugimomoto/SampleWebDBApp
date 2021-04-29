package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/books", BooksPage)
	http.ListenAndServe(":8081", nil)

}

func BooksPage(writer http.ResponseWriter, request *http.Request) {

	books, err := GetBooks()

	fmt.Println(books)

	if err != nil {
		panic(err)
	}

	t, err := template.ParseFiles("books.html")
	if err != nil {
		panic(err)
	}

	t.Execute(writer, books)
}

func GetBooks() (books Root, err error) {
	url := "https://oreillydemoapi.azurewebsites.net/api.rsc/OReillyBooks/?$top=20"
	authHeaderName := "x-cdata-authtoken"
	authHeaderValue := "7y3E6q4b6V1v9f0D2m9j"

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(authHeaderName, authHeaderValue)

	client := new(http.Client)
	resp, err := client.Do(req)

	// URLがnilだったり、Timeoutが発生した場合にエラーを返す模様。
	// サーバーからのレスポンスとなる 401 Unauthroized Error などはResponseをチェックする。
	// サーバーとの疎通が開始する前の動作のよう。
	if err != nil {
		fmt.Println("Error Request:", err)
		return Root{}, err
	}
	// resp.Bodyはクローズすること。クローズしないとTCPコネクションを開きっぱなしになる。
	defer resp.Body.Close()

	// 200 OK 以外の場合はエラーメッセージを表示して終了
	if resp.StatusCode != 200 {
		fmt.Println("Error Response:", resp.Status)
		return Root{}, err
	}

	// とりあえずResponsの構造体を全部出力
	fmt.Printf("%-v", resp)

	// Response Body を読み取り
	body, _ := io.ReadAll(resp.Body)

	// JSONを構造体にエンコード
	var Books Root
	json.Unmarshal(body, &Books)

	return Books, nil
}

type Root struct {
	Value []Books `json:"value"`
}

type Books struct {
	RowId       int
	ImageUrl    string
	ISBN        string
	Price       string
	PublishDate string
	Title       string
	URL         string
}
