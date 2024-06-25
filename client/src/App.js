import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import './App.css';
import Menu from './components/Menu.js';
import Home from './pages/Home';
import Movies from './pages/Movies';
import Genres from './pages/Genres';
import Admin from './pages/Admin';
import ShowMovie from "./pages/Movies/show.js";
import ShowMovieGenre from "./pages/Genres/Show.js";

function App() {
  return (
    <Router>
      <div className="container">
        <div className="row">
          <h1 className="mt-3">Go React Movie Project!</h1>
          <hr className="mb-3" />
        </div>
        <div className="row">
          <div className="col-2">
            <Menu />
          </div>
          <div className="col-10">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/movies" element={<Movies />} />
              <Route exact path="/movies/:id" element={<ShowMovie />} />
              <Route path="/genres" element={<Genres />} />
              <Route exact path="/genres/:id/movies" element={<ShowMovieGenre />} />
              <Route path="/admin" element={<Admin />} />
            </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
