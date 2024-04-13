import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
import Home from './pages/home';
import OAuth2 from './pages/oauth2';

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Home />} />
				<Route path="/oauth2" element={<OAuth2 />} />
				<Route
					path="/nouser"
					element={<div>LINE友達登録後、再度お試しください。</div>}
				/>
			</Routes>
		</BrowserRouter>
	);
}

export default App;
