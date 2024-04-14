import { useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

type AuthParams = {
	// Use snake case to match Google's API
	client_id: string;
	redirect_uri: string;
	response_type: string;
	access_type: string;
	scope: string[];
};

const generateAuthUrl = (authParams: AuthParams) => {
	const authParamsWithoutScope: Omit<AuthParams, 'scope'> = {
		client_id: authParams.client_id,
		redirect_uri: authParams.redirect_uri,
		response_type: authParams.response_type,
		access_type: authParams.access_type,
	};
	const queryParams = new URLSearchParams(authParamsWithoutScope);
	queryParams.append('scope', authParams.scope.join(','));
	return `https://accounts.google.com/o/oauth2/v2/auth?${queryParams.toString()}`;
};

const OAuth2 = () => {
	const location = useLocation();
	const navigate = useNavigate();
	const apiUrl = import.meta.env.VITE_API_URL;

	const query = new URLSearchParams(location.search);
	const lineId = query.get('id');

	const checkUserExist = async () => {
		if (!lineId) {
			console.error('id not found in query params');
			navigate('/nouser');
			return;
		}

		try {
			const response = await fetch(`${apiUrl}/user?id=${lineId}`);
			const data = await response.json();
			if (!response.ok) {
				console.error('Failed to fetch user data: ', data);
				navigate('/nouser');
			} else {
				sessionStorage.setItem('id', lineId);
			}
		} catch (error) {
			console.error('Failed to fetch user data: ', error);
			navigate('/nouser');
		}
	};

	const authGoogle = async () => {
		const authParams: AuthParams = {
			client_id: import.meta.env.VITE_GOOGLE_CLIENT_ID,
			redirect_uri: `${import.meta.env.VITE_ORIGIN_URL}/callback`,
			response_type: 'code',
			access_type: 'offline',
			scope: ['https://www.googleapis.com/auth/calendar'],
		};
		window.location.href = generateAuthUrl(authParams);
	};

	useEffect(() => {
		checkUserExist().then(() => {
			authGoogle();
		});
	}, [checkUserExist, authGoogle]);

	return <></>;
};

export default OAuth2;
