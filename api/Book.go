package api

//Book Resource
type Book struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Pages  int    `json:"pages" binding:"required,min=4,max=500"`
}
