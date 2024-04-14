import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
import Callback from './pages/callback';
import Home from './pages/home';
import OAuth2 from './pages/oauth2';
import Success from './pages/success';

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Home />} />
				<Route path="/oauth2" element={<OAuth2 />} />
				<Route path="/callback" element={<Callback />} />
				<Route path="/success" element={<Success />} />
				<Route
					path="/nouser"
					element={<div>LINE友達登録後、再度お試しください。</div>}
				/>
			</Routes>
		</BrowserRouter>
	);
}

export default App;
