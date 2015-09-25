## tree 
`import "github.com/go-kami/tree"`

tree is the routing tree ripped from [HttpRouter](https://github.com/julienschmidt/httprouter). Instead of storing a HTTP handler, this tree stores an `interface{}`, so you can put whatever you'd like inside of it.