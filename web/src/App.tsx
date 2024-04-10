import liff from '@line/liff';
import { useEffect, useState } from 'react';
import './App.css';
import GoogleLogin from './components/GoogleAuth';
import InExternalBrowser from './components/InExternalBrowser';

function App() {
	const [auth, setAuth] = useState<boolean>(false);
	const [userId, setUserId] = useState<string>('');
	const [displayName, setDisplayName] = useState<string>('');

	useEffect(() => {
		liff
			.init({ liffId: import.meta.env.VITE_LIFF_ID })
			.then(async () => {
				const profile = await liff.getProfile();
				setAuth(true);
				setUserId(profile.userId);
				setDisplayName(profile.displayName);
			})
			.catch((e: Error) => {
				setAuth(false);
				alert(e);
			});
	});

	return (
		<>
			{auth ? (
				<GoogleLogin props={{ userId, displayName }} />
			) : (
				<InExternalBrowser />
			)}
		</>
	);
}

export default App;
