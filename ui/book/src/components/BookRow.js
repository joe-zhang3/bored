import React from 'react'
import axios from 'axios'

export default class BookRow extends React.Component{
    state = {
        books: []
      }
    
      componentDidMount() {
        axios.get(`http://127.0.0.1:5000/api/books`)
          .then(res => {
            const books = res.data;
            this.setState({ books });
          })
      }
    
      render() {
        return (
              this.state.books
                .map(book =>
                    <tr>
                    <th >{book.id}</th>
                    <td>{book.name}</td>
                    <td>{book.price}</td>
                  </tr>
                )
        )
      }
}