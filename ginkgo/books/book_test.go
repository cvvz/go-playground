package books_test

import (
	"time"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"example/books"
)

var i bool

var _ = ginkgo.Describe("Book", func() {
	var foxInSocks, lesMis *books.Book

	ginkgo.BeforeEach(func() {
		lesMis = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		foxInSocks = &books.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	ginkgo.Describe("Categorizing books", func() {
		ginkgo.Context("with more than 300 pages", func() {
			ginkgo.BeforeEach(func() {
			})
			ginkgo.It("should be a novel", func() {
				Expect(lesMis.Category()).To(Equal(books.CategoryNovel))
			})
		})

		ginkgo.Context("with fewer than 300 pages", func() {
			ginkgo.It("should be a short story", func() {
				Expect(foxInSocks.Category()).To(Equal(books.CategoryShortStory))
			})
		})
	})

	ginkgo.It("can extract the author's last name", func() {
		book := &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		Expect(book.AuthorLastName()).To(Equal("Hugo"))
	})
})

var _ = ginkgo.It("can extract the author's first name", ginkgo.Serial, func() {
	book := &books.Book{
		Title:  "Les Miserables",
		Author: "Victor Hugo",
		Pages:  2783,
	}
	ginkgo.By("test output")

	Expect(book.AuthorFirstName()).To(Equal("Victor"))
})

var _ = ginkgo.Describe("Books", func() {
	ginkgo.It("can extract the author's last name", func() {
		book := &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		Expect(book.AuthorLastName()).To(Equal("Hugo"))
	})

	ginkgo.It("can fetch a summary of the book from the library service", func(ctx ginkgo.SpecContext) {
		flakehelper()
	}, ginkgo.SpecTimeout(time.Second))
})

var _ = ginkgo.Describe("book", func() {
	var book *books.Book

	ginkgo.BeforeEach(func() {

		book = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}
	})

	ginkgo.DescribeTable("Extracting the author's first and last name",
		func(author string, isValid bool, firstName string, lastName string) {
			book.Author = author
			Expect(book.AuthorFirstName()).To(Equal(firstName))
			Expect(book.AuthorLastName()).To(Equal(lastName))
		},
		ginkgo.Entry("When author has both names", "Victor Hugo", true, "Victor", "Hugo"),
		ginkgo.Entry("When author has one name", "Hugo", true, "", "Hugo"),
		ginkgo.Entry("When author has a middle name", "Victor Marie Hugo", true, "Victor", "Hugo"),
	)
})

// ginkgo -p -vv --flake-attempts=2
// Ran 6 of 6 Specs in 0.057 seconds
// SUCCESS! -- 6 Passed | 0 Failed | 1 Flaked | 0 Pending | 0 Skipped
func flakehelper() {
	ginkgo.GinkgoHelper()
	if !i {
		i = !i
		ginkgo.Fail("User is too young for this book")
	}
}
