from book import Book

books = []

def add_book(title, author, price, quantity):
    try:
        price = float(price)
        quantity = int(quantity)
        if price <= 0:
            raise ValueError("Inavalid input! Price must be positive numbers.")
        if quantity <= 0:
            raise ValueError("Inavalid input! Quantity must be positive numbers.")
        book = Book(title, author, price, quantity)
        books.append(book)
        return "Book added successfully!"
    except ValueError as e:
        return f"Error: {e}"

def view_books():
    return [book.display_details() for book in books]

def search_book(query):
    results = [book.display_details() for book in books if query.lower() in book.title.lower() or query.lower() in book.author.lower()]
    return results if results else "No matching books found."
