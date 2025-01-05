from flask import Flask, request, jsonify
import sqlite3

app = Flask(__name__)

DB_NAME = 'bookbuddy.db'

# Database Helper Functions
def execute_query(query, params=(), fetchone=False, fetchall=False):
    conn = sqlite3.connect(DB_NAME)
    conn.row_factory = sqlite3.Row
    cursor = conn.cursor()
    cursor.execute(query, params)
    data = None
    if fetchone:
        data = cursor.fetchone()
    elif fetchall:
        data = cursor.fetchall()
    conn.commit()
    conn.close()
    return data

# Routes
@app.route('/books', methods=['POST'])
def add_book():
    data = request.get_json()
    try:
        title = data['title']
        author = data['author']
        published_year = int(data['published_year'])
        genre = data['genre']

        query = "INSERT INTO books (title, author, published_year, genre) VALUES (?, ?, ?, ?)"
        
        conn = sqlite3.connect(DB_NAME)
        cursor = conn.cursor()
        cursor.execute(query, (title, author, published_year, genre))
        book_id = cursor.lastrowid  # Retrieve the ID of the inserted book
        conn.commit()
        conn.close()

        return jsonify({"message": "Book added successfully", "book_id": book_id}), 201
    except KeyError:
        return jsonify({"error": "Invalid data", "message": "Missing required fields"}), 400
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500


@app.route('/books', methods=['GET'])
def get_books():
    books = execute_query("SELECT * FROM books", fetchall=True)
    return jsonify([dict(book) for book in books]), 200

@app.route('/books/<int:book_id>', methods=['GET'])
def get_book(book_id):
    book = execute_query("SELECT * FROM books WHERE id = ?", (book_id,), fetchone=True)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404
    return jsonify(dict(book)), 200

@app.route('/books/<int:book_id>', methods=['PUT'])
def update_book(book_id):
    data = request.get_json()
    book = execute_query("SELECT * FROM books WHERE id = ?", (book_id,), fetchone=True)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    try:
        title = data.get('title', book['title'])
        author = data.get('author', book['author'])
        published_year = int(data.get('published_year', book['published_year']))
        genre = data.get('genre', book['genre'])

        query = "UPDATE books SET title = ?, author = ?, published_year = ?, genre = ? WHERE id = ?"
        execute_query(query, (title, author, published_year, genre, book_id))
        return jsonify({"message": "Book updated successfully"}), 200
    except Exception as e:
        return jsonify({"error": "Invalid data", "message": str(e)}), 400

@app.route('/books/<int:book_id>', methods=['DELETE'])
def delete_book(book_id):
    book = execute_query("SELECT * FROM books WHERE id = ?", (book_id,), fetchone=True)
    if not book:
        return jsonify({"error": "Book not found", "message": "No book exists with the provided ID"}), 404

    execute_query("DELETE FROM books WHERE id = ?", (book_id,))
    return jsonify({"message": "Book deleted successfully"}), 200

@app.route('/books', methods=['GET'])
def get_filtered_books():
    genre = request.args.get('genre')
    author = request.args.get('author')
    query = "SELECT * FROM books WHERE 1=1"
    params = []
    if genre:
        query += " AND genre = ?"
        params.append(genre)
    if author:
        query += " AND author = ?"
        params.append(author)
    books = execute_query(query, params, fetchall=True)
    return jsonify([dict(book) for book in books]), 200


if __name__ == '__main__':
    app.run(debug=True)
