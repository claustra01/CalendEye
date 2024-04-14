import { useEffect } from 'react';

const Success = () => {
	useEffect(() => {
		window.location.href = 'line://';
	}, []);

	return <div>ログインに成功しました。LINEアプリに戻ります。</div>;
};

export default Success;
