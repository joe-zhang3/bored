import React from 'react'
import { DataGrid } from '@material-ui/data-grid';
import API from '../model/api';

const columns = [
  { field: 'id', headerName: 'ID', width: 70 },
  { field: 'title', headerName: 'First name', width: 230 },
  { field: 'description', headerName: 'Last name', width: 500}
];

export default class BookList extends React.Component {
  state = {
    books: []
  }

  componentDidMount() {
    API.get('/books')
      .then(res => {
        const books = res.data;
        this.setState({ books });
      })
  }

  render() {
    return (
      <div style={{ height: 400, width: '100%' }}>
      <DataGrid rows={this.state.books} columns={columns} pageSize={5} />
      </div>
    )
  }
}
