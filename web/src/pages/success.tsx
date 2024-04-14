import { useEffect } from 'react';

const Success = () => {
	useEffect(() => {
		window.location.href = 'line://';
	}, []);

	return <div>ログインに成功しました。</div>;
};

export default Success;
