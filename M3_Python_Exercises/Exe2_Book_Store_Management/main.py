from book_management import add_book, view_books, search_book
from customer_management import add_customer, view_customers
from sales_management import sell_book, view_sales

def main():
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == '1':
            print("\n1. Add Book")
            print("2. View Books")
            print("3. Search Book")
            sub_choice = input("Enter your choice: ")
            if sub_choice == '1':
                title = input("Title: ")
                author = input("Author: ")
                price = input("Price: ")
                quantity = input("Quantity: ")
                print(add_book(title, author, price, quantity))
            elif sub_choice == '2':
                books = view_books()
                print("\n".join(books) if books else "No books available.")
            elif sub_choice == '3':
                query = input("Enter title or author to search: ")
                results = search_book(query)
                print("\n".join(results))
            else:
                print("Invalid choice.")

        elif choice == '2':
            print("\n1. Add Customer")
            print("2. View Customers")
            sub_choice = input("Enter your choice: ")
            if sub_choice == '1':
                name = input("Name: ")
                email = input("Email: ")
                phone = input("Phone: ")
                print(add_customer(name, email, phone))
            elif sub_choice == '2':
                customers = view_customers()
                print("\n".join(customers) if customers else "No customers available.")
            else:
                print("Invalid choice.")

        elif choice == '3':
            print("\n1. Sell Book")
            print("2. View Sales Records")
            sub_choice = input("Enter your choice: ")
            if sub_choice == '1':
                customer_name = input("Customer Name: ")
                book_title = input("Book Title: ")
                quantity = input("Quantity: ")
                print(sell_book(customer_name, book_title, quantity))
            elif sub_choice == '2':
                sales = view_sales()
                print("\n".join(sales) if sales else "No sales records available.")
            else:
                print("Invalid choice.")

        elif choice == '4':
            print("Thank you for using BookMart!")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()