from customer import Customer

customers = []

def add_customer(name, email, phone):
    if not email or not phone.isdigit():
        return "Error: Invalid customer details."
    customer = Customer(name, email, phone)
    customers.append(customer)
    return "Customer added successfully!"

def view_customers():
    return [customer.display_details() for customer in customers]
