import sqlite3

def initialize_db():
    conn = sqlite3.connect('bookbuddy.db')
    cursor = conn.cursor()

    # Create books table
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        author TEXT NOT NULL,
        published_year INTEGER NOT NULL,
        genre TEXT NOT NULL
    )
    ''')

    # Insert sample data
    sample_data = [
        ('The Great Gatsby', 'F. Scott Fitzgerald', 1925, 'Fiction'),
        ('To Kill a Mockingbird', 'Harper Lee', 1960, 'Fiction'),
        ('1984', 'George Orwell', 1949, 'Sci-Fi'),
        ('The Catcher in the Rye', 'J.D. Salinger', 1951, 'Fiction'),
    ]

    cursor.executemany('''
    INSERT INTO books (title, author, published_year, genre)
    VALUES (?, ?, ?, ?)
    ''', sample_data)

    conn.commit()
    conn.close()

if __name__ == '__main__':
    initialize_db()
    print("Database initialized and sample data added!")
