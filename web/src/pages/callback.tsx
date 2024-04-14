import { useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

const Callback = () => {
	const location = useLocation();
	const navigate = useNavigate();
	const apiUrl = import.meta.env.VITE_API_URL;

	const query = new URLSearchParams(location.search);

	const setToken = async () => {
		const code = query.get('code');
		if (!code) {
			console.error('code not found in query params');
			navigate('/nouser');
			return;
		}

		try {
			const lineId = sessionStorage.getItem('id');
			sessionStorage.removeItem('id');
			const response = await fetch(`${apiUrl}/token`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify({ id: lineId, code: code }),
			});
			if (!response.ok) {
				console.error('Failed to fetch token: ', response);
				navigate('/nouser');
			} else {
				navigate('/success');
			}
		} catch (error) {
			console.error('Failed to fetch token: ', error);
			navigate('/nouser');
		}
	};

	useEffect(() => {
		setToken();
	}, [setToken]);

	return <></>;
};

export default Callback;
