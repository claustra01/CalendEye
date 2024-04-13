import { useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

const OAuth2 = () => {
	const location = useLocation();
	const navigate = useNavigate();
	const apiUrl = import.meta.env.VITE_API_URL;

	const query = new URLSearchParams(location.search);
	const lineId = query.get('id');

	const fetchData = async () => {
		try {
			console.log(`${apiUrl}/users?id=${lineId}`);
			const response = await fetch(`${apiUrl}/user?id=${lineId}`);
			const data = await response.json();
			if (!response.ok) {
				console.error('Failed to fetch user data: ', data);
				navigate('/nouser');
			}
		} catch (error) {
			console.error('Failed to fetch user data: ', error);
			navigate('/nouser');
		}
	};

	useEffect(() => {
		fetchData();
	}, [fetchData]);

	return <></>;
};

export default OAuth2;
