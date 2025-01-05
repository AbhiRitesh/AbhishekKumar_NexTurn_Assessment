from transaction import Transaction
from book_management import books

sales = []

def sell_book(customer_name, book_title, quantity):
    try:
        quantity = int(quantity)
        if quantity <= 0:
            raise ValueError("Quantity must be a positive number.")
        for book in books:
            if book.title.lower() == book_title.lower():
                if book.quantity < quantity:
                    return f"Error: Only {book.quantity} copies available. Sale cannot be completed."
                book.quantity -= quantity
                sale = Transaction(customer_name, "", "", book.title, quantity)
                sales.append(sale)
                return f"Sale successful! Remaining quantity: {book.quantity}"
        return "Error: Book not found."
    except ValueError as e:
        return f"Error: {e}"

def view_sales():
    return [sale.display_details() for sale in sales]
