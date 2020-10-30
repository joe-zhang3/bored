import '../App.css';
import logo from '../logo.svg';

export default function Home() {

  return(
    <div id="content">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <a className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer">
          Home Page
        </a>
      </header>
    </div>  
  )
}