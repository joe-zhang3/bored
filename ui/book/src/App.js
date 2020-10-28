import logo from './logo.svg';
import './App.css';
import Header from './components/Header'

function App() {
  return (
  <div className="App" >
    <Header />
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
    </div>
  );
}

export default App;
