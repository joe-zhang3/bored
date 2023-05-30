import BookRow from "./BookRow"

export default function BookTable() {

    return (
        <table className="table table-bordered">
  <thead>
    <tr>
      <th scope="col">#</th>
      <th scope="col">First</th>
      <th scope="col">Last</th>
      <th scope="col">Handle</th>
    </tr>
  </thead>
  <tbody>
    <BookRow />
  </tbody>
</table>
    )
}