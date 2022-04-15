package books

import (
	"books/internal/auth"
	"github.com/gin-gonic/gin"
	"strings"
)

type Api struct {
	Service   BookServiceInterface
	SecretKey string
}

func (b *Api) checkToken(ctx *gin.Context) int {
	rawToken := ctx.GetHeader(`Authorization`)
	token := strings.Split(rawToken, " ")
	_, exc := auth.GetUserIdByJwtToken(token[1], b.SecretKey)
	if exc != nil {
		ctx.JSON(401, gin.H{`err`: exc.Error()})
		return 1
	}
	return 0
}

//CreateNewBook create new book endpoint
func (b *Api) CreateNewBook(ctx *gin.Context) {
	var schema CreateBookSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		book, exc := b.Service.CreateBook(schema.BookAuthor, schema.BookTitle, schema.YearOfRelease, schema.CoverUrl, schema.Description)
		if exc != nil {
			ctx.JSON(500, gin.H{`err`: exc.Error()})
			return
		}
		ctx.JSON(201, book)
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
func (b *Api) DeleteBook(c *gin.Context) {
	var schema BookIdSchema
	if err := c.ShouldBindJSON(&schema); err == nil {
		boolRes, exc := b.Service.DeleteBook(schema.BookId)
		if boolRes || exc == nil {
			c.JSON(200, gin.H{`success`: true})
		}
	} else {
		c.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
func (b *Api) GetBookByID(ctx *gin.Context) {
	var schema BookIdSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		book, exc := b.Service.GetBookById(schema.BookId)
		if exc != nil {
			ctx.JSON(500, gin.H{`err`: exc.Error()})
			return
		}
		ctx.JSON(200, book)
	} else {
		ctx.JSON(400, err)
		return
	}
}
func (b *Api) GetAllBooks(ctx *gin.Context) {
	books, exc := b.Service.GetAllBooks()
	if exc != nil {
		ctx.JSON(500, gin.H{`err`: exc.Error()})
		return
	}
	ctx.JSON(200, books)
}
func (b *Api) UpdateAuthor(ctx *gin.Context) {
	var schema UpdateAuthorSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		result, exc := b.Service.UpdateAuthor(schema.BookId, schema.NewAuthor)
		if result && exc == nil {
			ctx.JSON(200, gin.H{`success`: true})
		}
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
func (b *Api) UpdateYearOfRelease(ctx *gin.Context) {
	var schema UpdateYearOfReleaseSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		result, exc := b.Service.UpdateYearOfRelease(schema.BookId, schema.NewYearOfRelease)
		if result || exc == nil {
			ctx.JSON(200, gin.H{`success`: true})
		} else {
			ctx.JSON(400, gin.H{`err`: err.Error()})
			return
		}
	}
}
func (b *Api) UpdateDescription(ctx *gin.Context) {
	var schema UpdateDescriptionSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		result, exc := b.Service.UpdateDescription(schema.BookId, schema.NewDescription)
		if result || exc == nil {
			ctx.JSON(200, gin.H{`success`: true})
		} else {
			ctx.JSON(400, gin.H{`err`: err.Error()})
			return
		}
	}
}
func (b *Api) UpdateCover(ctx *gin.Context) {
	var schema UpdateCoverSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		result, exc := b.Service.UpdateCover(schema.BookId, schema.NewCoverUrl)
		if result || exc == nil {
			ctx.JSON(200, gin.H{`success`: true})
		} else {
			ctx.JSON(500, gin.H{`err`: exc.Error()})
			return
		}
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
func (b *Api) UpdateTitle(ctx *gin.Context) {
	var schema UpdateBookTitleSchema
	if err := ctx.ShouldBindJSON(&schema); err == nil {
		b.checkToken(ctx)
		result, exc := b.Service.UpdateBookTitle(schema.BookId, schema.NewBookTitle)
		if result || exc == nil {
			ctx.JSON(200, gin.H{`success`: true})
		} else {
			ctx.JSON(500, gin.H{`err`: exc.Error()})
			return
		}
	} else {
		ctx.JSON(400, gin.H{`err`: err.Error()})
		return
	}
}
func (b *Api) GetRandomBook(c *gin.Context) {
	book, exc := b.Service.GetRandomBook()
	if exc != nil {
		c.JSON(500, gin.H{`err`: exc.Error()})
		return
	}
	c.JSON(200, book)
}
func (b *Api) GetBookByAuthor(c *gin.Context) {
	var schema GetAllBooksByAuthorSchema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(400, gin.H{`err`: err.Error()})
		return
	}
	result, exc := b.Service.GetAllBooksByAuthor(schema.Author)
	if exc != nil {
		c.JSON(500, gin.H{`err`: exc.Error()})
		return
	}
	c.JSON(200, result)
}
