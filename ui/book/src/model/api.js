import axios from 'axios';

export default axios.create({
  baseURL: `http://book_app:5000/api`
});
