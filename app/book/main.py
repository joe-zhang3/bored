from flask import Flask
from flask_restful import reqparse, abort, Api, Resource
from flask_cors import CORS, cross_origin

app = Flask(__name__)
cors = CORS(app)
api = Api(app, prefix="/api")
app.config['CORS_HEADERS'] = 'Content-Type'

BOOKS = [
        {'id': 'book1', 'name': 'How to learn Java', 'price':100},
        {'id': 'book2', 'name': 'How to learn Python', 'price': 90},
        {'id': 'book3', 'name': 'How to learn C#','price':101},
        {'id': 'book4', 'name': 'How to learn Golang', 'price': 11},
        ]


def abort_if_book_doesnt_exist(book_id):
    if book_id not in BOOKS:
        abort(404, message="book {} doesn't exist".format(book_id))

parser = reqparse.RequestParser()
parser.add_argument('task')


# book
# shows a single book item and lets you delete a book item
class book(Resource):
    def get(self, book_id):
        abort_if_book_doesnt_exist(book_id)
        return BOOKS[book_id]

    def delete(self, book_id):
        abort_if_book_doesnt_exist(book_id)
        del BOOKS[book_id]
        return '', 204

    def put(self, book_id):
        args = parser.parse_args()
        task = {'task': args['task']}
        BOOKS[book_id] = task
        return task, 201


# bookList
# shows a list of all books, and lets you POST to add new tasks
class bookList(Resource):
    def get(self):
        return BOOKS

    def post(self):
        args = parser.parse_args()
        book_id = int(max(bookS.keys()).lstrip('book')) + 1
        book_id = 'book%i' % book_id
        BOOKS[book_id] = {'task': args['task']}
        return BOOKS[book_id], 201

##
## Actually setup the Api resource routing here
##
api.add_resource(bookList, '/books')
api.add_resource(book, '/books/<book_id>')


if __name__ == '__main__':
    app.run(debug=True,host='localhost', port=5000)
