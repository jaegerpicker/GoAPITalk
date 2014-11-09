##  Let's "Go"! Get to the point already

Go addresses each of the wish list items:

- Routing is a simple as r.Get('<url_goes_here>', func_handler_goes_here) in both martini and gorilla's mux lib
- Gorm provides a simple and clean orm to use in go, mgo is probably the best mongo driver available
- In Go it's incredibly simple using the marshal and unmarshal encoding libs in the standard library
- Go routines and channels allow you to write code exactly as if it was a sync code but have it be async, it's seriously magical but completely straight forward
and not hidden
- By default go apps are statically compiled, so deployed is a simple file copy and start on the server. Everything you need is contained in the app unless you choose to have external file dependencies.
