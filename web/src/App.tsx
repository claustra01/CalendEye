import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from './pages/home';
import './App.css';
import OAuth2 from "./pages/oauth2";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
				<Route path="/oauth2" element={<OAuth2 />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
