package main

// Reference: https://www.integralist.co.uk/posts/understanding-golangs-func-type/
func main() {
	// 版本一和版本二
	/*
		db := database{
			"foo": "football",
			"bar": "basketball",
		}
		log.Fatal(http.ListenAndServe(":3030", db))
	*/

	// 版本三
	/*
		db := database{
			"foo": "foo value",
			"bar": "bar value",
			"baz": "baz value",
		}

		mux := http.NewServeMux()
		mux.Handle("/foo", http.HandlerFunc(db.foo))
		mux.Handle("/bar", http.HandlerFunc(db.bar))

		mux.HandleFunc("baz", db.baz)
		log.Fatal(http.ListenAndServe(":3030", mux))
	*/

	// 版本四
	/*
		db := database{
			"foo": "foo value",
			"bar": "bar value",
			"baz": "baz value",
		}

		http.HandleFunc("/foo", db.foo)
		http.HandleFunc("/bar", db.bar)
		http.HandleFunc("/baz", db.baz)

		log.Fatal(http.ListenAndServe(":3030", nil))
	*/
}

type database map[string]string

// 版本一：响影所有的 path response
/*
func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s", r.Method, r.RequestURI)
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s", k, v)
	}
}
*/

// 版本二：针对 request 的 path 做分支处理，相当于自定义路由
/*
func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/foo":
		fmt.Fprintf(w, "this is foo")
	case "/bar":
		fmt.Fprintf(w, "this is bar")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "no route found, this is default")
	}
}
*/

/*
// 版本三/四共用：使用 func(w http.ResponseWriter, r *http.Request) 作为处理 http 请求的函数签名
func (db database) foo(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "this is from foo path", db["foo"])
}
func (db database) bar(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "this is from foo path", db["bar"])
}
func (db database) baz(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "this is from foo path", db["baz"])
}
*/
